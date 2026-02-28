package domain

import "context"

// SensorGroup is how LoRaWAN groups devices(sensors). Chirpstack calls this application
// SensorGroup is a prerequisite for adding sensors
// Gateway will only packet forward data. Application is responsable for managing devices per project
type SensorGroup struct {
	Id        string
	Name      string
	CompanyId string
	Location  string
}

type SensorGroupService interface {
	Create(ctx context.Context, payload SensorGroup) error
	Update(ctx context.Context, sensorGroupId string, payload SensorGroup) error
	GetAll(ctx context.Context) ([]SensorGroup, error)
	Delete(ctx context.Context, sensorGroupId string) error
}
