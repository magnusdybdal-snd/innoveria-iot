package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/authctx"
	"innoveria-iot/pkg/json"

	"github.com/google/uuid"
)

// PostFactoryArea handles factory area creation requests.
// The factory must belong to the authenticated user's company.
//
// @Summary Register a new factory area
// @Tags factory-areas
// @Accept json
// @Produce json
// @Param body body dto.CreateNewFactoryArea true "Factory area payload"
// @Success 201 {object} dto.FactoryAreaResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factory-areas [post]
func PostFactoryArea(svc domain.FactoryAreaService, factorySvc domain.FactoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		payload, err := json.Decode[dto.CreateNewFactoryArea](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.FactoryID = strings.TrimSpace(payload.FactoryID)
		payload.Name = strings.TrimSpace(payload.Name)
		if payload.Description != nil {
			description := strings.TrimSpace(*payload.Description)
			payload.Description = &description
		}

		factoryAreaDomain := dto.MapCreateFactoryAreaToDomain(payload)
		if factoryAreaDomain.FactoryID == "" || factoryAreaDomain.Name == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "factory_id and name are required")
			return
		}

		if _, err := uuid.Parse(factoryAreaDomain.FactoryID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory_id (uuid)")
			return
		}

		if _, err := factorySvc.GetOneFactory(ctx, auth.CompanyID, factoryAreaDomain.FactoryID); err != nil {
			if errors.Is(err, domain.ErrFactoryNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "factory not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		factoryAreaResp, err := svc.RegisterFactoryArea(ctx, factoryAreaDomain)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrFactoryNotFound):
				json.HandleError(w, http.StatusNotFound, err, "factory not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapFactoryAreaFromDomain(factoryAreaResp)
		if err := json.Encode(w, http.StatusCreated, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetAllFactoryAreas handles requests to fetch all factory areas for a given factory,
// scoped to the authenticated user's company.
//
// @Summary Get all factory areas
// @Tags factory-areas
// @Produce json
// @Param factory_id query string true "Factory ID"
// @Success 200 {object} dto.FactoryAreaListResponse
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /factory-areas [get]
func GetAllFactoryAreas(svc domain.FactoryAreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		factoryID := r.URL.Query().Get("factory_id")
		if factoryID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing factory_id query parameter"), "factory_id is required")
			return
		}
		if _, err := uuid.Parse(factoryID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory_id (uuid)")
			return
		}

		areas, err := svc.GetAllFactoryAreas(ctx, auth.CompanyID, factoryID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapFactoryAreaListFromDomain(areas)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetOneFactoryArea handles requests to fetch one factory area by ID,
// scoped to the authenticated user's company.
//
// @Summary Get one factory area
// @Tags factory-areas
// @Produce json
// @Param id path string true "Factory area ID"
// @Success 200 {object} dto.FactoryAreaResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factory-areas/{id} [get]
func GetOneFactoryArea(svc domain.FactoryAreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		areaID := r.PathValue("id")
		if areaID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		if _, err := uuid.Parse(areaID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory area id (uuid)")
			return
		}

		area, err := svc.GetOneFactoryArea(ctx, auth.CompanyID, areaID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrFactoryAreaNotFound):
				json.HandleError(w, http.StatusNotFound, err, "factory area not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapFactoryAreaFromDomain(area)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteFactoryArea handles requests to delete a factory area by ID,
// scoped to the authenticated user's company.
//
// @Summary Delete factory area
// @Tags factory-areas
// @Param id path string true "Factory area ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factory-areas/{id} [delete]
func DeleteFactoryArea(svc domain.FactoryAreaService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		areaID := r.PathValue("id")
		if areaID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		if _, err := uuid.Parse(areaID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory area id (uuid)")
			return
		}

		if err := svc.DeleteFactoryArea(ctx, auth.CompanyID, areaID); err != nil {
			switch {
			case errors.Is(err, domain.ErrFactoryAreaNotFound):
				json.HandleError(w, http.StatusNotFound, err, "factory area not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
