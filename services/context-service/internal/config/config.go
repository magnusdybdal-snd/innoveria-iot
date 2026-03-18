// Package config provides context-service runtime configuration.
package config

import (
	"fmt"
	"innoveria-iot/pkg/env"
)

// Config holds context-service runtime configuration.
type Config struct {
	Addr             string
	CollectionSvcURL string
	DB_URL           string
}

// Load reads configuration from environment variables.
func Load() *Config {

	dbHost := env.Get("DB_HOST", "context-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "context")
	dbPassword := env.Get("DB_PASSWORD", "context")
	dbName := env.Get("DB_NAME", "context")
	sslmode := env.Get("DB_SSLMODE", "disable")

	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		CollectionSvcURL: env.Get("COLLECTION_SERVICE", "http://collection-service:8080"),
		DB_URL: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
			sslmode,
		),
	}
}
