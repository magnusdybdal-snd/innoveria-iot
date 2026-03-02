package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
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

	data := mappers.MapCreateChirpstackApplication(payload)

	if err := d.cc.CreateApplication(ctx, data); err != nil {
		return err
	}
	return nil
}

func (d *SensorGroupServiceImpl) GetAll(ctx context.Context, limit int) ([]domain.SensorGroup, error) {
	resp, err := d.cc.GetAllApplication(ctx, limit)
	if err != nil {
		return nil, err
	}

	var result []domain.SensorGroup
	for _, sg := range resp.Result {
		result = append(result, mappers.MapChirpstackSensorGroupDtoToDomain(sg))
	}

	return result, nil
}

// These are not important for mvp
// TODO
func (d *SensorGroupServiceImpl) Update(ctx context.Context, deviceGroupId string, payload domain.SensorGroup) error {
	return nil
}

// TODO
func (d *SensorGroupServiceImpl) Delete(ctx context.Context, deviceGroupId string) error {
	return nil
}
