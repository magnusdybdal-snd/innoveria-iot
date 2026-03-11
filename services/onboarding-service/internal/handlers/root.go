// Package handlers provides HTTP handlers for the onboarding service.
package handlers

import (
	"log/slog"
	"net/http"

	"innoveria-iot/pkg/json"
)

// Root handles GET / and returns the service name and status. Returns 404 for any other path.
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.Encode(w, http.StatusOK, map[string]string{
		"name":   "onboarding-service",
		"status": "ok",
	})
	if err != nil {
		slog.Error("error encoding json", "err", err)
	}
}
