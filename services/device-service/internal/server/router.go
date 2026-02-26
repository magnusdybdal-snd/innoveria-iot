package server

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers"
)

func NewRouter(
	gatewaySvc domain.GatewayService,
	sensorSvc domain.SensorService,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Gateway Routes:
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGateways(gatewaySvc))
	mux.HandleFunc("POST "+GATEWAY_ROUTE, handlers.PostGateway(gatewaySvc))
	mux.HandleFunc("PUT "+GATEWAY_ROUTE_ID, handlers.PutGateway(gatewaySvc))

	// Sensor Routes:
	mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))

	return mux
}
