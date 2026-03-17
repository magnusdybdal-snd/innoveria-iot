// Package config provides context-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds context-service runtime configuration.
type Config struct {
	Addr             string
	CollectionSvcURL string
}

// Load reads configuration from environment variables.
func Load() *Config {
	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		CollectionSvcURL: env.Get("COLLECTION_SERVICE", "http://collection-service:8080"),
	}
}
