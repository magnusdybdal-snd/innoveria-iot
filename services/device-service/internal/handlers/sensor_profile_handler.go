package handlers

import (
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	"net/http"
	"strconv"
)

// GetAllSensorProfiles returns all available sensor profiles we provide
// This is not checked with authentication so every user can utilize this library
func GetAllSensorProfiles(svc domain.SensorProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		data, err := svc.GetAll(ctx, limit)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
		resp := dto.MapSensorProfileDomainToDTO(data)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// TODO: Handle custom device profiles
