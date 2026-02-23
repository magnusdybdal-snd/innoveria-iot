package handlers

import (
	"net/http"

	"innoveria-iot/collection-service/internal/domain"
)

func ReplaceName(svc domain.MeasurementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: Implementation
	}
}
