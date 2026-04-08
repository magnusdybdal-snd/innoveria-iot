package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/collection-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// HandlePayloadTags returns the unique payload tags for a device
//
// @Summary		Get payload tags
// @Tags		payload tags
// @Produce		json
// @Param		device_eui	query	string	true	"DeviceEUI"
// @Success		200		{object}	dto.PayloadTagsResponse
// @Failure		400
// @Failure     500
// @Router		/collection/payload-tags [get]
func HandlePayloadTags(svc domain.MeasurementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract the device eui query and check that it is not empty
		deviceEUI := r.URL.Query().Get("device_eui")
		if deviceEUI == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing device_eui"), "device_eui query parameter is required")
			return
		}

		// Get the payload tags for the device
		tags, err := svc.GetPayloadKeys(r.Context(), deviceEUI)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch payload tags")
			return
		}

		// Encode payload tags as a JSON response
		if err := json.Encode(w, http.StatusOK, dto.PayloadTagsResponse{Keys: tags}); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
