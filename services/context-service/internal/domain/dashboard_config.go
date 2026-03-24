// Package domain defines the context-service core interfaces and domain contracts.
package domain

import (
	"encoding/json"
	"time"
)

// DashboardConfig is the domain model for dashboard config
type DashboardConfig struct {
	ID         string
	CompanyID  string
	UserID     string
	Name       string
	ConfigJSON json.RawMessage
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
