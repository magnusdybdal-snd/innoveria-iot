package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

type SessionServiceImpl struct {
	Repo domain.SensorRepository
}

func NewSensorService(repo domain.SensorRepository) *SessionServiceImpl {
	return &SessionServiceImpl{
		Repo: repo,
	}
}

// Create will store the sensor data in a time series database
func (s *SessionServiceImpl) Create(ctx context.Context, measurement domain.SensorMeasurement) error {
	return s.Repo.Insert(ctx, measurement)
}

func (s *SessionServiceImpl) GetLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {
	return s.Repo.FindLatest(ctx, deviceEUI)
}

func (s *SessionServiceImpl) GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return s.Repo.FindByTimeRange(ctx, deviceEUI, from, to)
}
