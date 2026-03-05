package server

import (
	"innoveria-iot/auth-service/internal/handlers"
	"net/http"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	return mux
}
