package service

import "innoveria-iot/device-service/internal/chirpstackrest"

type SensorProfileServiceImpl struct {
	cc *chirpstackrest.Client
}

func NewSensorProfileServiceImpl(cc *chirpstackrest.Client) *SensorProfileServiceImpl {
	return &SensorProfileServiceImpl{
		cc: cc,
	}
}
