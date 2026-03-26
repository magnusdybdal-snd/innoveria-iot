// Package config provides erp-agent-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds erp-agent-service runtime configuration.
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
