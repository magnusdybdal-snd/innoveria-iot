package server

import (
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers"
	"net/http"
)

func NewRouter(svc domain.MeasurementService) *http.ServeMux {
	mux := http.NewServeMux()

	// Routes:
	mux.HandleFunc("GET /", handlers.Root)
	mux.HandleFunc("GET /latest", handlers.HandleLatestMeasurement(svc))
	mux.HandleFunc("GET /measurements", handlers.HandleMeasurementsByTimeRange(svc))
	mux.HandleFunc("GET /api/v1/collection/latest", handlers.HelloProxy) // test for proxy

	return mux
}
