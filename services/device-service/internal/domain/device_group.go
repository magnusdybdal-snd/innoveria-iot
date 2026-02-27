package domain

import "context"

// DeviceGroup is how LoRaWAN groups devices(sensors). Chirpstack calls this application
// DeviceGroup is a prerequisite for adding sensors
// Gateway will only packet forward data. Application is responsable for managing devices per project
type DeviceGroup struct {
	Id        string
	Name      string
	CompanyId string
	Location  string
}

type DeviceGroupService interface {
	Create(ctx context.Context, payload DeviceGroup) error
	Update(ctx context.Context, deviceGroupId string, payload DeviceGroup) error
	GetAll(ctx context.Context) ([]DeviceGroup, error)
	Delete(ctx context.Context, deviceGroupId string) error
}
