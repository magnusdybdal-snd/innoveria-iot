package domain

import (
	"context"
	"time"
)

// Sensor represents a LoRaWAN sensor in our system.
// A sensor is assigned to a chirpstack application and send data trough.
// whichever gateway is in range. There is no fixed gateway assignment in LoRaWAN.
type Sensor struct {
	Id                  string
	CompanyID           string
	DeviceEUI           string // hardware identifier, shared key with Chirpstack
	Name                string
	Description         *string
	State               DeviceState // administrative state: ACTIVE / INACTIVE
	FactoryAreaID       *string     // loose cross-service ref
	ProductionResource  *string     // loose cross-service ref
	ChirpstackProfileID string      // LoRaWAN template that describes device model, chosen on registration
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// Runtime fields - populated from Chirpstack, not stored in DB
	Status     Status // Chirpstack connectivity: 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

// SensorService is the interface for sensor methods
type SensorService interface {
	Create(ctx context.Context, payload Sensor) error
	Update(ctx context.Context, deviceID string, payload Sensor) error
	GetAll(ctx context.Context) ([]Sensor, error)
	Delete(ctx context.Context, deviceID string) error
}

// SensorRepository handles persistance of sensor meta data in our database.
// Chirpstack operations are handled in the service layer
type SensorRepository interface {
	Create(ctx context.Context, sensor Sensor) (Sensor, error)
	FindByID(ctx context.Context, sensorID string) (Sensor, error)
	FindAllByCompanyID(ctx context.Context, companyID string) ([]Sensor, error)
	FindByEUI(ctx context.Context, deviceEUI string) (Sensor, error)
	UpdateState(ctx context.Context, sensorID string, state DeviceState) error
	Update(ctx context.Context, deviceID string, sensor Sensor) error
	Delete(ctx context.Context, deviceID string) error
}
