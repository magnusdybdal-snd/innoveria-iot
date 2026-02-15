package main

import (
	"innoveria-iot/collection-simulator/internal/broker"
	"innoveria-iot/collection-simulator/internal/config"
	"innoveria-iot/pkg/logger"
	"log"
)

type sensorReading struct {
	DeviceID string `json:"DeviceEUI"`
}

func main() {
	logger.NewLogger("collection-simulator")

	cfg := config.Load()

	client, err := broker.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
