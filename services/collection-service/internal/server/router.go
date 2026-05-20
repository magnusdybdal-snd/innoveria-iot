// Package server configures and runs the collection service HTTP server.
package server

import (
	"net/http"

	_ "innoveria-iot/collection-service/docs"
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers"
	"innoveria-iot/pkg/middleware"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter registers all HTTP routes and returns the configured ServeMux.
func NewRouter(svc domain.MeasurementService, tenantMappingSvc domain.TenantMappingService, exportSvc domain.ExportService, enableSwagger bool) *http.ServeMux {
	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("GET /", handlers.Root)

	// Measurements
	mux.HandleFunc("GET "+LATEST_MEASUREMENT, handlers.HandleLatestMeasurement(svc))
	mux.HandleFunc("GET "+MEASUREMENTS_BY_TIME, handlers.HandleMeasurementsByTimeRange(svc))

	// Payload tags
	mux.HandleFunc("GET "+PAYLOAD_TAGS, middleware.AdminGuard(handlers.HandlePayloadTags(svc)))

	// Export
	mux.HandleFunc("GET "+EXPORT_MEASUREMENTS, handlers.HandleExportMeasurements(exportSvc))

	// Company config routes:
	mux.HandleFunc("POST "+COMPANY_CONFIG_ROUTE, handlers.PostTenantMapping(tenantMappingSvc))

	// Swagger
	if enableSwagger {
		mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	}

	return mux
}
