// Package config provides erp-agent-service runtime configuration.
package config

import (
	"fmt"
	"strconv"
	"time"

	"innoveria-iot/pkg/env"
)

// Config holds erp-agent-service runtime configuration.
type Config struct {
	Addr string

	MonitorERPHost          string // Monitor ERP host address
	MonitorERPPORT          string // Monitor ERP port
	MonitorERPCompanyNumber int    // Monitor ERP company number, (this is 1 by default)

	MonitorERPUsername string // Monitor ERP username
	MonitorERPPassword string // Monitor ERP password

	MonitorERPForceRelogin bool // Should be false; true logs out all active Monitor ERP sessions

	PollingInterval time.Duration
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	host, err := env.Required("MONITOR_ERP_HOST")
	if err != nil {
		return nil, err
	}
	monitorPort, err := env.Required("MONITOR_ERP_PORT")
	if err != nil {
		return nil, err
	}

	companyRaw, err := env.Required("MONITOR_ERP_COMPANY_NUMBER")
	if err != nil {
		return nil, err
	}

	companyNumber, err := strconv.Atoi(companyRaw)
	if err != nil {
		return nil, fmt.Errorf("invalid MONITOR_ERP_COMPANY_NUMBER: %w", err)
	}

	username, err := env.Required("MONITOR_ERP_USERNAME")
	if err != nil {
		return nil, err
	}

	password, err := env.Required("MONITOR_ERP_PASSWORD")
	if err != nil {
		return nil, err
	}

	pollingInterval, err := time.ParseDuration(env.Get("POLLING_INTERVAL", "10m"))
	if err != nil {
		return nil, fmt.Errorf("invalid POLLING_INTERVAL: %w", err)
	}

	return &Config{
		Addr:                    ":" + env.Get("PORT", "8080"),
		MonitorERPHost:          host,
		MonitorERPPORT:          monitorPort,
		MonitorERPCompanyNumber: companyNumber,
		MonitorERPUsername:      username,
		MonitorERPPassword:      password,
		MonitorERPForceRelogin:  env.GetBool("MONITOR_ERP_FORCE_RELOGIN", false),
		PollingInterval:         pollingInterval,
	}, nil
}
