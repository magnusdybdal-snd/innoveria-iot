// Package json provides HTTP response helpers for encoding JSON and writing error responses.
package json

import (
	"log/slog"
	"net/http"
)

// ErrorResponse is the top-level JSON envelope returned on error.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail carries the HTTP status code and a human-readable error message.
type ErrorDetail struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// HandleError logs the error and writes a JSON error response with the given status code and message.
func HandleError(w http.ResponseWriter, code int, err error, msg string) {
	if err == nil {
		slog.Warn("HandleError called with nil error", "code", code, "message", msg)
		return
	}

	slog.Error("request error", "code", code, "message", msg, "error", err)

	resp := ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: msg,
		},
	}
	if encErr := Encode(w, code, resp); encErr != nil {
		slog.Error("failed to write error response", "error", encErr)
	}
}
