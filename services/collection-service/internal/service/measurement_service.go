package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

type MeasurementService struct {
	Repo domain.MeasurementRepository
}

func NewMeasurementService(repo domain.MeasurementRepository) *MeasurementService {
	return &MeasurementService{
		Repo: repo,
	}
}

// Create will store the sensor data in a time series database
func (s *MeasurementService) Create(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	return s.Repo.Insert(ctx, measurement, tenantID)
}

func (s *MeasurementService) GetLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {
	return s.Repo.FindLatest(ctx, deviceEUI)
}

func (s *MeasurementService) GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return s.Repo.FindByTimeRange(ctx, deviceEUI, from, to)
}
