package handlers

import (
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

func GetSensors(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// 1. get the domain sensor data
		data, err := svc.GetAll(ctx)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		// 2. transform to sensor dto for response format
		resp := dto.MapSensorDomainToDTO(data)

		// 3. json encode
		if err := json.Encode(w, http.StatusOK, resp); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}
	}
}

func PostSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.Context()

	}
}
