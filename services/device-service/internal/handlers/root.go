package handlers

import (
	"net/http"

	"innoveria-iot/pkg/json"
)

// Home route
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	json.Encode(w, http.StatusOK, map[string]string{
		"name":   "device-service",
		"status": "ok",
	})
}
