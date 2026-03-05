// Package config provides auth-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds auth-service runtime configuration.
type Config struct {
	Addr string
}

// Load reads configuration from environment variables.
func Load() *Config {
	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
	}
}
