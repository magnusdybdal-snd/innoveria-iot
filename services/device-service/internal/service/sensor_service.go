package service

import (
	"context"
	"fmt"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
	"log/slog"
)

// SensorServiceImpl TODO(@vinjar): add proper documentation.
type SensorServiceImpl struct {
	cc             *chirpstackrest.Client
	companycfgRepo domain.CompanyConfigRepository
	sensorRepo     domain.SensorRepository
}

// NewSensorService TODO(@vinjar): add proper documentation.
func NewSensorService(cc *chirpstackrest.Client, sensorRepo domain.SensorRepository, companycfgRepo domain.CompanyConfigRepository) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc:             cc,
		sensorRepo:     sensorRepo,
		companycfgRepo: companycfgRepo,
	}
}

// Create TODO
func (s *SensorServiceImpl) Create(ctx context.Context, payload domain.Sensor) error {
	// Find the company's chirpstack tenant ID
	_, err := s.companycfgRepo.FindByCompanyID(ctx, payload.CompanyID)
	if err != nil {
		return fmt.Errorf("create sensor: finding company tenant ID: %w", err)
	}

	// Sending post request to chirpstack
	return nil
}

// Update TODO
func (s *SensorServiceImpl) Update(ctx context.Context, deviceID string, sensor domain.Sensor) error {
	return nil
}

// GetAll retrieves all sensors belonging to a companyID from the database and merges the
// response with the status from Chirpstack (status and last seen)
func (s *SensorServiceImpl) GetAll(ctx context.Context) ([]domain.Sensor, error) {

	// fetch all sensor belonging to the company in db
	sensors, err := s.sensorRepo.FindAllByCompanyID(ctx, "a0000000-0000-0000-0000-000000000001")
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

// Delete deletes a sensor from both Chirpstack and from the database. Deletion in
// Chirpstack is always tried first so we keep database entry if we fail.
// Any failure will return early to prevent desyncing chirpstack and the database.
func (s *SensorServiceImpl) Delete(ctx context.Context, deviceID string) error {
	// Get the sensor EUI from database
	sensor, err := s.sensorRepo.FindByID(ctx, deviceID)
	if err != nil {
		return fmt.Errorf("delete sensor: sensor %s not found in database: %w", deviceID, err)
	}

	// Delete in chirpstack
	err = s.cc.DeleteSensor(ctx, sensor.DeviceEUI)
	if err != nil {
		return fmt.Errorf("delete sensor: delete in chirpstack: %w", err)
	}

	// Delete in database after successfully deleting in Chirpstack
	err = s.sensorRepo.Delete(ctx, deviceID)
	if err != nil {
		return fmt.Errorf("delete sensor: delete in database: %w", err)
	}

	slog.Info("successfully deleted sensor", "id", deviceID)
	return nil
}
