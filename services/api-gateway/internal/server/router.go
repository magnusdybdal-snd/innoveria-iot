package server

import (
	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/handlers"
	"net/http"
)

// NewRouter configures the HTTP router
// Route path are found in the const.go file
func NewRouter(cfg *config.Config) *http.ServeMux  {
	mux := http.NewServeMux()
	
	// Routes:
	mux.HandleFunc(INDEX,handlers.Root)
	collProxy, err := handlers.NewUpstreamProxy(cfg.CollSvcURL)
	if err != nil {
		handlers.RegisterServiceError(mux,COLLECTION_ROUTE,"collection-service")
	} else {
		handlers.RegisterServiceInfo(mux,COLLECTION_ROUTE,"collection-service",[]string{
			"/latest", // latest sensor info
		})
		mux.Handle(COLLECTION_ROUTE+"/",collProxy)
	}
	mux.HandleFunc(AUTHENTICATION_ROUTE,handlers.Authentication)
	mux.HandleFunc(CONTEXT_ROUTE,handlers.Context)

	// Collection proxy

	return mux
}
