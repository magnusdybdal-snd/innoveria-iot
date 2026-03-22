// Package config TODO(@vinjar): add proper documentation.
package config

import (
	"log/slog"
	"os"
	"strings"

	"innoveria-iot/pkg/env"
)

// Config TODO(@vinjar): add proper documentation.
type Config struct {
	Addr               string
	CorsAllowedOrigins string
	CollSvcURL         string
	DeviceSvcURL       string
	AuthSvcURL         string
	OnboardingSvcURL   string
	JWTSecret          string
	JWTIssuer          string
	EnableSwagger      bool
}

// Load the spesific enviroment variables
func Load() *Config {
	jwtSecret, err := env.Required("JWT_SECRET")
	if err != nil {
		slog.Error("invalid jwt secret", "error", err)
		os.Exit(1)
	}
	cfg := Config{
		Addr:               ":" + env.Get("PORT", "8080"),
		CorsAllowedOrigins: env.Get("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
		CollSvcURL:         strings.TrimSpace(env.Get("COLLECTION_SERVICE", "http://collection-service:8080")),
		DeviceSvcURL:       strings.TrimSpace(env.Get("DEVICE_SERVICE", "http://device-service:8080")),
		AuthSvcURL:         strings.TrimSpace(env.Get("AUTH_SERVICE", "http://auth-service:8080")),
		OnboardingSvcURL:   strings.TrimSpace(env.Get("ONBOARDING_SERVICE", "http://onboarding-service:8080")),
		JWTSecret:          jwtSecret, // jwt secret laoding
		JWTIssuer:          env.Get("JWT_ISSUER", "auth-service"),
		EnableSwagger:      env.GetBool("ENABLE_SWAGGER", false),
	}

	return &cfg
}
