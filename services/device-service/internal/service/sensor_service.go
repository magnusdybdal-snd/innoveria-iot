package service

import (
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

type SensorServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewSensorService(cc *chirpstackrest.Client) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc: cc,
	}
}

func (s *SensorServiceImpl) GetAll() ([]domain.Sensor, error) {
	return nil, nil
}
