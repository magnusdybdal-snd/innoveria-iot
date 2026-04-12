package server

import (
	"net/http"

	"innoveria-iot/erp-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// ERP AGENT SERVICE Endpoints
	mux.HandleFunc("POST "+ORDER_OPERATIONS_ROUTE, handlers.PostIngestOrderOperations())
	mux.HandleFunc("POST "+ORDER_REPORTINGS_ROUTE, handlers.PostIngestOrderOperations())
	mux.HandleFunc("POST "+WORKCENTERS_ROUTE, handlers.PostIngestOrderOperations())

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
