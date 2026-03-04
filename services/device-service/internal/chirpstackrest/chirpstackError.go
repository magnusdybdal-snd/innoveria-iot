// Package chirpstackrest TODO(@vinjar): add proper documentation.
package chirpstackrest

import (
	"encoding/json"
	"errors"
	"fmt"
	"innoveria-iot/pkg/httpclient"
)

// ChirpstackError TODO(@vinjar): add proper documentation.
type ChirpstackError struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details"`
}

// ErrorDetail TODO(@vinjar): add proper documentation.
type ErrorDetail struct {
	Type  string            `json:"@type"`
	Props map[string]string `json:"-"`
}

// Error TODO(@vinjar): add proper documentation.
func (e *ChirpstackError) Error() string {
	return fmt.Sprintf("chirpstack error: %s (code=%d)", e.Message, e.Code)
}

// Error handling for chirpstack requests
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
