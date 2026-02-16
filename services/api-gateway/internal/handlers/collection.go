package handlers

import (
	"log"
	"net/http"
)

// TODO: Handle logic
func Collection(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := w.Write([]byte("hello collection"))
	if err != nil {
		log.Printf("Error writing response: %v\n", err)
		return
	}
}
