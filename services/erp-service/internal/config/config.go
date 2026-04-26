// Package config provides erp-service runtime configuration.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"innoveria-iot/pkg/env"
)

const defaultReconcileInterval = 30 * time.Second

// Config holds erp-service runtime configuration.
type Config struct {
	Addr  string
	DBURL string

	ReconcileInterval    time.Duration
	ERP_AGENT_JWT_SECRET string
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
	if err != nil {
		slog.Warn("ERP_RECONCILE_INTERVAL is not a valid duration, using default",
			"raw_value", reconcileIntervalRaw,
			"default", defaultReconcileInterval)
		reconcileInterval = defaultReconcileInterval
	} else if reconcileInterval <= 0 {
		slog.Warn("ERP_RECONCILE_INTERVAL must be positive, using default",
			"parsed_value", reconcileInterval,
			"default", defaultReconcileInterval)
		reconcileInterval = defaultReconcileInterval
	}

	erpAgentSecret, err := env.Required("ERP_AGENT_JWT_SECRET")
	if err != nil {
		slog.Error("invalid erp agent jwt secret", "error", err)
		os.Exit(1)
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
		ReconcileInterval:    reconcileInterval,
		ERP_AGENT_JWT_SECRET: erpAgentSecret,
	}
}
