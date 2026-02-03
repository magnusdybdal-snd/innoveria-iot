package routing

import (
	"innoveria-iot/api-gateway/internal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux  {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/",handlers.Root)

	mux.HandleFunc("/api/v1/collection",handlers.Collection)

	return mux
}
