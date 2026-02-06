package config

import (
	"innoveria-iot/pkg/env"
	"strings"
)

type Config struct {
	Addr       string
	CollSvcURL string
	// Some serviceURL
	// Some serviceAPIKEY?
}

// Loads the spesific enviroment variables
func Load() *Config {
	cfg := Config{
		Addr: ":" + env.Get("PORT", "8080"),
		CollSvcURL: strings.TrimSpace(env.Get("COLLECTION_SERVICE", "http://localhost:8082")),
		// Some serviceURL
	}

	return &cfg
}
