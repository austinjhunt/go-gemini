package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

// GetPublicEndpoint performs an HTTP GET request and unmarshals the JSON response into the provided target interface.
func GetPublicEndpoint(url string, target interface{}) error {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New("error reading response body: " + err.Error())
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return errors.New("error unmarshalling response: " + err.Error())
	}

	return nil
}


func DownloadFile(url string, filePath string) error {
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
 