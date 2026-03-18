// Package handlers provides HTTP handler functions for context-service.
package handlers

import (
	"log"
	"net/http"
)

// Root is the health-check route.
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Printf("error writing response: %v\n", err)
	}
}
