package service

import (
	"context"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

type SensorGroupServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewDeviceGroupService(cc *chirpstackrest.Client) *SensorGroupServiceImpl {
	return &SensorGroupServiceImpl{
		cc: cc,
	}
}

func (d *SensorGroupServiceImpl) Create(ctx context.Context, payload domain.SensorGroup) error {
	return nil
}
func (d *SensorGroupServiceImpl) Update(ctx context.Context, deviceGroupId string, payload domain.SensorGroup) error {
	return nil
}
func (d *SensorGroupServiceImpl) GetAll(ctx context.Context) ([]domain.SensorGroup, error) {
	return nil, nil
}

func (d *SensorGroupServiceImpl) Delete(ctx context.Context, deviceGroupId string) error {
	return nil
}
