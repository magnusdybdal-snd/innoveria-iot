package server

import (
	"innoveria-iot/collection-service/internal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET /collection", handlers.HelloProxy)
	mux.HandleFunc("GET /api/v1/collection/latest", handlers.HelloProxy) // test for proxy

	return mux
}
