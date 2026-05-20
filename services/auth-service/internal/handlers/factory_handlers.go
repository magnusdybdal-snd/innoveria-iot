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

// PostFactory handles factory creation requests.
// The factory is created under the company of the authenticated user.
//
// @Summary Register a new factory
// @Tags factories
// @Accept json
// @Produce json
// @Param body body dto.CreateNewFactory true "Factory payload"
// @Success 201 {object} dto.FactoryResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factories [post]
func PostFactory(svc domain.FactoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		payload, err := json.Decode[dto.CreateNewFactory](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Name = strings.TrimSpace(payload.Name)
		payload.Address = strings.TrimSpace(payload.Address)

		if payload.Name == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "name is required")
			return
		}

		factoryDomain := dto.MapCreateFactoryToDomain(payload, auth.CompanyID)

		factoryResp, err := svc.RegisterFactory(ctx, factoryDomain)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrCompanyNotFound):
				json.HandleError(w, http.StatusNotFound, err, "company not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapFactoryFromDomain(factoryResp)
		if err := json.Encode(w, http.StatusCreated, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetAllFactories handles requests to fetch all factories belonging to the authenticated user's company.
//
// @Summary Get all factories
// @Tags factories
// @Produce json
// @Success 200 {object} dto.FactoryListResponse
// @Failure 401
// @Failure 500
// @Router /factories [get]
func GetAllFactories(svc domain.FactoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		factories, err := svc.GetAllFactories(ctx, auth.CompanyID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapFactoryListFromDomain(factories)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// GetOneFactory handles requests to fetch one factory by ID, scoped to the authenticated user's company.
//
// @Summary Get one factory
// @Tags factories
// @Produce json
// @Param id path string true "Factory ID"
// @Success 200 {object} dto.FactoryResponse
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factories/{id} [get]
func GetOneFactory(svc domain.FactoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		factoryID := r.PathValue("id")
		if factoryID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		if _, err := uuid.Parse(factoryID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory id (uuid)")
			return
		}

		factory, err := svc.GetOneFactory(ctx, auth.CompanyID, factoryID)
		if err != nil {
			switch {
			case errors.Is(err, domain.ErrFactoryNotFound):
				json.HandleError(w, http.StatusNotFound, err, "factory not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		resp := dto.MapFactoryFromDomain(factory)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteFactory handles requests to delete a factory by ID, scoped to the authenticated user's company.
//
// @Summary Delete factory
// @Tags factories
// @Param id path string true "Factory ID"
// @Success 204
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /factories/{id} [delete]
func DeleteFactory(svc domain.FactoryService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		factoryID := r.PathValue("id")
		if factoryID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		if _, err := uuid.Parse(factoryID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory id (uuid)")
			return
		}

		if err := svc.DeleteFactory(ctx, auth.CompanyID, factoryID); err != nil {
			switch {
			case errors.Is(err, domain.ErrFactoryNotFound):
				json.HandleError(w, http.StatusNotFound, err, "factory not found")
			default:
				json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
