package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

type SensorProfileServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewSensorProfileService(cc *chirpstackrest.Client) *SensorProfileServiceImpl {
	return &SensorProfileServiceImpl{
		cc: cc,
	}
}

func (s *SensorProfileServiceImpl) GetAll(ctx context.Context) ([]domain.SensorProfile, error) {
	return nil, nil
}

func (s *SensorProfileServiceImpl) GetOne(ctx context.Context) (domain.SensorProfile, error) {
	return domain.SensorProfile{}, nil
}
