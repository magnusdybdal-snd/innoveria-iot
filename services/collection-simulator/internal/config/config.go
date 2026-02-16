package config

import (
	"log"
	"strconv"
	"strings"

	"innoveria-iot/pkg/env"
)

type Config struct {
	MQTTBrokerURL string
	Topic         string
	Interval      int // in seconds
	Devices       int
	Qos           byte
}

func Load() *Config {
	devices, err := strconv.Atoi(strings.TrimSpace(env.Get("Devices", "1")))
	if err != nil {
		log.Fatal("config loading failed")
	}

	qos, err := strconv.Atoi(strings.TrimSpace(env.Get("qos", "0"))) // 0, 1 or 2
	if err != nil {
		log.Fatal("config loading failed")
	}

	interval, err := strconv.Atoi(strings.TrimSpace(env.Get("interval", "10")))
	if err != nil {
		log.Fatal("config loading failed")
	}

	return &Config{
		MQTTBrokerURL: env.Get("MQTT_BROKER_URL", "tcp://mosquitto:1883"),
		Interval:      interval,
		Topic:         env.Get("MQTT_TOPIC", "test/sensors"),
		Devices:       devices,
		Qos:           byte(qos),
	}
}
