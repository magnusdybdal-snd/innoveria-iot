// Package json TODO(@vinjar): add proper documentation.
package json

import (
	"log"
	"net/http"
)

// ErrorResponse TODO(@vinjar): add proper documentation.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail TODO(@vinjar): add proper documentation.
type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleError TODO(@vinjar): add proper documentation.
func HandleError(w http.ResponseWriter, code int, err error, msg string) {
	if err != nil {
		// Log error for server
		log.Printf("ERROR: %v", err)

		// Build JSON response
		resp := ErrorResponse{
			Error: ErrorDetail{
				Code:    code,
				Message: msg,
			},
		}
		Encode(w, code, resp) //nolint:errcheck // TODO(vinjar): handle Encode error in error handler
	}
}
