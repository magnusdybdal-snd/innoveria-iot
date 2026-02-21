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
	mux.HandleFunc(INDEX, handlers.Root)

	/*
		Proxy routes microservice:
	*/
	// Device service
	handlers.RegisterProxyService(mux, DEVICE_ROUTE, "device-service", cfg.DeviceSvcURL, []string{
		"/gateways",
		"/sensors",
	})

	return mux
}
