package server

import (
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers"
	"net/http"
)

func NewRouter(gatewaySvc domain.GatewayService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGatways(gatewaySvc))

	return mux
}
