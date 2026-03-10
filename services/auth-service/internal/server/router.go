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

	// Company
	mux.HandleFunc("GET "+COMPANY_ROUTE, handlers.GetAllCompanies(authSvc))
	mux.HandleFunc("GET "+COMPANY_ID_ROUTE, handlers.GetOneCompany(authSvc))
	mux.HandleFunc("POST "+COMPANY_ROUTE, handlers.PostCompany(authSvc))
	mux.HandleFunc("DELETE "+COMPANY_ID_ROUTE, handlers.DeleteCompany(authSvc))

	// Factory
	mux.HandleFunc("GET "+FACTORY_ROUTE, handlers.GetAllFactories(authSvc))
	mux.HandleFunc("GET "+FACTORY_ROUTE, handlers.GetOneFactory(authSvc))
	mux.HandleFunc("POST "+FACTORY_ROUTE, handlers.PostFactory(authSvc))
	mux.HandleFunc("DELETE "+FACTORY_ROUTE, handlers.DeleteFactory(authSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
