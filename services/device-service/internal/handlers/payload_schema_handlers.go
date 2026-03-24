package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetPayloadSchemaDrafts returns ChirpStack profile IDs with unlabeled payload schema rows.
// Admin only — enforcement is handled at the API gateway level.
//
// @Summary		List profiles with unlabeled payload schema rows
// @Tags		payload-schema
// @Produce		json
// @Success		200	{object}	dto.DraftProfilesResponse
// @Failure		500
// @Router		/payload-schema/drafts [get]
func GetPayloadSchemaDrafts(svc domain.PayloadSchemaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		profileIDs, err := svc.GetDraftProfiles(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapDraftProfilesToDTO(profileIDs)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// GetPayloadSchemaByProfile returns all payload schema rows for a ChirpStack profile.
// Admin only — enforcement is handled at the API gateway level.
//
// @Summary		Get payload schema for a profile
// @Tags		payload-schema
// @Produce		json
// @Param		chirpstack_profile_id	path	string	true	"ChirpStack profile ID"
// @Success		200	{object}	dto.PayloadSchemaListResponse
// @Failure		400
// @Failure		500
// @Router		/payload-schema/{chirpstack_profile_id} [get]
func GetPayloadSchemaByProfile(svc domain.PayloadSchemaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

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

// PostDiscoverPayloadKeys creates draft payload schema rows for a profile from discovered keys.
// Admin only — enforcement is handled at the API gateway level.
//
// @Summary		Discover payload keys for a profile
// @Tags		payload-schema
// @Accept		json
// @Param		chirpstack_profile_id	path	string							true	"ChirpStack profile ID"
// @Param		body					body	dto.DiscoverPayloadKeysRequest	true	"Discovered payload keys"
// @Success		201
// @Failure		400
// @Failure		500
// @Router		/payload-schema/{chirpstack_profile_id}/discover [post]
func PostDiscoverPayloadKeys(svc domain.PayloadSchemaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		profileID := r.PathValue("chirpstack_profile_id")

		payload, err := json.Decode[dto.DiscoverPayloadKeysRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if len(payload.PayloadKeys) == 0 {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("payload_keys must not be empty"), "bad request")
			return
		}

		for i, key := range payload.PayloadKeys {
			payload.PayloadKeys[i] = strings.TrimSpace(key)
			if payload.PayloadKeys[i] == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("payload_keys must not contain empty strings"), "bad request")
				return
			}
		}

		if err := svc.DiscoverKeys(ctx, profileID, payload.PayloadKeys); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// PutPayloadSchemaLabels saves measurement type labels for a profile's payload schema rows.
// Admin only — enforcement is handled at the API gateway level.
//
// @Summary		Save payload schema labels for a profile
// @Tags		payload-schema
// @Accept		json
// @Param		chirpstack_profile_id	path	string								true	"ChirpStack profile ID"
// @Param		body					body	dto.SavePayloadSchemaLabelsRequest	true	"Labels to save"
// @Success		204
// @Failure		400
// @Failure		422
// @Failure		500
// @Router		/payload-schema/{chirpstack_profile_id} [put]
func PutPayloadSchemaLabels(svc domain.PayloadSchemaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

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

		for _, l := range payload.Labels {
			if strings.TrimSpace(l.PayloadKey) == "" {
				json.HandleError(w, http.StatusBadRequest, fmt.Errorf("payload_key must not be empty"), "bad request")
				return
			}
			if strings.TrimSpace(l.MeasurementType) == "" {
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
