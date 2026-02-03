package routing

import (
	"innoveria-iot/api-gateway/internal/handlers"
	"net/http"
)

// NewRouter configures the HTTP router
// Route path are found in the const.go file
func NewRouter() *http.ServeMux  {
	mux := http.NewServeMux()
	
	// Routes:
	mux.HandleFunc(INDEX,handlers.Root)
	mux.HandleFunc(COLLECTION_ROUTE,handlers.Collection)
	mux.HandleFunc(AUTHENTICATION_ROUTE,handlers.Authentication)
	mux.HandleFunc(CONTEXT_ROUTE,handlers.Context)

	return mux
}
