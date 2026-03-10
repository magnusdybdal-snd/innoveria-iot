package server

import (
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers"

	_ "innoveria-iot/auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(authSvc domain.AuthService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET "+COMPANY_ROUTE, handlers.GetAllCompanies(authSvc))
	mux.HandleFunc("POST "+COMPANY_ROUTE, handlers.PostCompany(authSvc))
	mux.HandleFunc("DELETE "+COMPANY_ROUTE+"/{companyID}", handlers.DeleteCompany(authSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
