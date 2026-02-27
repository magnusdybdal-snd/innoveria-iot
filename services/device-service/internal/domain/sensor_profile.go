package domain

import "context"

type SensorProfile struct {
	Id         string // device profile id
	Name       string
	Region     string // LoRaWAN region. Our domain is limited to EU868
	MACVersion string // LoRaWAN version
	VendorId   string // Identification of model producer
	VendorName string
	// IsCustom       bool // TODO: add a custom device profiles
}

type SensorProfileService interface {
	GetAll(ctx context.Context) ([]SensorProfile, error)
	GetOne(ctx context.Context) (SensorProfile, error)
}
