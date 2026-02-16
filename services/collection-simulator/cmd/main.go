package main

import (
	"log"
	"log/slog"

	"innoveria-iot/collection-simulator/internal/broker"
	"innoveria-iot/collection-simulator/internal/config"
	"innoveria-iot/collection-simulator/internal/simulator"
	"innoveria-iot/pkg/logger"
)

// Service to mock sensor data for the development process
// See config.go for what topic is used
func main() {
	logger.NewLogger("collection-simulator")

	cfg := config.Load()

	client, err := broker.New(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	sim := simulator.New(client, *cfg)

	slog.Info("Started a collection-simulator",
		"MQTT_Topic", cfg.Topic,
		"devices", cfg.Devices,
		"qos", cfg.Qos)

	sim.Start()

}
