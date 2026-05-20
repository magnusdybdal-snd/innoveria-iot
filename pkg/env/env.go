// Package env provides helpers for reading environment variables with fallbacks and type conversion.
package env

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Enviroment variable handeling

// Get will fallback if the value is empty
func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Required will return error if env value is not set
// used for secrets loading
func Required(key string) (string, error) {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return "", fmt.Errorf("missing required env: %s", key)
	}
	return v, nil
}

// GetBool will fallback if the value is missing or not a bool
func GetBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		boolVal, err := strconv.ParseBool(v)
		if err != nil {
			return fallback
		}
		return boolVal
	}
	return fallback
}

// GetFile is for reading enviroment variables from a file
func GetFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
