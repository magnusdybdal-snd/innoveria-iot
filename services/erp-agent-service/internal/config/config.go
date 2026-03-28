// Package config provides erp-agent-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
	"strconv"
)

// Config holds erp-agent-service runtime configuration.
type Config struct {
	Addr string

	MonitorERPHost          string // Monitor ERP host address
	MonitorERPCompanyNumber int    // Monitor ERP company number, (this is 1 by default)

	MonitorERPUsername string // Monitor ERP username
	MonitorERPPassword string // Monitor ERP password

	MonitorERPForceRelogin bool // Should be false; true logs out all active Monitor ERP sessions
	// EnableSwagger bool
}

// Load reads configuration from environment variables.
func Load() *Config {

	companyNumber, err := strconv.Atoi(env.Get("MONITOR_ERP_COMPANY_NUMBER", "1"))
	if err != nil {
		companyNumber = 1 // defaults to 1
	}

	return &Config{
		Addr:                    ":" + env.Get("PORT", "8080"),
		MonitorERPHost:          env.Get("MONITOR_ERP_HOST", ""),
		MonitorERPCompanyNumber: companyNumber,
		MonitorERPUsername:      env.Get("MONITOR_ERP_USERNAME", ""),
		MonitorERPPassword:      env.Get("MONITOR_ERP_PASSWORD", ""),
		MonitorERPForceRelogin:  env.GetBool("MONITOR_ERP_FORCE_RELOGIN", false),
	}
}
