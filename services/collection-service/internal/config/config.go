package config

import (
	"innoveria-iot/pkg/env"
	"strconv"
)

type Config struct {
	// Server Configs
	Addr string

	// MQTT Configs
	MQTTBrokerURL   string
	MQTTClientId    string
	MQTTUsername    string
	MQTTPassword    string
	MQTTTopic       string
	MQTTWorkerCount int
}

// Loads the spesific enviroment variables
func Load() *Config {

	workerCount, err := strconv.Atoi(env.Get("MQTT_WORKER_COUNT", "10"))
	if err != nil {
		workerCount = 10
	}
	
	cfg := Config{
		Addr: ":" + env.Get("PORT", "8080"),

		MQTTBrokerURL: env.Get("MQTT_BROKER_URL", "tcp://localhost:1883"), // Mqtt broker
		MQTTClientId: env.Get("MQTT_CLIENT_ID", ""),
		MQTTUsername: env.Get("MQTT_USERNAME", ""),
		MQTTPassword: env.Get("MQTT_PASSWORD", ""),
		MQTTTopic: env.Get("MQTT_TOPIC", "application/+/device/+/event/up"), // Handles all uplink events
		MQTTWorkerCount: workerCount,

	}

	return &cfg
}
