package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// SensorProfileServiceImpl implements domain.SensorProfileService, fetching sensor profiles from Chirpstack.
type SensorProfileServiceImpl struct {
	cc *chirpstackrest.Client
}

// NewSensorProfileService creates a new SensorProfileServiceImpl with the given Chirpstack client.
func NewSensorProfileService(cc *chirpstackrest.Client) *SensorProfileServiceImpl {
	return &SensorProfileServiceImpl{
		cc: cc,
	}
}

// GetAll retrieves all available sensor profiles from Chirpstack up to the given limit.
func (s *SensorProfileServiceImpl) GetAll(ctx context.Context, limit int) ([]domain.SensorProfile, error) {
	resp, err := s.cc.GetAllSensorProfiles(ctx)
	if err != nil {
		return nil, err
	}
	var result []domain.SensorProfile
	for _, sp := range resp {
		result = append(result, mappers.MapChirpstackDeviceProfilesToDomain(sp))
	}

	return result, nil
}

// GetOne is not yet implemented.
func (s *SensorProfileServiceImpl) GetOne(ctx context.Context) (domain.SensorProfile, error) {
	return domain.SensorProfile{}, nil
}
