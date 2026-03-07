// Package handlers TODO(@vinjar): add proper documentation.
package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetGateways TODO(@vinjar): add proper documentation.
func GetGateways(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the domain data from service layer
		data, err := svc.GetAll(ctx)
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

// PostGateway TODO(@vinjar): add proper documentation.
func PostGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateGatewayRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapGatewayDTOToDomain(payload)

		if err := svc.Create(ctx, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// PutGateway updates a gateway by its internal ID
func PutGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no gateway id found"), "bad request")
			return
		}

		payload, err := json.Decode[dto.CreateGatewayRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapGatewayDTOToDomain(payload)

		if err := svc.Update(ctx, id, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteGateway TODO(@vinjar): add proper documentation.
func DeleteGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no gateway id found"), "bad request")
			return
		}

		if err := svc.Delete(ctx, id); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
