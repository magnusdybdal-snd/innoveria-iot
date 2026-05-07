package domain

import (
	"context"
	"time"
)

// ExportColumn describes a single CSV column: the payload key used to look up values
// and the human-readable header to write in the CSV file.
type ExportColumn struct {
	PayloadKey string
	Header     string
}

// ExportData is the structured result returned by ExportService, ready for CSV rendering.
type ExportData struct {
	Columns      []ExportColumn
	Measurements []SensorMeasurement
}

// ExportService defines the business logic for preparing sensor measurement exports.
type ExportService interface {
	GetExportData(ctx context.Context, companyID, deviceEUI string, from, to time.Time) (ExportData, error)
}
