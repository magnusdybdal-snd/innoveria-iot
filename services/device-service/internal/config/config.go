package config

import (
	"innoveria-iot/pkg/env"
	"log"
	"net/url"
)

type Config struct {
	Addr string

	ChirpstackURL    url.URL
	ChirpstackSecret string
}

func Load() *Config {
	secret := env.GetFile("/secrets/chirpstack-api-key")

	chripstackURL, err := url.Parse(env.Get("Chirpstack_REST", "http://chirpstack-rest-api:8090"))
	if err != nil {
		log.Fatal("config loading failed")
	}

	return &Config{
		Addr:             ":" + env.Get("PORT", "8080"),
		ChirpstackURL:    *chripstackURL,
		ChirpstackSecret: secret,
	}
}
