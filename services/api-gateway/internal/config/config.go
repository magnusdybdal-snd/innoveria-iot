// Package config TODO(@vinjar): add proper documentation.
package config

import (
	"strings"

	"innoveria-iot/pkg/env"
)

// Config TODO(@vinjar): add proper documentation.
type Config struct {
	Addr             string
	CollSvcURL       string
	DeviceSvcURL     string
	AuthSvcURL       string
	OnboardingSvcURL string
	ContextSvcURL    string
	// Some serviceURL
	// Some serviceAPIKEY?
}

// Load the spesific enviroment variables
func Load() *Config {
	cfg := Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		CollSvcURL:       strings.TrimSpace(env.Get("COLLECTION_SERVICE", "http://collection-service:8080")),
		DeviceSvcURL:     strings.TrimSpace(env.Get("DEVICE_SERVICE", "http://device-service:8080")),
		AuthSvcURL:       strings.TrimSpace(env.Get("AUTH_SERVICE", "http://auth-service:8080")),
		OnboardingSvcURL: strings.TrimSpace(env.Get("ONBOARDING_SERVICE", "http://onboarding-service:8080")),
		ContextSvcURL:    strings.TrimSpace(env.Get("CONTEXT_SERVICE", "http://context-service:8080")),
	}

	return &cfg
}
