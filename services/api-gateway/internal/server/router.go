package server

import (
	"net/http"

	"innoveria-iot/api-gateway/internal/config"
	"innoveria-iot/api-gateway/internal/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
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
		"/factories",
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

	// Onboarding service
	handlers.RegisterProxyService(mux, ONBOARDING_ROUTE, "onboarding-service", cfg.OnboardingSvcURL, []string{
		"/company",
	})

	// Swagger — merged spec from all services, served via the gateway
	mux.HandleFunc("GET /swagger/doc.json", handlers.MergedSwaggerSpec(cfg.DeviceSvcURL, cfg.CollSvcURL, cfg.AuthSvcURL, cfg.OnboardingSvcURL))
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	return mux
}
