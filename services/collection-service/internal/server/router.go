// Package server provides the HTTP server and routing for the device service API.
package server

import (
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers"
	"net/http"
)

// NewRouter TODO(@Magnus Dybdal): add proper documentation.
func NewRouter(svc domain.MeasurementService) *http.ServeMux {
	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET "+LATEST_MEASUREMENT, handlers.HandleLatestMeasurement(svc))
	mux.HandleFunc("GET "+MEASUREMENTS_BY_TIME, handlers.HandleMeasurementsByTimeRange(svc))

	return mux
}
