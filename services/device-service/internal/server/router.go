// Package server provides HTTP server for startup
package server

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers"

	_ "innoveria-iot/device-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter creates and returns an HTTP ServeMux with all device service routes registered.
func NewRouter(
	gatewaySvc domain.GatewayService,
	sensorSvc domain.SensorService,
	sensorProfileSvc domain.SensorProfileService,
	companyConfigSvc domain.CompanyConfigService,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Gateway Routes:
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGateways(gatewaySvc))
	mux.HandleFunc("POST "+GATEWAY_ROUTE, handlers.PostGateway(gatewaySvc))
	mux.HandleFunc("PUT "+GATEWAY_ROUTE_ID, handlers.PutGateway(gatewaySvc))
	mux.HandleFunc("DELETE "+GATEWAY_ROUTE_ID, handlers.DeleteGateway(gatewaySvc))

	// Sensor Routes:
	mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))
	mux.HandleFunc("POST "+SENSOR_ROUTE, handlers.PostSensor(sensorSvc))
	mux.HandleFunc("DELETE "+SENSOR_ROUTE_ID, handlers.DeleteSensor(sensorSvc))
	mux.HandleFunc("PUT "+SENSOR_ROUTE_ID, handlers.PutSensor(sensorSvc))

	// Sensor profile routes:
	mux.HandleFunc("GET "+SENSOR_PROFILE_ROUTE, handlers.GetAllSensorProfiles(sensorProfileSvc))

	// Company config routes:
	mux.HandleFunc("POST "+COMPANY_CONFIG_ROUTE, handlers.PostCompanyConfig(companyConfigSvc))
	mux.HandleFunc("DELETE "+COMPANY_CONFIG_ROUTE_ID, handlers.DeleteCompanyConfig(companyConfigSvc))

	// Swagger docs
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
