package service

import (
	"context"

	"innoveria-iot/device-service/internal/domain"
)

// SensorProfileConfigServiceImpl implements domain.SensorProfileConfigService.
type SensorProfileConfigServiceImpl struct {
	repo domain.SensorProfileConfigRepository
}

// NewSensorProfileConfigService creates a new SensorProfileConfigServiceImpl with the given repository.
func NewSensorProfileConfigService(repo domain.SensorProfileConfigRepository) *SensorProfileConfigServiceImpl {
	return &SensorProfileConfigServiceImpl{repo: repo}
}

// Get implements SensorProfileConfigService.
func (s *SensorProfileConfigServiceImpl) Get(ctx context.Context, chirpstackProfileID string) (domain.SensorProfileConfig, error) {
	return s.repo.Get(ctx, chirpstackProfileID)
}

// Upsert implements SensorProfileConfigService.
func (s *SensorProfileConfigServiceImpl) Upsert(ctx context.Context, sensorProfileCfg domain.SensorProfileConfig) error {
	return s.repo.Upsert(ctx, sensorProfileCfg)
}
