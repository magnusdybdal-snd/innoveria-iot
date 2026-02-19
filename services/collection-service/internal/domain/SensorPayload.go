package domain

import (
	"context"
	"time"
)

type SensorPayload struct {
	DeviceEUI string    `json:"device_eui"`
	Timestamp time.Time `json:"timestamp"`
	Data      string    `json:"data"`
	CompanyId string    `json:"company_id"`
}

type SensorRepository interface {
	FindOne() SensorPayload
	FindAll() []SensorPayload
	Insert(ctx context.Context, payload SensorPayload) error
}

type SensorService interface {
	GetOne() SensorPayload
	GetAll() []SensorPayload
	Create(ctx context.Context, payload SensorPayload) error
}
