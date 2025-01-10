package util

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512" // This package provides sha384.New
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func StringContainsSubstring(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func ArrayContainsString(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
} 

func getEnvOrDefault(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getBaseAPIUrl() string {
	apiEnvironment := getEnvOrDefault("GEMINI_EXCHANGE_API_ENVIRONMENT", "production")
	envUrls := map[string]string{
		"sandbox": "https://api.sandbox.gemini.com",
		"production": "https://api.gemini.com",
	}
	baseUrl := envUrls[apiEnvironment]
	return baseUrl
}


func GetPublicEndpoint(endpoint string, target interface{}) error {
	/* 
	Perform an HTTP GET request on a public Gemini API endpoint and unmarshal the JSON response into the provided target interface.

	Args: 
	endpoint (string) - API endpoint to use
	target - any type of JSON object in which the JSON response gets stored, passed as &target (pointer to a variable in which response is to be stored) when method is invoked

	Returns nothing if successful, returns error if it fails

	*/ 

	url := getBaseAPIUrl() + endpoint

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return errors.New("error creating request: " + err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return errors.New("error making request: " + err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("received non-200 status code: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return errors.New("error reading response body: " + err.Error())
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("error unmarshalling response: " + err.Error())
	}

	return nil
}


func DownloadPublicFile(url string, filePath string) error {
	/* 
	Download a file from a url to a specific file path

	Args: 
		url : url from which to download file	
		filePath: path on local file system to which downloaded file will be saved

	Returns:
		error: Returns an error if the file cannot be downloaded.
	*/

	// Perform the HTTP GET request to download the file.
	response, err := http.Get(url)
	if err != nil {
		return errors.New("failed to download funding amount report: " + err.Error())
	}
	defer response.Body.Close()

	// Check if the response status is OK.
	if response.StatusCode != http.StatusOK {
		return errors.New("unexpected HTTP status: " + response.Status)
	}

	// Create the file locally.
	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("failed to create the file: " + err.Error())
	}
	defer file.Close()

	// Copy the response body to the file.
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return errors.New("failed to save the report: " + err.Error())
	}

	log.Printf("Report downloaded successfully and saved to %s\n", filePath) 
	return nil
}


func PostPrivateEndpoint(endpoint string, payload map[string]interface{}, target interface{}) error {
	/* 
	Perform an HTTP POST request on a private Gemini API endpoint and unmarshal the JSON response into the provided target interface.

	Args: 
	endpoint (string) - endpoint path (e.g., /v1/order/new) to make request against. Environment variable API_ENVIRONMENT determines whether to use api.gemini.com (prod) or api.sandbox.gemini.com (sandbox)
	target - any type of JSON object in which the JSON response gets stored, passed as &target (pointer to a variable in which response is to be stored) when method is invoked

	Returns nothing if successful, returns error if it fails
	*/

	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Unable to load .env file:", err)
	}

	// Ensure required API keys are set
	if os.Getenv("GEMINI_EXCHANGE_API_KEY") == "" || len([]byte(os.Getenv("GEMINI_EXCHANGE_API_SECRET"))) == 0 {
		log.Println("GEMINI_EXCHANGE_API_KEY and GEMINI_EXCHANGE_API_SECRET are not both present in environment")
	} else {
		log.Println("Private credentials from .env successfully loaded for private API functions!")
	}

	// Retrieve API key and secret from environment
	apiKey := getEnvOrDefault("GEMINI_EXCHANGE_API_KEY", "")
	apiSecret := []byte(getEnvOrDefault("GEMINI_EXCHANGE_API_SECRET", ""))

	// Build the full URL
	url := getBaseAPIUrl() + endpoint

	// Create the nonce for the payload
	nonce := time.Now().UTC().Unix()

	// Merge the provided payload with auto-generated values
	mergedPayload := make(map[string]interface{})
	for key, value := range payload {
		mergedPayload[key] = value
	}
	mergedPayload["request"] = endpoint
	mergedPayload["nonce"] = strconv.FormatInt(nonce, 10)

	// Encode the payload to JSON
	payloadJSON, err := json.Marshal(mergedPayload)
	if err != nil {
		return errors.New("Error encoding payload: " + err.Error())
	}

	// Base64 encode the JSON payload
	b64Payload := base64.StdEncoding.EncodeToString(payloadJSON)

	// Create the HMAC signature using SHA384 
	h := hmac.New(sha512.New384, apiSecret)
	h.Write([]byte(b64Payload))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	// Prepare the HTTP request headers
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return errors.New("Error creating request: " + err.Error())
	}

	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Content-Length", "0")
	req.Header.Set("X-GEMINI-APIKEY", apiKey)
	req.Header.Set("X-GEMINI-PAYLOAD", b64Payload)
	req.Header.Set("X-GEMINI-SIGNATURE", signature)
	req.Header.Set("Cache-Control", "no-cache")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New("Error sending request: " + err.Error())
	}
	defer resp.Body.Close()

	// Read the response
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	// Handle non-OK HTTP status codes
	if resp.StatusCode != http.StatusOK {
		return errors.New("Error: status " + resp.Status + ", response: " + buf.String())
	}

	// Parse the response JSON into the target interface
	if err := json.Unmarshal(buf.Bytes(), target); err != nil {
		return errors.New("Error parsing response JSON: " + err.Error())
	}

	return nil
}
