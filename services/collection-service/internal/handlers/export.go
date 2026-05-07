package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"time"

	"innoveria-iot/collection-service/internal/domain"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/csvwriter"
	"innoveria-iot/pkg/json"
)

var deviceEUIRegex = regexp.MustCompile(`^[0-9a-fA-F]{16}$`)

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
func HandleExportMeasurements(exportSvc domain.ExportService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		q := r.URL.Query()

		deviceEUI := q.Get("device_eui")
		if !deviceEUIRegex.MatchString(deviceEUI) {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("invalid device_eui"), "device_eui must be a 16-character hex string")
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
			} else {
				slog.WarnContext(r.Context(), "unrecognised timezone, falling back to UTC", "timezone", tzName)
			}
		}

		data, err := exportSvc.GetExportData(r.Context(), auth.CompanyID, deviceEUI, from, to)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "failed to fetch export data")
			return
		}

		filename := fmt.Sprintf("%s_%s_%s.csv", deviceEUI, from.Format(time.DateOnly), to.Format(time.DateOnly))
		w.Header().Set("Content-Type", "text/csv")
		w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)

		columns := make([]csvwriter.Column, len(data.Columns))
		for i, col := range data.Columns {
			columns[i] = csvwriter.Column{PayloadKey: col.PayloadKey, Header: col.Header}
		}

		rows := make([]csvwriter.Row, len(data.Measurements))
		for i, meas := range data.Measurements {
			rows[i] = csvwriter.Row{Timestamp: meas.Timestamp, Payload: meas.Payload}
		}

		if err := csvwriter.Write(w, columns, rows, loc); err != nil {
			slog.ErrorContext(r.Context(), "csv write failed", "device_eui", deviceEUI, "error", err)
		}
	}
}
