package service

import (
	"context"
	"fmt"
	"log/slog"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
	"innoveria-iot/pkg/ptrutil"
)

// SensorServiceImpl implements domain.SensorService, coordinating between the database and Chirpstack.
type SensorServiceImpl struct {
	cc                   *chirpstackrest.Client
	companycfgRepo       domain.CompanyConfigRepository
	sensorRepo           domain.SensorRepository
	sensorProfileService domain.SensorProfileService
}

// NewSensorService creates a new SensorServiceImpl with the given Chirpstack client and repositories.
func NewSensorService(cc *chirpstackrest.Client, sensorRepo domain.SensorRepository, companycfgRepo domain.CompanyConfigRepository, sensorProfileService domain.SensorProfileService) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc:                   cc,
		sensorRepo:           sensorRepo,
		companycfgRepo:       companycfgRepo,
		sensorProfileService: sensorProfileService,
	}
}

// Create adds a new sensor to Chirpstack and the database.
// Chirpstack is updated first. If the database insert fails, the sensor is deleted from
// Chirpstack as a compensating transaction to keep both systems in sync.
func (s *SensorServiceImpl) Create(ctx context.Context, payload domain.Sensor) error {
	// Fetch the company's Chirpstack tenant ID
	cfg, err := s.companycfgRepo.FindByCompanyID(ctx, payload.CompanyID)
	if err != nil {
		return fmt.Errorf("create sensor: finding company tenant ID: %w", err)
	}

	// Ensure a tenant-level copy of the selected profile exists, get its ID.
	tenantProfileID, err := s.sensorProfileService.EnsureTenantProfile(ctx, payload.ChirpstackProfileID, cfg.ChirpstackTenantID) // payload.ChirpstackProfileID may be global or tenant-level
	if err != nil {
		return fmt.Errorf("create sensor: ensure tenant profile: %w", err)
	}
	payload.ChirpstackProfileID = tenantProfileID

	// Post request to Chirpstack
	sensorReq := mappers.MapChirpstackSensorRequest(payload, cfg.ChirpstackApplicationID)
	if err := s.cc.CreateSensor(ctx, sensorReq); err != nil {
		return fmt.Errorf("create sensor: add to chirpstack: %w", err)
	}

	// Set AppKey in Chirpstack
	sensorKeyReq := mappers.MapChirpstackSensorKeyRequest(payload)
	if err := s.cc.SetSensorKey(ctx, sensorKeyReq); err != nil {
		// Compensate: Remove from Chirpstack so system stays in sync.
		if compErr := s.cc.DeleteSensor(ctx, payload.DeviceEUI); compErr != nil {
			slog.Error("saga compensation failed: could not delete sensor from chirpstack after set key failure",
				"eui", payload.DeviceEUI, "error", compErr)
		}
		return fmt.Errorf("create sensor: set app key: %w", err)
	}

	// If successfully stored in Chirpstack, try to store in database.
	sensor, err := s.sensorRepo.Create(ctx, payload)
	if err != nil {
		// Compensate: remove from Chirpstack so systems stay in sync.
		if compErr := s.cc.DeleteSensor(ctx, payload.DeviceEUI); compErr != nil {
			slog.Error("saga compensation failed: could not delete sensor from chirpstack after db insert failure",
				"eui", payload.DeviceEUI, "error", compErr)
		}
		return fmt.Errorf("create sensor: add to database: %w", err)
	}

	slog.Info("successfully created sensor", "id", sensor.Id)
	return nil
}

