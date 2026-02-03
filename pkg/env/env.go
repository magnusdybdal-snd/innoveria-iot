package env

import "os"

// Enviroment variable handler
// Will fallback if the value is empty
func Get(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
