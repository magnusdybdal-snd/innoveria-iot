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

// GetMeasurementTypes returns all active (non-deprecated) measurement types.
//
// @Summary		List active measurement types
// @Tags		measurement-types
// @Produce		json
// @Success		200	{object}	dto.MeasurementTypeListResponse
// @Failure		500
// @Router		/measurement-types [get]
func GetMeasurementTypes(svc domain.MeasurementTypeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := svc.ListActive(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapMeasurementTypeDomainToDTO(data)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// GetAllMeasurementTypes returns all measurement types including deprecated. For admin use only.
//
// @Summary		List all measurement types including deprecated
// @Tags		measurement-types
// @Produce		json
// @Success		200	{object}	dto.MeasurementTypeListResponse
// @Failure		500
// @Router		/measurement-types/all [get]
func GetAllMeasurementTypes(svc domain.MeasurementTypeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := svc.ListAll(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapMeasurementTypeDomainToDTO(data)

		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PostMeasurementType creates a new measurement type.
//
// @Summary		Create a measurement type
// @Tags		measurement-types
// @Accept		json
// @Param		body	body	dto.CreateMeasurementTypeRequest	true	"Measurement type payload"
// @Success		201
// @Failure		400
// @Failure		409
// @Failure		500
// @Router		/measurement-types [post]
func PostMeasurementType(svc domain.MeasurementTypeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateMeasurementTypeRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Slug = strings.TrimSpace(payload.Slug)
		payload.DisplayName = strings.TrimSpace(payload.DisplayName)

		if payload.Slug == "" || payload.DisplayName == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("slug and display_name are required"), "bad request")
			return
		}

		data := dto.MapCreateMeasurementTypeDTOToDomain(payload)

		if err := svc.Create(ctx, data); err != nil {
			if errors.Is(err, domain.ErrAlreadyExists) {
				json.HandleError(w, http.StatusConflict, err, "measurement type with this slug already exists")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// PatchDeprecateMeasurementType marks a measurement type as deprecated.
//
// @Summary		Deprecate a measurement type
// @Tags		measurement-types
// @Param		slug	path	string	true	"Measurement type slug"
// @Success		204
// @Failure		400
// @Failure		404
// @Failure		500
// @Router		/measurement-types/{slug}/deprecate [patch]
func PatchDeprecateMeasurementType(svc domain.MeasurementTypeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		slug := r.PathValue("slug")
		if slug == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("slug is required"), "bad request")
			return
		}

		if err := svc.Deprecate(ctx, slug); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "measurement type not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
