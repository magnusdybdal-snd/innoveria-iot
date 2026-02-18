package domain

import "context"

type SensorPayload struct {
	Data string `json:"data"` // base64 encoded
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
