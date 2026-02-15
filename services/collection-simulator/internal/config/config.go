package config

import (
	"innoveria-iot/pkg/env"
	"log"
	"strconv"
	"strings"
)

type Config struct {
	MQTTBrokerURL string
	Devices int
}

func Load() *Config {
	devices,err := strconv.Atoi(strings.TrimSpace(env.Get("Devices", "100")))
	if err != nil {
		log.Fatal("config loading failed")
	}

	return &Config{
		MQTTBrokerURL: env.Get("MQTT_BROKER_URL", "tcp://mosquitto:1883"),
		Devices: devices,
	}
}
