package handlers

import (
	"net/http"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/json"
)

func HandleLatestSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract the device eui query and check that it is not empty
		deviceEUI := r.URL.Query().Get("device_eui")
		if deviceEUI == "" {
			json.HandleError(w, http.StatusBadRequest, nil, "device_eui query parameter is required")
			return
		}

		// Get the latest measurement
		measurement, err := svc.GetLatest(r.Context(), deviceEUI)
		if err != nil {
			// TODO: Any db error is now sent as 404, need more specific from repo
			json.HandleError(w, http.StatusNotFound, err, "no measurement found for device")
			return
		}

		// Encode measurement as a JSON response
		if err := json.Encode(w, http.StatusOK, measurement); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
