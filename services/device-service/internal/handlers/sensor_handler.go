package handlers

import (
	"fmt"
	"net/http"

	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/handlers/dto"
	"innoveria-iot/pkg/json"
)

// GetSensors returns all GetSensors
//
// @Summary 	Lists all sensors.
// @Tags 		sensors
// @Produce 	json
// @Success 	200 {object} dto.SensorListResponse
// @Failure 	400
// @Failure 	500
// @Router 		/sensors [get]
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

// PostSensor registers a new sensor in ChirpStack and the database.
//
// @Summary		Create a sensor
// @Tags 		sensors
// @Accept		json
// @Param		body   body   dto.CreateSensorRequest   true   "Sensor payload"
// @Success		201
// @Failure		400
// @Failure 	500
// @Router		/sensors [post]
func PostSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := json.Decode[dto.CreateSensorRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		data := dto.MapCreateSensorDTOToDomain(payload)

		if err := svc.Create(ctx, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusCreated)

	}
}

// PatchSensor updates a sensor by its internal ID
//
// @Summary 	Update a sensor
// @Tags		sensors
// @Accept		json
// @Param		id     path   string                    true   "Sensor ID"
// @Param		body   body   dto.UpdateSensorRequest   true   "Update payload"
// @Success		204
// @Failure		400
// @Failure		500
// @Router		/sensors/{id} [patch]
func PatchSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no sensor id found"), "bad request")
			return
		}

		payload, err := json.Decode[dto.UpdateSensorRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if payload.Name == nil && payload.Description == nil && payload.FactoryID == nil && payload.FactoryAreaID == nil && payload.ChirpstackProfileID == nil && payload.ProductionResource == nil {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no fields provided"), "bad request")
			return
		}

		data := dto.MapUpdateSensorDTOToDomain(payload)

		if err := svc.Update(ctx, id, data); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteSensor deletes a sensor by its internal ID.
//
// @Summary		Delete a sensor
// @Tags 		sensors
// @Param		id		path	string		true	"SensorID"
// @Success		204
// @Failure		400
// @Failure		500
// @Router		/sensors/{id} [delete]
func DeleteSensor(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no sensor id found"), "bad request")
			return
		}

		if err := svc.Delete(ctx, id); err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
