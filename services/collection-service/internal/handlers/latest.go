package handlers

import (
	"net/http"

	"innoveria-iot/collection-service/internal/domain"
)

// TODO: fix latest sensor read
func HandleLatestSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
