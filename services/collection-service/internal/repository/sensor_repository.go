package repository

import (
	"context"
	"innoveria-iot/collection-service/internal/db"
	"innoveria-iot/collection-service/internal/domain"
)

type SensorRepository struct {
	db *db.DB
}

func NewSensorRepository(db *db.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (s *SensorRepository) FindOne() domain.SensorPayload {
	var data domain.SensorPayload
	return data
}

func (s *SensorRepository) FindAll() []domain.SensorPayload {
	return nil
}

func (s *SensorRepository) Insert(ctx context.Context, payload domain.SensorPayload) error {
	return nil
}
