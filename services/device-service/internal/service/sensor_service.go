package service

import (
	"context"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

type SensorServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewSensorService(cc *chirpstackrest.Client) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc: cc,
	}
}

func (s *SensorServiceImpl) GetAll(ctx context.Context) ([]domain.Sensor, error) {
	// 1. get sensor meta data from database
	limit := 1
	applicationId := ""

	// 2. Get status from chirpstack
	resp, err := s.cc.GetAllSensors(ctx, limit, applicationId)
	if err != nil {
		return nil, err
	}
	var result []domain.Sensor

	for _, sensor := range resp.Result {
		result = append(result, mappers.MapChirpstackSensor(sensor))
	}
	// 3. merge to sensor domain

	return result, nil
}
