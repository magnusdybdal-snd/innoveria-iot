package domain

import "context"

// SensorProfile is the LoRaWAN device profile template used when registering sensors in the network.
// It defines the network configuration for a specific hardware model.
type SensorProfile struct {
	Id         string // device profile id
	Name       string
	Region     string // LoRaWAN region. Our domain is limited to EU868
	MACVersion string // LoRaWAN version
	VendorId   string // Identification of model producer
	VendorName string
	// IsCustom       bool // TODO: add a custom device profiles
}

// SensorProfileService defines the business logic operations for sensor profiles.
type SensorProfileService interface {
	GetAll(ctx context.Context) ([]SensorProfile, error)
	GetOne(ctx context.Context) (SensorProfile, error)
	EnsureTenantProfile(ctx context.Context, profileID string, tenantID string) (string, error)
}
