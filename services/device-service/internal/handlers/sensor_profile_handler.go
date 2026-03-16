package handlers

import (
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	"net/http"
	"strconv"
)

// GetAllSensorProfiles returns all available sensor profiles.
//
// @Summary		List all sensor profiles
// @Description	Not authenticated — every user can access this list.
// @Tags		sensor-profiles
// @Produce		json
// @Param		limit	query	int	true	"Max number of profiles to return"
// @Success		200	{object}	dto.SensorProfileListResponse
// @Failure		400
// @Failure		500
// @Router		/sensor-profiles [get]
func GetAllSensorProfiles(svc domain.SensorProfileService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
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
