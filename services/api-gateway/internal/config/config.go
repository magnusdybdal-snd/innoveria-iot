package config

import (
	"strings"

	"innoveria-iot/pkg/env"
)

type Config struct {
	Addr         string
	CollSvcURL   string
	DeviceSvcURL string
	// Some serviceURL
	// Some serviceAPIKEY?
}

// Loads the spesific enviroment variables
func Load() *Config {
	cfg := Config{
		Addr:         ":" + env.Get("PORT", "8080"),
		CollSvcURL:   strings.TrimSpace(env.Get("COLLECTION_SERVICE", "http://collection-service:8080")),
		DeviceSvcURL: strings.TrimSpace(env.Get("DEVICE_SERVICE", "http://device-service:8080")),
		// Some serviceURL
	}

	return &cfg
}
