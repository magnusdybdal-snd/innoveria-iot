// Package handlers provides HTTP handler functions for the collection service API.
package handlers

import (
	"fmt"
	"net/http"
	"time"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/json"
)

// HandleMeasurementsByTimeRange returns measurements for a device within a time range.
//
// @Summary		Get measurements by time range
// @Tags		measurements
// @Produce		json
// @Param		device_eui	query	string	true	"Device EUI"
// @Param		from		query	string	true	"Start time (RFC3339, e.g. 2024-01-01T00:00:00Z)"
// @Param		to			query	string	true	"End time (RFC3339, e.g. 2024-01-02T00:00:00Z)"
// @Success		200	{array}		domain.SensorMeasurement
// @Failure		400
// @Failure		500
// @Router		/measurements [get]
func HandleMeasurementsByTimeRange(svc domain.MeasurementService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()

		// Check for all query parameters
		deviceStr := q.Get("device_eui")
		if deviceStr == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing device_eui"), "device_eui query parameter is required")
			return
		}

		fromStr := q.Get("from")
		if fromStr == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing from"), "from query parameter is required")
			return
		}

		toStr := q.Get("to")
		if toStr == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing to"), "to query parameter is required")
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
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("from is not before to"), "from must be before to")
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
