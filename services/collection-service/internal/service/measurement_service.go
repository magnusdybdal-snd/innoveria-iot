// Package service implements the business logic layer for the collection service.
package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

// MeasurementService implements domain.MeasurementService.
type MeasurementService struct {
	Repo domain.MeasurementRepository
}

// NewMeasurementService creates a new MeasurementService backed by the given repository.
func NewMeasurementService(repo domain.MeasurementRepository) *MeasurementService {
	return &MeasurementService{
		Repo: repo,
	}
}

// Create will store the sensor data in a time series database
func (s *MeasurementService) Create(ctx context.Context, measurement domain.SensorMeasurement, tenantID string) error {
	return s.Repo.Insert(ctx, measurement, tenantID)
}

// GetLatest returns the most recent measurement for the given device, scoped to the caller's company.
func (s *MeasurementService) GetLatest(ctx context.Context, companyID string, deviceEUI string) (domain.SensorMeasurement, error) {
	return s.Repo.FindLatest(ctx, companyID, deviceEUI)
}

// GetByTimeRange returns all measurements for a device within the given time window, scoped to the caller's company.
func (s *MeasurementService) GetByTimeRange(ctx context.Context, companyID string, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return s.Repo.FindByTimeRange(ctx, companyID, deviceEUI, from, to)
}

// GetPayloadKeys returns the distinct payload keys seen in recent measurements for the given device.
func (s *MeasurementService) GetPayloadKeys(ctx context.Context, deviceEUI string) ([]string, error) {
	return s.Repo.FindPayloadKeys(ctx, deviceEUI)
}
