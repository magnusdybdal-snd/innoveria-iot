package config

import "innoveria-iot/pkg/env"

type Config struct {
	Port string
	// Some serviceURL
	// Some serviceURL
	// Some serviceAPIKEY?
}

// Loads the spesific enviroment variables
func Load() Config {
	cfg := Config{
		Port: env.Get("PORT", "8080"),
		// Some serviceURL
	}

	// TODO: Handle edge cases, if env is empty

	return cfg
}
