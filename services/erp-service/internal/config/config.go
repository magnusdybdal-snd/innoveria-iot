// Package config provides erp-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds erp-service runtime configuration.
type Config struct {
	Addr string
	// EnableSwagger bool
}

// Load reads configuration from environment variables.
func Load() *Config {
	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
	}
}
