// Package config TODO(@Magnus Dybdal): add proper documentation.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config TODO(@Magnus Dybdal): add proper documentation.
type Config struct {
	Addr   string
	DB_url string

	ChirpstackURL    string // chirpstack rest api url
	ChirpstackSecret string // chirpstack api token (bearer token)
}

// Load TODO(@Magnus Dybdal): add proper documentation.
func Load() *Config {
	secret := env.GetFile("/secrets/chirpstack-api-key")

	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		DB_url:           env.Get("DB_url", "postgres://device:device@device-db:5432/device?sslmode=disable"),
		ChirpstackURL:    env.Get("CHIRPSTACK_REST", "http://chirpstack-rest-api:8090"),
		ChirpstackSecret: secret,
	}
}
