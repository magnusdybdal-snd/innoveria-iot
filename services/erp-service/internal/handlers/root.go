// Package handlers provides HTTP handler functions for erp-service.
package handlers

import (
	"log/slog"
	"net/http"

	"innoveria-iot/pkg/json"
)

// Root is the health-check route.
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	err := json.Encode(w, http.StatusOK, map[string]string{
		"name":   "erp-service",
		"status": "ok",
	})
	if err != nil {
		slog.Error("error encoding json", "err", err)
	}
}
