package domain

import (
	"context"
	"time"
)

// Sensor represents a LoRaWAN sensor in our system.
// A sensor is assigned to a Chirpstack application and sends data through
// whichever gateway is in range. There is no fixed gateway assignment in LoRaWAN.
type Sensor struct {
	Id                  string
	CompanyID           string
	DeviceEUI           string // hardware identifier, shared key with Chirpstack
	AppKey              string // password for sensor to connect to an application
	Name                string
	Description         *string
	State               DeviceState // administrative state: ACTIVE / INACTIVE
	FactoryID           string      // loose cross-service ref
	FactoryAreaID       string      // loose cross-service ref
	ProductionResource  *string     // loose cross-service ref
	ChirpstackProfileID string      // LoRaWAN template that describes device model, chosen on registration
	CreatedAt           time.Time
	UpdatedAt           time.Time

	// Runtime fields - populated from Chirpstack, not stored in DB
	Status     Status // Chirpstack connectivity: 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

// SensorService defines the business logic operations for sensors.
type SensorService interface {
	Create(ctx context.Context, payload Sensor) error
	Update(ctx context.Context, companyID string, deviceID string, payload Sensor) error
	GetAll(ctx context.Context, companyID string) ([]Sensor, error)
	GetByID(ctx context.Context, companyID string, sensorID string) (Sensor, error)
	GetByProductionResourceID(ctx context.Context, companyID string, productionResourceID string) ([]Sensor, error)
	Delete(ctx context.Context, companyID string, deviceID string) error
}

// SensorRepository handles persistence of sensor metadata in the database.
// Chirpstack operations are handled in the service layer.
type SensorRepository interface {
	Create(ctx context.Context, sensor Sensor) (Sensor, error)
	FindByID(ctx context.Context, companyID string, sensorID string) (Sensor, error)
	FindByProductionResourceID(ctx context.Context, companyID string, productionResourceID string) ([]Sensor, error)
	FindAllByCompanyID(ctx context.Context, companyID string) ([]Sensor, error)
	FindByEUI(ctx context.Context, deviceEUI string) (Sensor, error)
	UpdateState(ctx context.Context, companyID string, sensorID string, state DeviceState) error
	Update(ctx context.Context, companyID string, deviceID string, sensor Sensor) error
	Delete(ctx context.Context, companyID string, deviceID string) error
}
