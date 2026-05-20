// Package handlers implements the HTTP handlers for the device service.
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

	"github.com/google/uuid"
)

// GetGateways returns all gateways.
//
// @Summary		List all gateways
// @Tags		gateways
// @Produce		json
// @Success		200	{object}	dto.GatewayListResponse
// @Failure		401
// @Failure		500
// @Router		/gateways [get]
func GetGateways(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		// Get the domain data from service layer
		data, err := svc.GetAll(ctx, auth.CompanyID)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		// Map device gateway domain to data transfer object
		resp := dto.MapGatewayDomainToDTO(data)

		// Send dto json over http
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

// PostGateway creates a new gateway.
//
// @Summary		Create a gateway
// @Tags		gateways
// @Accept		json
// @Param		body	body	dto.CreateGatewayRequest	true	"Gateway payload"
// @Success		201
// @Failure		400
// @Failure		401
// @Failure		500
// @Router		/gateways [post]
func PostGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		payload, err := json.Decode[dto.CreateGatewayRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.GatewayEUI = strings.TrimSpace(payload.GatewayEUI)
		payload.Name = strings.TrimSpace(payload.Name)
		payload.FactoryID = strings.TrimSpace(payload.FactoryID)
		payload.FactoryAreaID = strings.TrimSpace(payload.FactoryAreaID)

		if payload.GatewayEUI == "" || payload.Name == "" || payload.FactoryID == "" || payload.FactoryAreaID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("gateway_eui, name, factory_id and factory_area_id are required"), "bad request")
			return
		}

		data := dto.MapGatewayDTOToDomain(payload, auth.CompanyID)

		if err := svc.Create(ctx, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// PatchGateway updates a gateway by its internal ID.
//
// @Summary		Update a gateway
// @Tags		gateways
// @Accept		json
// @Param		id		path	string						true	"Gateway ID"
// @Param		body	body	dto.UpdateGatewayRequest	true	"Update payload"
// @Success		204
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router		/gateways/{id} [patch]
func PatchGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no gateway id found"), "bad request")
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload, err := json.Decode[dto.UpdateGatewayRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if payload.Name == nil && payload.Description == nil && payload.FactoryID == nil && payload.FactoryAreaID == nil {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no fields provided"), "bad request")
			return
		}

		for _, field := range []*string{payload.FactoryID, payload.FactoryAreaID} {
			if field != nil {
				if _, err := uuid.Parse(*field); err != nil {
					json.HandleError(w, http.StatusBadRequest, err, "bad request")
					return
				}
			}
		}

		data := dto.MapUpdateGatewayDTOToDomain(payload)

		if err := svc.Update(ctx, auth.CompanyID, id, data); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "gateway not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteGateway deletes a gateway by its internal ID.
//
// @Summary		Delete a gateway
// @Tags		gateways
// @Param		id	path	string	true	"Gateway ID"
// @Success		204
// @Failure		400
// @Failure		401
// @Failure		404
// @Failure		500
// @Router		/gateways/{id} [delete]
func DeleteGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusUnauthorized, err, "unauthorized")
			return
		}

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no gateway id found"), "bad request")
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if err := svc.Delete(ctx, auth.CompanyID, id); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "gateway not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
