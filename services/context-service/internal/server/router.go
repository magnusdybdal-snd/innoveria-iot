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
	mux.HandleFunc("GET "+RULES_ROUTE, handlers.GetRules(ruleSvc))
	mux.HandleFunc("GET "+CONTEXT_DATA_ROUTE, handlers.GetContextData(contextSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
