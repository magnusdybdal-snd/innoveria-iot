package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

type CollectionService struct {
	Repo domain.SensorRepository
}

func NewSensorService(repo domain.SensorRepository) *CollectionService {
	return &CollectionService{
		Repo: repo,
	}
}

// Create will store the sensor data in a time series database
func (s *CollectionService) Create(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	return s.Repo.Insert(ctx, measurement, tenantID)
}

func (s *CollectionService) GetLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {
	return s.Repo.FindLatest(ctx, deviceEUI)
}

func (s *CollectionService) GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return s.Repo.FindByTimeRange(ctx, deviceEUI, from, to)
}
