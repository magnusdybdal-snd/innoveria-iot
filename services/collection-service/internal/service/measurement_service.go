// Package service TODO(@Magnus Dybdal): add proper documentation.
package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

// MeasurementService TODO(@Magnus Dybdal): add proper documentation.
type MeasurementService struct {
	Repo domain.MeasurementRepository
}

// NewMeasurementService TODO(@Magnus Dybdal): add proper documentation.
func NewMeasurementService(repo domain.MeasurementRepository) *MeasurementService {
	return &MeasurementService{
		Repo: repo,
	}
}

// Create will store the sensor data in a time series database
func (s *MeasurementService) Create(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	return s.Repo.Insert(ctx, measurement, tenantID)
}

// GetLatest TODO(@Magnus Dybdal): add proper documentation.
func (s *MeasurementService) GetLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {
	return s.Repo.FindLatest(ctx, deviceEUI)
}

// GetByTimeRange TODO(@Magnus Dybdal): add proper documentation.
func (s *MeasurementService) GetByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return s.Repo.FindByTimeRange(ctx, deviceEUI, from, to)
}
