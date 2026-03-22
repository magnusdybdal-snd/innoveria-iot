package handlers

import (
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	"net/http"
)

// GetAllSensorProfiles returns all available sensor profiles.
//
// @Summary		List all sensor profiles
// @Description	Not authenticated — every user can access this list.
// @Tags		sensor-profiles
// @Produce		json
// @Success		200	{object}	dto.SensorProfileListResponse
// @Failure		400
// @Failure		500
// @Router		/sensor-profiles [get]
func GetAllSensorProfiles(svc domain.SensorProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := svc.GetAll(ctx)
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
