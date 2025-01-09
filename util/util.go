package util

import (
	"io/ioutil" 
	"net/http"
	"encoding/json"
	"strings"
	"errors"
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
