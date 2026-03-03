package env

import (
	"os"
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

// GetFile is for reading enviroment variables from a file
func GetFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
