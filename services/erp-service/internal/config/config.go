// Package config provides erp-service runtime configuration.
package config

import (
	"fmt"
	"innoveria-iot/pkg/env"
	"time"
)

const defaultReconcileInterval = 30 * time.Second

// Config holds erp-service runtime configuration.
type Config struct {
	Addr  string
	DBURL string

	ReconcileInterval time.Duration
	// EnableSwagger bool
}

// Load reads configuration from environment variables.
func Load() *Config {
	dbHost := env.Get("DB_HOST", "erp-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "erp")
	dbPassword := env.Get("DB_PASSWORD", "erp")
	dbName := env.Get("DB_NAME", "erp")
	sslmode := env.Get("DB_SSLMODE", "disable")

	reconcileIntervalRaw := env.Get("ERP_RECONCILE_INTERVAL", "30s")
	reconcileInterval, err := time.ParseDuration(reconcileIntervalRaw)
	if err != nil || reconcileInterval <= 0 {
		reconcileInterval = defaultReconcileInterval
	}

	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
		DBURL: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
			sslmode,
		),
		ReconcileInterval: reconcileInterval,
	}
}
