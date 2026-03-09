package server

import (
	"net/http"

	"innoveria-iot/onboarding-service/internal/domain"
	"innoveria-iot/onboarding-service/internal/handlers"

	_ "innoveria-iot/onboarding-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter creates and returns an HTTP ServeMux with all onboarding service routes registered.
func NewRouter(
	onboardingSvc domain.OnboardingService,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Onboarding routes:
	mux.HandleFunc("POST "+COMPANY_ROUTE, handlers.PostCompany(onboardingSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
