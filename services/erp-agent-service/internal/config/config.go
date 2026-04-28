// Package config provides erp-agent-service runtime configuration.
package config

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"innoveria-iot/pkg/env"
)

// Config holds erp-agent-service runtime configuration.
type Config struct {
	Addr     string
	GOEnv    string
	JWTToken string

	UseMockMonitor bool

	MonitorERPHost          string // Monitor ERP host address
	MonitorERPPort          string // Monitor ERP port
	MonitorERPCompanyNumber int    // Monitor ERP company number, (this is 1 by default)

	MonitorERPUsername string // Monitor ERP username
	MonitorERPPassword string // Monitor ERP password

	MonitorERPForceRelogin bool // Should be false; true logs out all active Monitor ERP sessions

	PollingInterval time.Duration
	CycleTimeout    time.Duration
	MaxBackoffTime  time.Duration
	ErpSvcURL       string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	goEnv, err := env.Required("GO_ENV")
	if err != nil {
		return nil, err
	}

	var jwtToken string
	if goEnv == "development" {
		jwtToken = "dev"
	} else {
		jwtToken, err = env.Required("JWT_TOKEN")
		if err != nil {
			return nil, err
		}
	}

	useMockMonitor := env.GetBool("MOCK_MONITOR", goEnv == "development")
	if useMockMonitor {
		slog.Warn("MOCK_MONITOR is enabled; using mock Monitor ERP integration")
	}

	if useMockMonitor && goEnv != "development" {
		return nil, fmt.Errorf("mock monitor is only allowed in development")
	}

	host := env.Get("MONITOR_ERP_HOST", "")
	monitorPort := env.Get("MONITOR_ERP_PORT", "")
	companyRaw := env.Get("MONITOR_ERP_COMPANY_NUMBER", "0")
	username := env.Get("MONITOR_ERP_USERNAME", "")
	password := env.Get("MONITOR_ERP_PASSWORD", "")

	if !useMockMonitor {
		var err error
		host, err = env.Required("MONITOR_ERP_HOST")
		if err != nil {
			return nil, err
		}

		monitorPort, err = env.Required("MONITOR_ERP_PORT")
		if err != nil {
			return nil, err
		}

		companyRaw, err = env.Required("MONITOR_ERP_COMPANY_NUMBER")
		if err != nil {
			return nil, err
		}

		username, err = env.Required("MONITOR_ERP_USERNAME")
		if err != nil {
			return nil, err
		}

		password, err = env.Required("MONITOR_ERP_PASSWORD")
		if err != nil {
			return nil, err
		}

	}

	companyNumber, err := strconv.Atoi(companyRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid MONITOR_ERP_COMPANY_NUMBER: %w", err)
	}

	pollingInterval, err := time.ParseDuration(env.Get("POLLING_INTERVAL", "10m"))
	if err != nil {
		return nil, fmt.Errorf("invalid POLLING_INTERVAL: %w", err)
	}

	cycleTimeout, err := time.ParseDuration(env.Get("CYCLE_TIMEOUT", "2m"))
	if err != nil {
		return nil, fmt.Errorf("invalid CYCLE_TIMEOUT: %w", err)
	}

	maxBackoffTime, err := time.ParseDuration(env.Get("MAX_BACKOFF_TIME", "5m"))
	if err != nil {
		return nil, fmt.Errorf("invalid MAX_BACKOFF_TIME: %w", err)
	}

	const defaultERPSvcURL = "http://erp-service:8080"
	erpSvcURL := env.Get("ERP_SERVICE", "")
	if erpSvcURL == "" {
		erpSvcURL = defaultERPSvcURL
		slog.Warn("ERP_SERVICE is not set; using default ERP service URL", "erp_service_url", erpSvcURL)
	}

	return &Config{
		Addr:                    ":" + env.Get("PORT", "8080"),
		GOEnv:                   goEnv,
		JWTToken:                jwtToken,
		UseMockMonitor:          useMockMonitor,
		MonitorERPHost:          host,
		MonitorERPPort:          monitorPort,
		MonitorERPCompanyNumber: companyNumber,
		MonitorERPUsername:      username,
		MonitorERPPassword:      password,
		MonitorERPForceRelogin:  env.GetBool("MONITOR_ERP_FORCE_RELOGIN", false),
		PollingInterval:         pollingInterval,
		CycleTimeout:            cycleTimeout,
		MaxBackoffTime:          maxBackoffTime,
		ErpSvcURL:               erpSvcURL,
	}, nil
}
