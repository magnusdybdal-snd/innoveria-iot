package handlers

import (
	"log"
	"net/http"
)

// Home route
func Root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := w.Write([]byte("ok"))
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
		return
	}
}
