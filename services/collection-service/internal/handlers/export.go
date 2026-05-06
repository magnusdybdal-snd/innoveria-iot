package handlers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"innoveria-iot/collection-service/internal/clients"
	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
)

// HandleExportMeasurements streams a CSV file of measurements for a device within a time range.
//
// @Summary		Export measurements as CSV
// @Tags		measurements
// @Produce		text/csv
// @Param		device_eui	query	string	true	"Device EUI"
// @Param		from		query	string	true	"Start time (RFC3339, e.g. 2024-01-01T00:00:00Z)"
// @Param		to			query	string	true	"End time (RFC3339, e.g. 2024-01-02T00:00:00Z)"
// @Success		200
// @Failure		400
// @Failure		401
// @Failure		500
// @Router		/export [get]
func HandleExportMeasurements(svc domain.MeasurementService, deviceClient *clients.DeviceClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		q := r.URL.Query()

		deviceEUI := q.Get("device_eui")
		if deviceEUI == "" {
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

		if !from.Before(to) {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("from is not before to"), "from must be before to")
			return
		}

		metrics, err := deviceClient.GetSensorMetrics(r.Context(), deviceEUI)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch sensor metrics")
			return
		}
		if len(metrics) == 0 {
			json.HandleError(w, http.StatusUnprocessableEntity, fmt.Errorf("no metrics configured for sensor %s", deviceEUI), "sensor has no metrics configured — configure metrics before exporting")
			return
		}

		measurements, err := svc.GetByTimeRange(r.Context(), auth.CompanyID, deviceEUI, from, to)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch measurements")
			return
		}

		filename := fmt.Sprintf("%s_%s_%s.csv", deviceEUI, from.Format(time.DateOnly), to.Format(time.DateOnly))
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)

		cw := csv.NewWriter(w)

		header := []string{"timestamp"}
		for _, m := range metrics {
			col := m.MeasurementType
			if m.Unit != nil && *m.Unit != "" {
				col = col + " (" + *m.Unit + ")"
			}
			header = append(header, col)
		}
		cw.Write(header) //nolint:errcheck

		for _, meas := range measurements {
			row := []string{meas.Timestamp.Format(time.RFC3339)}
			for _, m := range metrics {
				val := fmt.Sprintf("%v", meas.Payload[m.PayloadKey])
				row = append(row, val)
			}
			cw.Write(row) //nolint:errcheck
		}

		cw.Flush()
	}
}
