package repository

import (
	"context"
	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/domain"
	"time"
)

type SensorRepository struct {
	db *db.DB
}

func NewSensorRepository(db *db.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (s *SensorRepository) Insert(ctx context.Context, measurement domain.SensorMeasurement) error {
	return nil
}

func (s *SensorRepository) FindLatest(ctx context.Context, deviceEUI string) (domain.SensorMeasurement, error) {
	return domain.SensorMeasurement{}, nil
}

func (s *SensorRepository) FindByTimeRange(ctx context.Context, deviceEUI string, from, to time.Time) ([]domain.SensorMeasurement, error) {
	return nil, nil
}
