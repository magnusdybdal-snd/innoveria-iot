package handlers

import (
	"net/http"
	"strconv"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// PostSensorGroup creates a new sensor group.
//
// @Summary		Create a sensor group
// @Tags		sensor-groups
// @Accept		json
// @Param		body	body	dto.CreateSensorGroup	true	"Sensor group payload"
// @Success		201
// @Failure		400
// @Failure		500
// @Router		/sensor-groups [post]
func PostSensorGroup(svc domain.SensorGroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateSensorGroup](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapSensorGroupToDomain(payload)

		if err := svc.Create(ctx, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// GetAllSensorGroups returns all sensor groups.
//
// @Summary		List all sensor groups
// @Tags		sensor-groups
// @Produce		json
// @Param		limit	query	int	true	"Max number of groups to return"
// @Success		200	{object}	dto.SensorGroupListResponse
// @Failure		400
// @Failure		500
// @Router		/sensor-groups [get]
func GetAllSensorGroups(svc domain.SensorGroupService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		limitStr := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}
		resp, err := svc.GetAll(ctx, limit)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
		data := dto.MapSensorGroupToDTO(resp)

		if err := json.Encode(w, http.StatusOK, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
