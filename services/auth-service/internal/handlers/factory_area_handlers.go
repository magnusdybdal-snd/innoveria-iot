package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/auth-service/internal/domain"
	"innoveria-iot/auth-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"

	"github.com/google/uuid"
)

// PostFactoryArea handles factory area creation requests.
//
// @Summary Register a new factory area
// @Tags factory-areas
// @Accept json
// @Produce json
// @Param body body dto.CreateNewFactoryArea true "Factory area payload"
// @Success 201 {object} dto.FactoryAreaResponse
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /factory-areas [post]
func PostFactoryArea(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateNewFactoryArea](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
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

// GetAllFactoryAreas handles requests to fetch all factory areas.
//
// @Summary Get all factory areas
// @Tags factory-areas
// @Produce json
// @Success 200 {object} dto.FactoryAreaListResponse
// @Failure 500
// @Router /factory-areas [get]
func GetAllFactoryAreas(svc domain.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		areas, err := svc.GetAllFactoryAreas(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		resp := dto.MapFactoryAreaListFromDomain(areas)
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server")
		}
	}
}
