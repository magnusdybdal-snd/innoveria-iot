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
	measurementTypeSvc domain.MeasurementTypeService,
	enableSwagger bool,
) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.Root)

	// Gateway Routes:
	mux.HandleFunc("GET "+GATEWAY_ROUTE, handlers.GetGateways(gatewaySvc))
	mux.HandleFunc("POST "+GATEWAY_ROUTE, handlers.PostGateway(gatewaySvc))
	mux.HandleFunc("PATCH "+GATEWAY_ROUTE_ID, handlers.PatchGateway(gatewaySvc))
	mux.HandleFunc("DELETE "+GATEWAY_ROUTE_ID, handlers.DeleteGateway(gatewaySvc))

	// Sensor Routes:
	mux.HandleFunc("GET "+SENSOR_ROUTE, handlers.GetSensors(sensorSvc))
	mux.HandleFunc("POST "+SENSOR_ROUTE, handlers.PostSensor(sensorSvc))
	mux.HandleFunc("DELETE "+SENSOR_ROUTE_ID, handlers.DeleteSensor(sensorSvc))
	mux.HandleFunc("PATCH "+SENSOR_ROUTE_ID, handlers.PatchSensor(sensorSvc))

	// Sensor profile routes:
	mux.HandleFunc("GET "+SENSOR_PROFILE_ROUTE, handlers.GetAllSensorProfiles(sensorProfileSvc))

	// Company config routes:
	mux.HandleFunc("POST "+COMPANY_CONFIG_ROUTE, handlers.PostCompanyConfig(companyConfigSvc))
	mux.HandleFunc("DELETE "+COMPANY_CONFIG_ROUTE_ID, handlers.DeleteCompanyConfig(companyConfigSvc))

	// Measurement type routes:
	mux.HandleFunc("GET "+MEASUREMENT_TYPE_ROUTE, handlers.GetMeasurementTypes(measurementTypeSvc))
	mux.HandleFunc("GET "+MEASUREMENT_TYPE_ROUTE_ALL, handlers.GetAllMeasurementTypes(measurementTypeSvc))
	mux.HandleFunc("POST "+MEASUREMENT_TYPE_ROUTE, handlers.PostMeasurementType(measurementTypeSvc))
	mux.HandleFunc("PATCH "+MEASUREMENT_TYPE_ROUTE_DEPRECATE, handlers.PatchDeprecateMeasurementType(measurementTypeSvc))

	// Swagger docs
	if enableSwagger {
		mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	}

	return mux
}
