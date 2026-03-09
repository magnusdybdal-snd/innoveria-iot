// Package server configures and runs the collection service HTTP server.
package server

import (
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers"
	"net/http"

	_ "innoveria-iot/collection-service/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter registers all HTTP routes and returns the configured ServeMux.
func NewRouter(svc domain.MeasurementService) *http.ServeMux {
	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET "+LATEST_MEASUREMENT, handlers.HandleLatestMeasurement(svc))
	mux.HandleFunc("GET "+MEASUREMENTS_BY_TIME, handlers.HandleMeasurementsByTimeRange(svc))

	// Swagger
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	return mux
}
