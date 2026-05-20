package server

import (
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers"

	_ "innoveria-iot/context-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(contextSvc domain.ContextService, ruleSvc domain.RuleService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	// Rules endpoints
	mux.HandleFunc("GET "+RULES_ROUTE, handlers.GetRules(ruleSvc))
	mux.HandleFunc("POST "+RULES_ROUTE, handlers.CreateRule(ruleSvc))
	mux.HandleFunc("DELETE "+RULES_ID_ROUTE, handlers.DeleteRule(ruleSvc))
	// Context data endpoint
	mux.HandleFunc("GET "+CONTEXT_DATA_ROUTE, handlers.GetContextData(contextSvc))
	// Orders endpoints
	mux.HandleFunc("GET "+ORDERS_ROUTE, handlers.GetOrders(contextSvc))
	mux.HandleFunc("GET "+ORDER_BY_ID_ROUTE, handlers.GetOrderByID(contextSvc))
	mux.HandleFunc("GET "+ORDER_CONTEXT_ROUTE, handlers.GetOrderContext(contextSvc))
	// Production resources endpoint
	mux.HandleFunc("GET "+PRODUCTION_RESOURCES_ROUTE, handlers.GetProductionResources(contextSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
