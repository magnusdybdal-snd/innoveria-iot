package server

import (
	"net/http"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(contextSvc domain.ContextService, ruleSvc domain.RuleService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET /rules", handlers.GetRules(ruleSvc))

	// Swagger docs
	// mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
