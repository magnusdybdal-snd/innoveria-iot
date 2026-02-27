package handlers

import (
	"innoveria-iot/device-service/internal/domain"
	"net/http"
)

// GetAllSensorProfiles returns all available sensor profiles we provide
// This is not checked with authentication so every user can utilize this library
func GetAllSensorProfiles(svc domain.SensorProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
