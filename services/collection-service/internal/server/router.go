package server

import (
	"innoveria-iot/collection-service/internal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux  {
	mux := http.NewServeMux()
	
	// Routes:
	mux.HandleFunc("GET /",handlers.Root)

	return mux
}
