// Package json TODO(@vinjar): add proper documentation.
package json

import (
	"log/slog"
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
