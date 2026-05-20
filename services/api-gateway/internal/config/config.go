// Package config loads api-gateway configuration from environment variables.
package config

import (
	"log/slog"
	"os"
	"strings"

	"innoveria-iot/pkg/env"
)

// Config holds all configuration values for the API gateway.
type Config struct {
	Addr               string
	CollSvcURL         string
	DeviceSvcURL       string
	AuthSvcURL         string
	OnboardingSvcURL   string
	ContextSvcURL      string
	ErpSvcURL          string
	CorsAllowedOrigins string
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
		CollSvcURL:         strings.TrimSpace(env.Get("COLLECTION_SERVICE", "http://collection-service:8080")),
		DeviceSvcURL:       strings.TrimSpace(env.Get("DEVICE_SERVICE", "http://device-service:8080")),
		AuthSvcURL:         strings.TrimSpace(env.Get("AUTH_SERVICE", "http://auth-service:8080")),
		OnboardingSvcURL:   strings.TrimSpace(env.Get("ONBOARDING_SERVICE", "http://onboarding-service:8080")),
		ContextSvcURL:      strings.TrimSpace(env.Get("CONTEXT_SERVICE", "http://context-service:8080")),
		ErpSvcURL:          strings.TrimSpace(env.Get("ERP_SERVICE", "http://erp-service:8080")),
		CorsAllowedOrigins: env.Get("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
		JWTSecret:          jwtSecret, // jwt secret laoding
		JWTIssuer:          env.Get("JWT_ISSUER", "auth-service"),
		EnableSwagger:      env.GetBool("ENABLE_SWAGGER", false),
	}

	return &cfg
}
