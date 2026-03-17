package domain

import "time"

// ContextData is the domain model for contextdata
type ContextData struct {
	ID                   string
	CompanyID            string
	ContextType          string
	OrderID              string
	ProductionResourceID string
	DeviceEUI            string
	Value                float64
	Unit                 string
	PeriodStart          time.Time
	PeriodEnd            time.Time
	CalculatedAt         time.Time
}
