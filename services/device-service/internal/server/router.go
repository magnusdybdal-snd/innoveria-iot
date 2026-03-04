package server

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers"
)

// NewRouter TODO(@vinjar): add proper documentation.
func NewRouter(
	gatewaySvc domain.GatewayService,
	sensorSvc domain.SensorService,
	sensorProfileSvc domain.SensorProfileService,
	sensorGroupSvc domain.SensorGroupService,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Gateway Routes:
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGateways(gatewaySvc))
	mux.HandleFunc("POST "+GATEWAY_ROUTE, handlers.PostGateway(gatewaySvc))
	mux.HandleFunc("PUT "+GATEWAY_ROUTE_ID, handlers.PutGateway(gatewaySvc))
	mux.HandleFunc("DELETE "+GATEWAY_ROUTE_ID, handlers.DeleteGateway(gatewaySvc))

	// Sensor Routes:
	// mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))
	// mux.HandleFunc("POST "+SENSOR_ROUTE, handlers.PostSensors(sensorSvc)) // TODO: add this, when logic is right

	// Sensor profile routes:
	mux.HandleFunc("GET "+SENSOR_PROFILE_ROUTE, handlers.GetAllSensorProfiles(sensorProfileSvc))

	// SensorGroup routes
	mux.HandleFunc("POST "+SENSOR_GROUP_ROUTE, handlers.PostSensorGroup(sensorGroupSvc))
	mux.HandleFunc("GET "+SENSOR_GROUP_ROUTE, handlers.GetAllSensorGroups(sensorGroupSvc))

	return mux
}
