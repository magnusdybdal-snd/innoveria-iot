package server

import (
	"net/http"

	"innoveria-iot/erp-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
