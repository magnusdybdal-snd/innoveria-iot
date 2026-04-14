package server

import (
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/erp-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(ingestSvc domain.Ingest) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// ERP AGENT SERVICE Endpoints
	mux.HandleFunc("POST "+ORDER_ROUTE, handlers.PostIngestWorkCenters(ingestSvc))
	mux.HandleFunc("POST "+ORDER_OPERATIONS_ROUTE, handlers.PostIngestOrder(ingestSvc))
	mux.HandleFunc("POST "+ORDER_REPORTINGS_ROUTE, handlers.PostIngestOrderOperations(ingestSvc))
	mux.HandleFunc("POST "+WORKCENTERS_ROUTE, handlers.PostIngestOrderReports(ingestSvc))

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
