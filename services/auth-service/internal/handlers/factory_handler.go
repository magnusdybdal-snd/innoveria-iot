package handlers

import (
	"errors"
	"fmt"
	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
	"net/http"

	"github.com/google/uuid"
)

// PostFactory handles factory creation requests.
//
// @Summary Register a new factory
// @Tags factories
// @Accept json
// @Produce json
// @Param body body dto.CreateNewFactory true "Factory payload"
// @Success 201 {object} dto.FactoryResponse
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /factories [post]
func PostFactory(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateNewFactory](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		factoryDomain := dto.MapCreateFactoryToDomain(payload)
		if factoryDomain.CompanyID == "" || factoryDomain.Name == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing required fields"), "company_id and name are required")
			return
		}

		// error check for invalid input
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

// GetAllFactories handles requests to fetch all factories.
//
// @Summary Get all factories
// @Tags factories
// @Produce json
// @Success 200 {object} dto.FactoryListResponse
// @Failure 500
// @Router /factories [get]
func GetAllFactories(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		factories, err := svc.GetAllFactories(ctx)
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

// GetOneFactory handles requests to fetch one factory by ID.
//
// @Summary Get one factory
// @Tags factories
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} dto.FactoryResponse
// @Failure 400
// @Failure 500
// @Router /factories/{id} [get]
func GetOneFactory(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		factoryID := r.PathValue("id")
		if factoryID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		factory, err := svc.GetOneFactory(ctx, factoryID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapFactoryFromDomain(factory)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
		}
	}
}

// DeleteFactory handles requests to delete a factory by ID.
//
// @Summary Delete factory
// @Tags factories
// @Param id path string true "id"
// @Success 204
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /factories/{id} [delete]
func DeleteFactory(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		factoryID := r.PathValue("id")
		if factoryID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("missing id path parameter"), "id is required")
			return
		}

		// Check for valid uuid
		if _, err := uuid.Parse(factoryID); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "invalid factory id (uuid)")
			return
		}

		if err := svc.DeleteFactory(ctx, factoryID); err != nil {
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
