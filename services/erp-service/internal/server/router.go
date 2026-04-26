package server

import (
	"net/http"

	"innoveria-iot/erp-service/internal/domain"
	"innoveria-iot/erp-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(
	ingestSvc domain.Ingest,
	prodResSvc domain.ProductionResourceSvc,
	orderSvc domain.OrderService,
	jwtSecret string,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// ERP AGENT SERVICE Endpoints
	mux.Handle("POST "+ORDER_ROUTE, RequireAgentAuth(jwtSecret, handlers.PostIngestOrder(ingestSvc)))
	mux.Handle("POST "+ORDER_OPERATIONS_ROUTE, RequireAgentAuth(jwtSecret, handlers.PostIngestOrderOperations(ingestSvc)))
	mux.Handle("POST "+ORDER_REPORTINGS_ROUTE, RequireAgentAuth(jwtSecret, handlers.PostIngestOrderReports(ingestSvc)))
	mux.Handle("POST "+WORKCENTERS_ROUTE, RequireAgentAuth(jwtSecret, handlers.PostIngestWorkCenters(ingestSvc)))

	// Extracting monitor erp data
	mux.HandleFunc("GET "+PRODUCTION_RESOURCE_ROUTE, handlers.GetAllProductionResources(prodResSvc))
	mux.HandleFunc("GET "+PRODUCTION_RESOURCE_ID_ROUTE, handlers.GetOneProductionResource(prodResSvc))
	mux.HandleFunc("GET "+ORDER_SUMMARY_ROUTE, handlers.GetOrderSummaryHandler(orderSvc))
	mux.HandleFunc("GET "+ORDER_ID_ROUTE, handlers.GetOneOrderHandler(orderSvc))

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
