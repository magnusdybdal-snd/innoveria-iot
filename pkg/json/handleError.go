package json

import (
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

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
