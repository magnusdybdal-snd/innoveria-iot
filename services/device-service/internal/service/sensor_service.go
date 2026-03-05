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
	cc          *chirpstackrest.Client
	companyRepo domain.CompanyConfigRepository
	sensorRepo  domain.SensorRepository
}

// NewSensorService TODO(@vinjar): add proper documentation.
func NewSensorService(cc *chirpstackrest.Client, sensorRepo domain.SensorRepository, companyRepo domain.CompanyConfigRepository) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc:          cc,
		sensorRepo:  sensorRepo,
		companyRepo: companyRepo,
	}
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

// Create TODO(@vinjar): add proper documentation.
func (s *SensorServiceImpl) Create(ctx context.Context) error {
	return nil
}
