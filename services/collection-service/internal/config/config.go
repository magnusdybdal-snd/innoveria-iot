package config

import (
	"strconv"

	"innoveria-iot/pkg/env"
)

type Config struct {
	// Server Configs
	Addr   string
	DB_url string

	// MQTT Configs
	MQTTBrokerURL   string
	MQTTClientId    string
	MQTTUsername    string
	MQTTPassword    string
	MQTTTopic       string //"application/+/device/+/event/up" will listen to all
	MQTTWorkerCount int
}

// Loads the spesific enviroment variables
func Load() *Config {
	clientID := env.Get("MQTT_CLIENT_ID", "")
	if clientID == "" {
		clientID = "collection-service"
	}

	workerCount, err := strconv.Atoi(env.Get("MQTT_WORKER_COUNT", "10"))
	if err != nil {
		workerCount = 10
	}

	cfg := Config{
		Addr:   ":" + env.Get("PORT", "8080"),
		DB_url: env.Get("DB_url", "postgres://collection:collection@collection-db:5432/collection?sslmode=disable"),

		MQTTBrokerURL:   env.Get("MQTT_BROKER_URL", "tcp://mosquitto:1883"), // Mqtt broker
		MQTTClientId:    clientID,
		MQTTUsername:    env.Get("MQTT_USERNAME", ""),
		MQTTPassword:    env.Get("MQTT_PASSWORD", ""),
		MQTTTopic:       env.Get("MQTT_TOPIC", "test/sensors/#"), // Handles all uplink events
		MQTTWorkerCount: workerCount,
	}

	return &cfg
}
