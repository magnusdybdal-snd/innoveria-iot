package config

import (
	"innoveria-iot/pkg/env"
)

type Config struct {
	Addr string

	ChirpstackURL    string
	ChirpstackSecret string
}

func Load() *Config {
	secret := env.GetFile("/secrets/chirpstack-api-key")

	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		ChirpstackURL:    env.Get("Chirpstack_REST", "http://chirpstack-rest-api:8090"),
		ChirpstackSecret: secret,
	}
}
