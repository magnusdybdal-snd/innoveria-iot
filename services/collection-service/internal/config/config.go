package config

import "innoveria-iot/pkg/env"

type Config struct {
	Addr string
}

// Loads the spesific enviroment variables
func Load() *Config {
	cfg := Config{
		Addr: ":" + env.Get("PORT", "8080"),
	}

	return &cfg
}
