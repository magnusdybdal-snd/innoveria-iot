package handlers

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/pkg/json"
)

func GetGateways(svc domain.GatewayService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data, err := svc.GetAll(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		if err := json.Encode(w, http.StatusOK, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}
