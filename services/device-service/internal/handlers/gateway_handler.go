package handlers

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetGateway returns a json with all available gateways and their status from chirpstack
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

func PostGateway(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateGatewayRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapGatewayDTOToDomain(payload)

		if err := svc.Create(ctx, data, payload.CompanyId); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if err := json.Encode(w, http.StatusCreated, map[string]string{}); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
