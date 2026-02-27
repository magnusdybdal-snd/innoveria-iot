package domain

import "context"

type SensorProfile struct {
	Id             string
	ChirpProfileId string
	Name           string
	Region         string
	Vendor         string
	IsCustom       bool // TODO: add a custom device profiles
}

type SensorProfileService interface {
	GetAll(ctx context.Context) ([]SensorProfile, error)
	GetOne(ctx context.Context) (SensorProfile, error)
}
