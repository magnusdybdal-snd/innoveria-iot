package domain

import "context"

// SensorProfile is the LoRaWAN device profile template
// used when register sensors in the network
// It defines the network configuration for a specific hardware model
type SensorProfile struct {
	Id         string // device profile id
	Name       string
	Region     string // LoRaWAN region. Our domain is limited to EU868
	MACVersion string // LoRaWAN version
	VendorId   string // Identification of model producer
	VendorName string
	// IsCustom       bool // TODO: add a custom device profiles
}

// SensorProfileService TODO(@vinjar): add proper documentation.
type SensorProfileService interface {
	GetAll(ctx context.Context, limit int) ([]SensorProfile, error)
	GetOne(ctx context.Context) (SensorProfile, error)
}