// Update updates a sensor's metadata in Chirpstack first, then in the database.
// If the database update fails, the Chirpstack update is reverted as a compensating transaction.
func (s *SensorServiceImpl) Update(ctx context.Context, companyID string, sensorID string, payload domain.Sensor) error {
	// Verify that the sensor exists in the database and belongs to the caller's company.
	sensor, err := s.sensorRepo.FindByID(ctx, companyID, sensorID)
	if err != nil {
		return fmt.Errorf("update sensor: sensor %s not found in database: %w", sensorID, err)
	}

	// Get company's Chirpstack tenant ID for call to Chirpstack.
	companycfg, err := s.companycfgRepo.FindByCompanyID(ctx, sensor.CompanyID)
	if err != nil {
		return fmt.Errorf("update sensor: finding company tenant ID: %w", err)
	}

	// Ensure a tenant-level copy of the selected profile exists, get its ID.
	tenantProfileID, err := s.sensorProfileService.EnsureTenantProfile(ctx, payload.ChirpstackProfileID, companycfg.ChirpstackTenantID)
	if err != nil {
		return fmt.Errorf("update sensor: ensure tenant profile: %w", err)
	}
	payload.ChirpstackProfileID = tenantProfileID

	// Retain old values before merging payload, needed for potential compensation.
	oldSensor := sensor

	// Update values if not nil / empty string.
	if payload.Name != "" {
		sensor.Name = payload.Name
	}
	if payload.Description != nil {
		sensor.Description = payload.Description
	}
	if payload.ChirpstackProfileID != "" {
		sensor.ChirpstackProfileID = payload.ChirpstackProfileID
	}
	if payload.FactoryID != "" {
		sensor.FactoryID = payload.FactoryID
	}
	if payload.FactoryAreaID != "" {
		sensor.FactoryAreaID = payload.FactoryAreaID
	}
	if payload.ProductionResource != nil {
		sensor.ProductionResource = payload.ProductionResource
	}

	// Only call Chirpstack if a Chirpstack-stored field changed (name, description, profile).
	// DB-only fields (factory_id, factory_area_id, production_resource) skip Chirpstack entirely.
	chirpstackChanged := sensor.Name != oldSensor.Name ||
		ptrutil.Deref(sensor.Description) != ptrutil.Deref(oldSensor.Description) ||
		sensor.ChirpstackProfileID != oldSensor.ChirpstackProfileID

	if chirpstackChanged {
		// Get company's Chirpstack application ID for the API call.
		companycfg, err := s.companycfgRepo.FindByCompanyID(ctx, sensor.CompanyID)
		if err != nil {
			return fmt.Errorf("update sensor: finding company tenant ID: %w", err)
		}

		newReq := mappers.MapChirpstackSensorRequest(sensor, companycfg.ChirpstackApplicationID)
		if err := s.cc.UpdateSensor(ctx, newReq); err != nil {
			return fmt.Errorf("update sensor: update in chirpstack: %w", err)
		}

		// If successfully updated in Chirpstack, try to update in database.
		if err := s.sensorRepo.Update(ctx, companyID, sensorID, sensor); err != nil {
			// Compensate: revert Chirpstack to old values.
			oldReq := mappers.MapChirpstackSensorRequest(oldSensor, companycfg.ChirpstackApplicationID)
			if compErr := s.cc.UpdateSensor(ctx, oldReq); compErr != nil {
				slog.Error("saga compensation failed: could not revert sensor in chirpstack after db update failure",
					"id", sensorID, "error", compErr)
			}
			return fmt.Errorf("update sensor: update in database: %w", err)
		}
	} else {
		// No Chirpstack fields changed — update DB only.
		if err := s.sensorRepo.Update(ctx, companyID, sensorID, sensor); err != nil {
			return fmt.Errorf("update sensor: update in database: %w", err)
		}
	}

	slog.Info("successfully updated sensor", "id", sensorID)
	return nil
}

// GetAll retrieves all sensors belonging to a companyID from the database and merges the
// response with the status from Chirpstack (status and last seen).
func (s *SensorServiceImpl) GetAll(ctx context.Context, companyID string) ([]domain.Sensor, error) {

	// fetch all sensor belonging to the company in db
	sensors, err := s.sensorRepo.FindAllByCompanyID(ctx, companyID)
	if err != nil {
		return nil, fmt.Errorf("get all sensors: getting sensors from db: %w", err)
	}

	// Sensors to be returned
	var result []domain.Sensor

	// Loop over sensors and get their chirpstack status, merge and append response
	for _, sensor := range sensors {
		status, err := s.cc.GetOneSensor(ctx, sensor.DeviceEUI)
		if err != nil {
			// If no status form Chirpstack, append sensor without status / last seen
			slog.Warn("failed to fetch sensor from chirpstack", "eui", sensor.DeviceEUI, "error", err)
			result = append(result, sensor)
			continue
		}
		result = append(result, mappers.MergeSensor(status, sensor))
	}

	return result, nil
}

