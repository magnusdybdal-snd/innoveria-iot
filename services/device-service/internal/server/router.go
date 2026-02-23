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
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGateways(gatewaySvc))
	mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))

	return mux
}
