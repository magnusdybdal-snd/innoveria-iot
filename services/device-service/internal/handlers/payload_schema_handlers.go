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

// GetPayloadSchemaByProfile returns all payload schema rows for a ChirpStack profile. Admin only.
//
// @Summary		Get payload schema for a profile
// @Tags		payload-schema
// @Produce		json
// @Param		chirpstack_profile_id	path	string	true	"ChirpStack profile ID"
// @Success		200	{object}	dto.PayloadSchemaListResponse
// @Failure		400
// @Failure		401
// @Failure		403
// @Failure		500
// @Router		/payload-schema/{chirpstack_profile_id} [get]
func GetPayloadSchemaByProfile(svc domain.PayloadSchemaService) http.HandlerFunc {
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

		profileID := r.PathValue("chirpstack_profile_id")

		schemas, err := svc.GetByProfile(ctx, profileID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapPayloadSchemaDomainToDTO(schemas)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PutPayloadSchemaLabels saves measurement type labels for a profile's payload schema rows. Admin only.
//
// @Summary		Save payload schema labels for a profile
// @Tags		payload-schema
// @Accept		json
// @Param		chirpstack_profile_id	path	string								true	"ChirpStack profile ID"
// @Param		body					body	dto.SavePayloadSchemaLabelsRequest	true	"Labels to save"
// @Success		204
// @Failure		400
// @Failure		401
// @Failure		403
// @Failure		422
// @Failure		500
// @Router		/payload-schema/{chirpstack_profile_id} [put]
func PutPayloadSchemaLabels(svc domain.PayloadSchemaService) http.HandlerFunc {
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

		profileID := r.PathValue("chirpstack_profile_id")

		payload, err := json.Decode[dto.SavePayloadSchemaLabelsRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if len(payload.Labels) == 0 {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("labels must not be empty"), "bad request")
			return
		}

		seen := make(map[string]struct{}, len(payload.Labels))
		for i, label := range payload.Labels {
			payload.Labels[i].PayloadKey = strings.TrimSpace(label.PayloadKey)
			if payload.Labels[i].PayloadKey == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("payload_key must not be empty"), "bad request")
				return
			}

			if _, dup := seen[payload.Labels[i].PayloadKey]; dup {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("duplicate payload_key: %q", payload.Labels[i].PayloadKey), "bad request")
				return
			}
			seen[payload.Labels[i].PayloadKey] = struct{}{}

			payload.Labels[i].MeasurementType = strings.TrimSpace(label.MeasurementType)
			if payload.Labels[i].MeasurementType == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("measurement_type must not be empty"), "bad request")
				return
			}
		}

		schemas := dto.MapSaveLabelsRequestToDomain(profileID, payload)

		if err := svc.SaveLabels(ctx, schemas); err != nil {
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
