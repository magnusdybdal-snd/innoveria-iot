package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetSensorProfileConfig returns the configuration for a Chirpstack device profile.
// If no configuration exists, a default with configurable_schema = false is returned.
//
// @Summary		Get sensor profile config
// @Tags		sensor-profile-config
// @Produce		json
// @Param		profile_id	path		string	true	"Chirpstack profile ID"
// @Success		200			{object}	dto.SensorProfileConfigResponse
// @Failure		500
// @Router		/sensor-profile-config/{profile_id} [get]
func GetSensorProfileConfig(svc domain.SensorProfileConfigService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		profileID := r.PathValue("profile_id")
		if profileID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("profile_id is required"), "bad request")
			return
		}

		data, err := svc.Get(ctx, profileID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapSensorProfileConfigDomainToDTO(data)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PatchSensorProfileConfig updates the configuration for a Chirpstack device profile.
//
// @Summary		Update sensor profile config
// @Tags		sensor-profile-config
// @Accept		json
// @Param		profile_id	path	string								true	"Chirpstack profile ID"
// @Param		body		body	dto.PatchSensorProfileConfigRequest	true	"Sensor profile config payload"
// @Success		204
// @Failure		400
// @Failure		500
// @Router		/sensor-profile-config/{profile_id} [patch]
func PatchSensorProfileConfig(svc domain.SensorProfileConfigService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		profileID := r.PathValue("profile_id")
		if profileID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("profile_id is required"), "bad request")
			return
		}

		payload, err := json.Decode[dto.PatchSensorProfileConfigRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapPatchSensorProfileConfigDTOToDomain(profileID, payload)

		if err := svc.Upsert(ctx, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
