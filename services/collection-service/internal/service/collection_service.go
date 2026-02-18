package service

import (
	"context"
	"innoveria-iot/collection-service/internal/domain"
)

type SessionServiceImpl struct {
	Repo domain.SensorRepository
}

func NewSensorService(repo domain.SensorRepository) *SessionServiceImpl {
	return &SessionServiceImpl{
		Repo: repo,
	}
}

func (s *SessionServiceImpl) GetOne() domain.SensorPayload {
	return s.Repo.FindOne()
}

func (s *SessionServiceImpl) GetAll() []domain.SensorPayload {
	return s.Repo.FindAll()
}

// Create will store the sensor data in a time series database
func (s *SessionServiceImpl) Create(ctx context.Context, payload domain.SensorPayload) error {
	return s.Repo.Insert(ctx, payload)
}
