// Package config TODO(@vinjar): add proper documentation.
package config

import (
	"fmt"
	"strconv"

	"innoveria-iot/pkg/env"
)

// Config TODO(@vinjar): add proper documentation.
type Config struct {
	// Server Configs
	Addr          string
	DB_url        string
	EnableSwagger bool

	// MQTT Configs
	MQTTBrokerURL   string
	MQTTClientId    string
	MQTTUsername    string
	MQTTPassword    string
	MQTTTopic       string // "application/+/device/+/event/up" will listen to all
	MQTTWorkerCount int
}

// Load the spesific enviroment variables
func Load() *Config {
	clientID := env.Get("MQTT_CLIENT_ID", "")
	if clientID == "" {
		clientID = "collection-service"
	}

	workerCount, err := strconv.Atoi(env.Get("MQTT_WORKER_COUNT", "10"))
	if err != nil {
		workerCount = 10
	}
	// database config
	dbHost := env.Get("DB_HOST", "collection-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "collection")
	dbPassword := env.Get("DB_PASSWORD", "collection")
	dbName := env.Get("DB_NAME", "collection")
	sslmode := env.Get("DB_SSLMODE", "disable")

	cfg := Config{
		Addr: ":" + env.Get("PORT", "8080"),
		DB_url: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
			sslmode,
		),
		EnableSwagger: env.GetBool("ENABLE_SWAGGER", false),

		MQTTBrokerURL:   env.Get("MQTT_BROKER_URL", "tcp://mosquitto:1883"), // Mqtt broker
		MQTTClientId:    clientID,
		MQTTUsername:    env.Get("MQTT_USERNAME", ""),
		MQTTPassword:    env.Get("MQTT_PASSWORD", ""),
		MQTTTopic:       env.Get("MQTT_TOPIC", "test/sensors/#"), // Handles all uplink events
		MQTTWorkerCount: workerCount,
	}

	return &cfg
}
