package service

import "innoveria-iot/device-service/internal/chirpstackrest"

type SensorServiceImpl struct {
	cc *chirpstackrest.Client
}
