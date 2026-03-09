// Package config loads and provides runtime configuration for the onboarding service.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds all runtime configuration values for the onboarding service.
type Config struct {
	Addr string

	AuthSvcURL       string
	DeviceSvcURL     string
	CollectionSvcURL string
}

// Load reads configuration from environment variables and returns a populated Config.
func Load() *Config {
	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		AuthSvcURL:       env.Get("AUTH_SERVICE", "http://auth-service:8080"),
		DeviceSvcURL:     env.Get("DEVICE_SERVICE", "http://device-service:8080"),
		CollectionSvcURL: env.Get("COLLECTION_SERVICE", "http://collection-service:8080"),
	}
}
