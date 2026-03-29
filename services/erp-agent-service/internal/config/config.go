// Package config provides erp-agent-service runtime configuration.
package config

import (
	"innoveria-iot/pkg/env"
)

// Config holds erp-agent-service runtime configuration.
type Config struct {
	Addr string

	MonitorERPHost          string // Monitor ERP host address
	MonitorERPCompanyNumber string // Monitor ERP company number

	MonitorERPUsername string // Monitor ERP username
	MonitorERPPassword string // Monitor ERP password

	MonitorERPForceRelogin bool // Should be false; true logs out all active Monitor ERP sessions
	// EnableSwagger bool
}

// Load reads configuration from environment variables.
func Load() *Config {
	return &Config{
		Addr: ":" + env.Get("PORT", "8080"),

		MonitorERPHost:          env.Get("MONITOR_ERP_HOST", ""),
		MonitorERPCompanyNumber: env.Get("MONITOR_ERP_COMPANY_NUMBER", ""),
		MonitorERPUsername:      env.Get("MONITOR_ERP_USERNAME", ""),
		MonitorERPPassword:      env.Get("MONITOR_ERP_PASSWORD", ""),
		MonitorERPForceRelogin:  env.GetBool("MONITOR_ERP_FORCE_RELOGIN", false),
	}
}
