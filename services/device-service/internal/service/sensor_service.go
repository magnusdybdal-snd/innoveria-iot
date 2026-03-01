package service

import (
	"context"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

type SensorServiceImpl struct {
	cc          *chirpstackrest.Client
	companyRepo domain.CompanyConfigRepository
}

func NewSensorService(cc *chirpstackrest.Client, companyRepo domain.CompanyConfigRepository) *SensorServiceImpl {
	return &SensorServiceImpl{
		cc:          cc,
		companyRepo: companyRepo,
	}
}

func (s *SensorServiceImpl) GetAll(ctx context.Context) ([]domain.Sensor, error) {
	// 1. get sensor meta data from database

	// TODO: fix this when tennant system is working
	limit := 1

	// TODO: REPLACE HARDCODED companyID once auth exists
	cfg, err := s.companyRepo.FindByCompanyID(ctx, "a0000000-0000-0000-0000-000000000001")
	if err != nil {
		return nil, err
	}
	applicationID := cfg.ChirpstackApplicationID

	// 2. Get status from chirpstack
	resp, err := s.cc.GetAllSensors(ctx, limit, applicationID)
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

func (s *SensorServiceImpl) Create(ctx context.Context) error {
	return nil
}