// GetByProductionResourceID  retrieves all sensors attatched to one production resource ID from the database and merges the
// response with the status from Chirpstack (status and last seen).
func (s *SensorServiceImpl) GetByProductionResourceID(ctx context.Context, companyID string, productionResourceID string) ([]domain.Sensor, error) {

	sensors, err := s.sensorRepo.FindByProductionResourceID(ctx, companyID, productionResourceID)
	if err != nil {
		return nil, fmt.Errorf("get sensors by production resource: %w", err)
	}

	var result []domain.Sensor

	// Loop over sensors and get their chirpstack status, merge and append response
	for _, sensor := range sensors {
		status, err := s.cc.GetOneSensor(ctx, sensor.DeviceEUI)
		if err != nil {
			// If no status from Chirpstack, append sensor without status / last seen
			slog.Warn("failed to fetch sensor from chirpstack", "eui", sensor.DeviceEUI, "error", err)
			result = append(result, sensor)
			continue
		}
		result = append(result, mappers.MergeSensor(status, sensor))
	}

	return result, nil
}

// GetByID retrieves a single sensor by its ID, scoped to the caller's company.
func (s *SensorServiceImpl) GetByID(ctx context.Context, companyID string, sensorID string) (domain.Sensor, error) {
	sensor, err := s.sensorRepo.FindByID(ctx, companyID, sensorID)
	if err != nil {
		return domain.Sensor{}, fmt.Errorf("get sensor by id: %w", err)
	}

	return sensor, nil
}

// Delete removes a sensor from Chirpstack and then from the database.
// If the database delete fails, the sensor is re-created in Chirpstack as a compensating
// transaction to keep both systems in sync.
func (s *SensorServiceImpl) Delete(ctx context.Context, companyID string, deviceID string) error {
	// Retrieve the sensor from the database, scoped to the caller's company.
	sensor, err := s.sensorRepo.FindByID(ctx, companyID, deviceID)
	if err != nil {
		return fmt.Errorf("delete sensor: sensor %s not found in database: %w", deviceID, err)
	}

	// Fetch companycfg before deletion — needed for compensation if DB delete fails.
	companycfg, err := s.companycfgRepo.FindByCompanyID(ctx, sensor.CompanyID)
	if err != nil {
		return fmt.Errorf("delete sensor: finding company tenant ID: %w", err)
	}

	// Delete request to Chirpstack.
	if err := s.cc.DeleteSensor(ctx, sensor.DeviceEUI); err != nil {
		return fmt.Errorf("delete sensor: delete in chirpstack: %w", err)
	}

	// If successfully deleted in Chirpstack, try to delete from database.
	if err := s.sensorRepo.Delete(ctx, companyID, deviceID); err != nil {
		// Compensate: re-create in Chirpstack so systems stay in sync.
		sensorReq := mappers.MapChirpstackSensorRequest(sensor, companycfg.ChirpstackApplicationID)
		if compErr := s.cc.CreateSensor(ctx, sensorReq); compErr != nil {
			slog.Error("saga compensation failed: could not re-create sensor in chirpstack after db delete failure",
				"eui", sensor.DeviceEUI, "error", compErr)
		} else {
			// If compensation is successfull we also set the key.
			keyReq := mappers.MapChirpstackSensorKeyRequest(sensor)
			if compErr := s.cc.SetSensorKey(ctx, keyReq); compErr != nil {
				slog.Error("saga compensation failed: could not re-set app key in chirpstack after db delete failure",
					"eui", sensor.DeviceEUI, "error", compErr)
			}
		}
		return fmt.Errorf("delete sensor: delete in database: %w", err)
	}

	slog.Info("successfully deleted sensor", "id", deviceID)
	return nil
}
