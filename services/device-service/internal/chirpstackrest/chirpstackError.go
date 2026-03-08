// Package chirpstackrest provides an HTTP client for communicating with the Chirpstack REST API.
package chirpstackrest

import (
	"encoding/json"
	"errors"
	"fmt"
	"innoveria-iot/pkg/httpclient"
)

// ChirpstackError represents a structured error response from the Chirpstack API.
type ChirpstackError struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}

// ErrorDetail holds additional context about a specific Chirpstack error.
type ErrorDetail struct {
	Type  string            `json:"@type"`
	Props map[string]string `json:"-"`
}

// Error returns a formatted string representation of the Chirpstack error.
func (e *ChirpstackError) Error() string {
	return fmt.Sprintf("chirpstack error: %s (code=%d)", e.Message, e.Code)
}

// handleChirpstackError parses an HTTP error into a typed ChirpstackError if possible,
// or falls back to a plain error wrapping the status code and body.
func handleChirpstackError(err error) error {
	var httpErr *httpclient.HTTPError
	if !errors.As(err, &httpErr) {
		return err
	}

	var apiErr ChirpstackError
	if json.Unmarshal(httpErr.Body, &apiErr) == nil {
		// Return typed error so service layer can inspect it
		return &ChirpstackError{
			Code:    apiErr.Code,
			Message: apiErr.Message,
			Details: apiErr.Details,
		}
	}

	// Fallback if body isn't structured JSON
	return fmt.Errorf("chirpstack error (%d): %s",
		httpErr.StatusCode,
		string(httpErr.Body),
	)
}
