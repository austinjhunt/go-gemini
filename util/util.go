package util

import (

	// This package provides sha384.New

	"log"
	"os"
	"strings"
)

// Function to check if a slice contains a string
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func Debug(msg string) {
	LOGLEVEL := strings.ToLower(GetEnvOrDefault("LOGLEVEL", "INFO"))
	if LOGLEVEL == "debug" {
		log.Printf("[DEBUG] %s", msg)
	}
}
func Info(msg string) {
	LOGLEVEL := strings.ToLower(GetEnvOrDefault("LOGLEVEL", "INFO"))
	levels := []string{"debug", "info"}
	if contains(levels, LOGLEVEL) {
		log.Printf("[INFO] %s", msg)
	}
}
func Warn(msg string) {
	LOGLEVEL := strings.ToLower(GetEnvOrDefault("LOGLEVEL", "INFO"))
	levels := []string{"debug", "info", "warn"}
	if contains(levels, LOGLEVEL) {
		log.Printf("[INFO] %s", msg)
	}
}
func Err(msg string) {
	log.Printf("[INFO] %s", msg)
}

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

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetBaseAPIUrl() string {
	apiEnvironment := GetEnvOrDefault("GEMINI_EXCHANGE_API_ENVIRONMENT", "production")
	envUrls := map[string]string{
		"sandbox":    "https://api.sandbox.gemini.com",
		"production": "https://api.gemini.com",
	}
	baseUrl := envUrls[apiEnvironment]
	return baseUrl
}
