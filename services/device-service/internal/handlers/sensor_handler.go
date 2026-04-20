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

// GetSensors returns all sensors, optionally filtered by production resource.
//
// @Summary 	Lists all sensors.
// @Tags 		sensors
// @Produce 	json
// @Param       production_resource_id    query    string    false    "Production resource"
// @Success 	200 {object} dto.SensorListResponse
// @Failure 	400
// @Failure 	500
// @Router 		/sensors [get]
func GetSensors(svc domain.SensorService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		var data []domain.Sensor

		productionResourceID := r.URL.Query().Get("production_resource_id")

		if productionResourceID != "" {
			if _, parseErr := uuid.Parse(productionResourceID); parseErr != nil {
				json.HandleError(w, http.StatusBadRequest, parseErr, "bad request")
				return
			}
			data, err = svc.GetByProductionResourceID(ctx, auth.CompanyID, productionResourceID)
		} else {
			data, err = svc.GetAll(ctx, auth.CompanyID)
		}

		// error check for both paths above in if/else
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		// Transform to sensor dto for response format
		resp := dto.MapSensorDomainToDTO(data)

		// Json encode
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
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		payload, err := json.Decode[dto.CreateSensorRequest](r)
		if err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		payload.Name = strings.TrimSpace(payload.Name)
		payload.DeviceEUI = strings.TrimSpace(payload.DeviceEUI)
		payload.AppKey = strings.TrimSpace(payload.AppKey)
		payload.ChirpstackProfileID = strings.TrimSpace(payload.ChirpstackProfileID)
		payload.FactoryID = strings.TrimSpace(payload.FactoryID)
		payload.FactoryAreaID = strings.TrimSpace(payload.FactoryAreaID)

		if payload.Name == "" || payload.DeviceEUI == "" || payload.AppKey == "" || payload.ChirpstackProfileID == "" || payload.FactoryID == "" || payload.FactoryAreaID == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("name, device_eui, app_key, device_profile_id, factory_id and factory_area_id are required"), "bad request")
			return
		}

		data := dto.MapCreateSensorDTOToDomain(payload, auth.CompanyID)

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
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no sensor id found"), "bad request")
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
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

		for _, field := range []*string{payload.FactoryID, payload.FactoryAreaID, payload.ChirpstackProfileID} {
			if field != nil {
				if _, err := uuid.Parse(*field); err != nil {
					json.HandleError(w, http.StatusBadRequest, err, "bad request")
					return
				}
			}
		}

		data := dto.MapUpdateSensorDTOToDomain(payload)

		if err := svc.Update(ctx, auth.CompanyID, id, data); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "sensor not found")
				return
			}
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
		auth, err := authctx.FromRequest(r)
		if err != nil {
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		id := r.PathValue("id")
		if id == "" {
			json.HandleError(w, http.StatusBadRequest, fmt.Errorf("no sensor id found"), "bad request")
			return
		}

		if _, err := uuid.Parse(id); err != nil {
			json.HandleError(w, http.StatusBadRequest, err, "bad request")
			return
		}

		if err := svc.Delete(ctx, auth.CompanyID, id); err != nil {
			if errors.Is(err, domain.ErrNotFound) {
				json.HandleError(w, http.StatusNotFound, err, "sensor not found")
				return
			}
			json.HandleError(w, http.StatusInternalServerError, err, "internal server error")
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
