package config

import "innoveria-iot/pkg/env"

type Config struct {
	Addr string
}

func Load() *Config {
	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
	}
}
