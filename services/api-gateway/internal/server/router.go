package server

import (
	"net/http"

	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/handlers"
)

// NewRouter configures the HTTP router
// Route path are found in the const.go file
func NewRouter(cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()

	// Root handler
	// mux.HandleFunc("GET /", handlers.Root)

	/*
		Proxy routes microservice:
	*/

	// auth service
	handlers.RegisterProxyService(mux, AUTHENTICATION_ROUTE, "auth-service", cfg.AuthSvcURL, []string{
		"/companies",
	})

	// Device service
	handlers.RegisterProxyService(mux, DEVICE_ROUTE, "device-service", cfg.DeviceSvcURL, []string{
		"/gateways",
		"/sensors",
		"/sensor-profiles",
		"/sensor-group",
	})

	// Collection service
	handlers.RegisterProxyService(mux, COLLECTION_ROUTE, "collection-service", cfg.CollSvcURL, []string{
		"/latest",
		"/measurements",
	})

	return mux
}
