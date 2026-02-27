package service

import (
	"context"
	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
)

type DeviceGroupServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewDeviceGroupService(cc *chirpstackrest.Client) *DeviceGroupServiceImpl {
	return &DeviceGroupServiceImpl{
		cc: cc,
	}
}

func (d *DeviceGroupServiceImpl) Create(ctx context.Context, payload domain.DeviceGroup) error {
	return nil
}
func (d *DeviceGroupServiceImpl) Update(ctx context.Context, deviceGroupId string, payload domain.DeviceGroup) error {
	return nil
}
func (d *DeviceGroupServiceImpl) GetAll(ctx context.Context) ([]domain.DeviceGroup, error) {
	return nil, nil
}

func (d *DeviceGroupServiceImpl) Delete(ctx context.Context, deviceGroupId string) error {
	return nil
}
