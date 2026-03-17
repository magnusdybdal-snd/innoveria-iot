package server

import (
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers"

	_ "innoveria-iot/auth-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(companySvc domain.CompanyService, factorySvc domain.FactoryService, factoryAreaSvc domain.FactoryAreaService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Company
	mux.HandleFunc("GET "+COMPANY_ROUTE, handlers.GetAllCompanies(companySvc))
	mux.HandleFunc("GET "+COMPANY_ID_ROUTE, handlers.GetOneCompany(companySvc))
	mux.HandleFunc("POST "+COMPANY_ROUTE, handlers.PostCompany(companySvc))
	mux.HandleFunc("DELETE "+COMPANY_ID_ROUTE, handlers.DeleteCompany(companySvc))

	// Factory
	mux.HandleFunc("GET "+FACTORY_ROUTE, handlers.GetAllFactories(factorySvc))
	mux.HandleFunc("GET "+FACTORY_ID_ROUTE, handlers.GetOneFactory(factorySvc))
	mux.HandleFunc("POST "+FACTORY_ROUTE, handlers.PostFactory(factorySvc))
	mux.HandleFunc("DELETE "+FACTORY_ID_ROUTE, handlers.DeleteFactory(factorySvc))

	// Factory Area
	mux.HandleFunc("POST "+FACTORY_AREA_ROUTE, handlers.PostFactoryArea(factoryAreaSvc))
	mux.HandleFunc("GET "+FACTORY_AREA_ROUTE, handlers.GetAllFactoryAreas(factoryAreaSvc))
	mux.HandleFunc("GET "+FACTORY_AREA_ID_ROUTE, handlers.GetOneFactoryArea(factoryAreaSvc))
	mux.HandleFunc("DELETE "+FACTORY_AREA_ID_ROUTE, handlers.DeleteFactoryArea(factoryAreaSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
