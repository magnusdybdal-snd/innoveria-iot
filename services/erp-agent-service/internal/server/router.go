package server

import (
	"net/http"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	//mux.HandleFunc("GET /", handlers.Root)

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
