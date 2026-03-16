// Package handlers provides HTTP endpoints for auth service
package handlers

import (
	"log/slog"
	"net/http"

	"innoveria-iot/pkg/json"
)

// Root handles the base endpoint and returns a simple service status payload.
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.Encode(w, http.StatusOK, map[string]string{
		"name":   "auth-service",
		"status": "ok",
	})
	if err != nil {
		slog.Error("error encoding json", "err", err)
	}
}
