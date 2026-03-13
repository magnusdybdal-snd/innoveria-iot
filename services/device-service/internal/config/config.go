// Package config loads and provides runtime configuration for the device service.
package config

import (
	"fmt"

	"innoveria-iot/pkg/env"
)

// Config holds all runtime configuration values for the device service.
type Config struct {
	Addr   string
	DB_url string

	ChirpstackURL        string // chirpstack rest api url
	ChirpstackSecretPath string // chirpstack api token (bearer token)
}

// Load reads configuration from environment variables and mounted secrets, returning a populated Config.
func Load() *Config {
	// database config
	dbHost := env.Get("DB_HOST", "device-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "device")
	dbPassword := env.Get("DB_PASSWORD", "device")
	dbName := env.Get("DB_NAME", "device")
	sslmode := env.Get("DB_SSLMODE", "disable")

	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
		DB_url: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
			sslmode,
		),
		ChirpstackURL:        env.Get("CHIRPSTACK_REST", "http://chirpstack-rest-api:8090"),
		ChirpstackSecretPath: "/secrets/chirpstack-api-key",
	}
}
