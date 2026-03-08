package server

import (
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers"
)

// NewRouter builds and returns the service HTTP router.
func NewRouter(authSvc domain.AuthService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("POST "+COMPANY_ROUTE, handlers.PostCompany(authSvc))

	return mux
}
