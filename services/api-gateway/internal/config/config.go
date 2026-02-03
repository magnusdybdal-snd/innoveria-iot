package config

import "innoveria-iot/pkg/env"

type Config struct {
	Addr string
	// Some serviceURL
	// Some serviceURL
	// Some serviceAPIKEY?
}

// Loads the spesific enviroment variables
func Load() *Config {
	cfg := Config{
		Addr: ":" + env.Get("PORT", "8080"),
		// Some serviceURL
	}

	// TODO: Handle edge cases, if env is empty

	return &cfg
}
