// Package handlers TODO(@Magnus Dybdal): add proper documentation.
package handlers

import (
	"net/http"
	"time"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/json"
)

// HandleMeasurementsByTimeRange TODO(@Magnus Dybdal): add proper documentation.
func HandleMeasurementsByTimeRange(svc domain.MeasurementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		// Check for all query parameters
		deviceStr := q.Get("device_eui")
		if deviceStr == "" {
			json.HandleError(w, http.StatusBadRequest, nil, "device_eui query parameter is required")
			return
		}

		fromStr := q.Get("from")
		if fromStr == "" {
			json.HandleError(w, http.StatusBadRequest, nil, "from query parameter is required")
			return
		}

		toStr := q.Get("to")
		if toStr == "" {
			json.HandleError(w, http.StatusBadRequest, nil, "to query parameter is required")
			return
		}

		// Parse RFC3339 timestamps
		from, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "from must be a valid RFC3339 timestamp")
			return
		}

		to, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "to must be a valid RFC3339 timestamp")
			return
		}

		// Sanity check: timestamp from must come before timestamp to
		if !from.Before(to) {
			json.HandleError(w, http.StatusBadRequest, nil, "timestamp error: from must be before to")
			return
		}

		// Delegate to the service
		measurements, err := svc.GetByTimeRange(r.Context(), deviceStr, from, to)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch measurements")
			return
		}

		// Return the slice as json
		if err := json.Encode(w, http.StatusOK, measurements); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to encode response")
		}
	}
}
