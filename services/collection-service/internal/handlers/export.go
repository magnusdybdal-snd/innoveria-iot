package handlers

import (
	"encoding/csv"
	"fmt"
	"math"
	"net/http"
	"strconv"
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
// @Param		timezone	query	string	false	"IANA timezone name (e.g. Europe/Oslo). Defaults to UTC."
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

		loc := time.UTC
		if tzName := q.Get("timezone"); tzName != "" {
			if parsed, err := time.LoadLocation(tzName); err == nil {
				loc = parsed
			}
		}

		metrics, err := deviceClient.GetSensorMetrics(r.Context(), deviceEUI)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch sensor metrics")
			return
		}

		measurements, err := svc.GetByTimeRange(r.Context(), auth.CompanyID, deviceEUI, from, to)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch measurements")
			return
		}

		// Build a payloadKey → column header map from configured metrics.
		metricHeader := make(map[string]string, len(metrics))
		for _, m := range metrics {
			col := m.MeasurementType
			if m.Unit != nil && *m.Unit != "" {
				col = col + " (" + *m.Unit + ")"
			}
			metricHeader[m.PayloadKey] = col
		}

		// Collect the ordered column set: configured metric keys first (preserving
		// device-service order), then any unconfigured keys found in the measurements.
		seen := make(map[string]bool, len(metrics))
		columns := make([]string, 0, len(metrics))
		for _, m := range metrics {
			seen[m.PayloadKey] = true
			columns = append(columns, m.PayloadKey)
		}
		for _, meas := range measurements {
			for key := range meas.Payload {
				if !seen[key] {
					seen[key] = true
					columns = append(columns, key)
				}
			}
		}

		filename := fmt.Sprintf("%s_%s_%s.csv", deviceEUI, from.Format(time.DateOnly), to.Format(time.DateOnly))
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)

		cw := csv.NewWriter(w)

		tzLabel := loc.String()

		header := []string{"date", "time (" + tzLabel + ")"}
		for _, key := range columns {
			if col, ok := metricHeader[key]; ok {
				header = append(header, col)
			} else {
				header = append(header, key)
			}
		}
		cw.Write(header) //nolint:errcheck

		for _, meas := range measurements {
			t := meas.Timestamp.In(loc)
			row := []string{
				t.Format(time.DateOnly),
				t.Format("15:04:05"),
			}
			for _, key := range columns {
				row = append(row, formatPayloadValue(meas.Payload[key]))
			}
			cw.Write(row) //nolint:errcheck
		}

		cw.Flush()
	}
}

// formatPayloadValue formats a JSON payload value for CSV output.
// Float64 values are rounded to 4 decimal places and trailing zeros are stripped.
// All other types use their default string representation.
func formatPayloadValue(v any) string {
	f, ok := v.(float64)
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	rounded := math.Round(f*10000) / 10000
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}
