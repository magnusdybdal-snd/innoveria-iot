package domain

import (
	"context"
	"time"
)

// Gateway represents a LoRaWAN gateway in our system.
// They receive signals from sensors and forward to Chirpstack.
// A gateway has no knowledge of which sensors they serve.
type Gateway struct {
	Id            string
	CompanyId     string
	GatewayEUI    string
	Name          string
	Description   *string
	State         DeviceState // Administrative state: ACTIVE / INACTIVE
	FactoryID     string
	FactoryAreaID string
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Runtime fields - not stored in our database
	Status     Status // Chirpstack connectivity: 0=online, 1=never_seen, 2=offline
	LastSeenAt string
}

// GatewayService defines the business logic operations for gateways.
type GatewayService interface {
	Create(ctx context.Context, payload Gateway) error
	Update(ctx context.Context, gatewayId string, payload Gateway) error
	GetAll(ctx context.Context) ([]Gateway, error)
	Delete(ctx context.Context, gatewayID string) error
}

// GatewayRepository handles persistance of gateway metadata in our database.
// Chirpstack operations are handled seperately in service layer
type GatewayRepository interface {
	Create(ctx context.Context, gateway Gateway) (Gateway, error)
	FindByID(ctx context.Context, gatewayID string) (Gateway, error)
	FindAllByCompanyID(ctx context.Context, companyID string) ([]Gateway, error)
	FindByEUI(ctx context.Context, gatewayEUI string) (Gateway, error)
	UpdateState(ctx context.Context, gatewayID string, state DeviceState) error
	Update(ctx context.Context, gatewayID string, payload Gateway) error
	Delete(ctx context.Context, gatewayID string) error
}
