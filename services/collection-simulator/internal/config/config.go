package config

import "innoveria-iot/pkg/env"

type Config struct {
	MQTTBrokerURL string
}

func Load() *Config {
	return &Config{
		MQTTBrokerURL: env.Get("MQTT_BROKER_URL", "tcp://mosquitto:1883"),
	}
}
