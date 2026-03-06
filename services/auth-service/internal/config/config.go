// Package config provides auth-service runtime configuration.
package config

import (
	"fmt"
	"innoveria-iot/pkg/env"
)

// Config holds auth-service runtime configuration.
type Config struct {
	Addr   string
	DB_URL string
}

// Load reads configuration from environment variables.
func Load() *Config {

	dbHost := env.Get("DB_HOST", "auth-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "auth")
	dbPassword := env.Get("DB_PASSWORD", "auth")
	dbName := env.Get("DB_NAME", "auth")
	sslmode := env.Get("DB_SSLMODE", "disable")

	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
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
