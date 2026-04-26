package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"
)

// GetSensorMetrics returns the effective payload key mappings for a sensor.
// Checks per-sensor metrics first, falls back to profile-level payload schema.
// Returns an empty list if the sensor exists but is not yet configured.
// Returns 404 if the sensor EUI does not exist.
//
// @Summary		Get effective metrics for a sensor
// @Tags		sensor-metrics
// @Produce		json
// @Param		eui	path	string	true	"Device EUI"
// @Success		200	{object}	dto.SensorMetricListResponse
// @Failure		404
// @Failure		500
// @Router		/sensors/{eui}/metrics [get]
func GetSensorMetrics(svc domain.SensorMetricService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		eui := r.PathValue("eui")

		metrics, err := svc.GetEffectiveMetrics(ctx, eui)
		if err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "sensor not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapSensorMetricDomainToDTO(metrics)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PutSensorMetrics saves operator-defined metric labels for a configurable sensor. Admin only.
//
// @Summary		Save metric labels for a sensor
// @Tags		sensor-metrics
// @Accept		json
// @Param		eui		path	string							true	"Device EUI"
// @Param		body	body	dto.UpsertSensorMetricsRequest	true	"Metrics to save"
// @Success		204
// @Failure		400
// @Failure		401
// @Failure		403
// @Failure		404
// @Failure		422
// @Failure		500
// @Router		/sensors/{eui}/metrics [put]
func PutSensorMetrics(svc domain.SensorMetricService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}
		if !auth.IsAdmin() {
			json.HandleError(w, http.StatusForbidden, fmt.Errorf("forbidden"), "forbidden")
			return
		}

		eui := r.PathValue("eui")

		payload, err := json.Decode[dto.UpsertSensorMetricsRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if len(payload.Metrics) == 0 {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("metrics must not be empty"), "bad request")
			return
		}

		for i, m := range payload.Metrics {
			payload.Metrics[i].PayloadKey = strings.TrimSpace(m.PayloadKey)
			if payload.Metrics[i].PayloadKey == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("payload_key must not be empty"), "bad request")
				return
			}
			payload.Metrics[i].MeasurementType = strings.TrimSpace(m.MeasurementType)
			if payload.Metrics[i].MeasurementType == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("measurement_type must not be empty"), "bad request")
				return
			}
		}

		metrics := dto.MapUpsertMetricsRequestToDomain(payload)

		if err := svc.UpsertMetrics(ctx, eui, metrics); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "sensor not found")
				return
			}
			if errors.Is(err, domain.ErrInvalidMeasurementType) {
				json.HandleError(w, http.StatusUnprocessableEntity, err, "one or more measurement types do not exist in the vocabulary")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
