package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// SensorProfileServiceImpl TODO(@vinjar): add proper documentation.
type SensorProfileServiceImpl struct {
	cc *chirpstackrest.Client
}

// NewSensorProfileService TODO(@vinjar): add proper documentation.
func NewSensorProfileService(cc *chirpstackrest.Client) *SensorProfileServiceImpl {
	return &SensorProfileServiceImpl{
		cc: cc,
	}
}

// GetAll returns a list of the domain sensor profiles
func (s *SensorProfileServiceImpl) GetAll(ctx context.Context, limit int) ([]domain.SensorProfile, error) {
	resp, err := s.cc.GetAllSensorProfiles(ctx, limit)
	if err != nil {
		return nil, err
	}
	var result []domain.SensorProfile
	for _, sp := range resp.Result {
		result = append(result, mappers.MapChirpstackDeviceProfilesToDomain(sp))
	}

	return result, nil
}

// GetOne TODO(@vinjar): add proper documentation.
func (s *SensorProfileServiceImpl) GetOne(ctx context.Context) (domain.SensorProfile, error) {
	return domain.SensorProfile{}, nil
}
