// Package config provides auth-service runtime configuration.
package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"innoveria-iot/pkg/env"
)

// Config holds auth-service runtime configuration.
type Config struct {
	Addr               string
	DB_URL             string
	JWT_SECRET         string
	JWTIssuer          string
	JWTAccessTTL       time.Duration
	JWTRefreshTokenTTL time.Duration
	EnableSwagger      bool
}

// Load reads configuration from environment variables.
func Load() *Config {
	dbHost := env.Get("DB_HOST", "auth-db")
	dbPort := env.Get("DB_PORT", "5432")
	dbUser := env.Get("DB_USER", "auth")
	dbPassword := env.Get("DB_PASSWORD", "auth")
	dbName := env.Get("DB_NAME", "auth")
	sslmode := env.Get("DB_SSLMODE", "disable")

	jwtAccessTTL, err := time.ParseDuration(env.Get("JWT_ACCESS_TTL", "15m"))
	if err != nil {
		slog.Error("invalid JWT_ACCESS_TTL", "error", err)
		os.Exit(1)
	}

	jwtRefreshTTL, err := time.ParseDuration(env.Get("JWT_REFRESH_TTL", "720h")) // 30 days
	if err != nil {
		slog.Error("invalid JWT_REFRESH_TTL", "error", err)
		os.Exit(1)
	}

	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),
		DB_URL: fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=%s",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
			sslmode,
		),
		JWT_SECRET:         env.Get("JWT_SECRET", "secret"), // jwt secret laoding
		JWTIssuer:          env.Get("JWT_ISSUER", "auth-service"),
		JWTAccessTTL:       jwtAccessTTL,
		JWTRefreshTokenTTL: jwtRefreshTTL,
		EnableSwagger:      env.GetBool("ENABLE_SWAGGER", false),
	}
}
