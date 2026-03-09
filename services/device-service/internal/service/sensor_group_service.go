package service

import (
	"context"

	"innoveria-iot/device-service/internal/chirpstackrest"
	"innoveria-iot/device-service/internal/domain"
	"innoveria-iot/device-service/internal/service/mappers"
)

// SensorGroupServiceImpl implements domain.SensorGroupService, managing sensor groups via Chirpstack applications.
type SensorGroupServiceImpl struct {
	cc *chirpstackrest.Client
}

// NewDeviceGroupService creates a new SensorGroupServiceImpl with the given Chirpstack client.
func NewDeviceGroupService(cc *chirpstackrest.Client) *SensorGroupServiceImpl {
	return &SensorGroupServiceImpl{
		cc: cc,
	}
}

// Create registers a new sensor group as a Chirpstack application.
func (d *SensorGroupServiceImpl) Create(ctx context.Context, payload domain.SensorGroup) error {

	data := mappers.MapCreateChirpstackApplication(payload)

	if err := d.cc.CreateApplication(ctx, data); err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all sensor groups from Chirpstack up to the given limit.
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

// Update is not yet implemented.
func (d *SensorGroupServiceImpl) Update(ctx context.Context, deviceGroupId string, payload domain.SensorGroup) error {
	return nil
}

// Delete is not yet implemented.
func (d *SensorGroupServiceImpl) Delete(ctx context.Context, deviceGroupId string) error {
	return nil
}
