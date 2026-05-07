// Package csvwriter provides helpers for rendering structured export data as CSV.
package csvwriter

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"strconv"
	"time"
)

// Column describes a single CSV column: the key used to look up payload values
// and the human-readable header to write in the file.
type Column struct {
	PayloadKey string
	Header     string
}

// Row represents a single measurement row with its timestamp and payload.
type Row struct {
	Timestamp time.Time
	Payload   map[string]any
}

// Write renders columns and rows as CSV into w.
// The first two columns are always date and local time; remaining columns follow the provided order.
// Returns any I/O error that occurred during writing or flushing.
func Write(w io.Writer, columns []Column, rows []Row, loc *time.Location) error {
	cw := csv.NewWriter(w)

	header := []string{"date", "time (" + loc.String() + ")"}
	for _, col := range columns {
		header = append(header, col.Header)
	}
	cw.Write(header) //nolint:errcheck

	for _, row := range rows {
		t := row.Timestamp.In(loc)
		record := []string{
			t.Format(time.DateOnly),
			t.Format("15:04:05"),
		}
		for _, col := range columns {
			record = append(record, formatValue(row.Payload[col.PayloadKey]))
		}
		cw.Write(record) //nolint:errcheck
	}

	cw.Flush()
	return cw.Error()
}

// formatValue formats a JSON payload value for CSV output.
// Float64 values are rounded to 4 decimal places and trailing zeros are stripped.
// All other types use their default string representation.
func formatValue(v any) string {
	f, ok := v.(float64)
	if !ok {
		return fmt.Sprintf("%v", v)
	}
	rounded := math.Round(f*10000) / 10000
	return strconv.FormatFloat(rounded, 'f', -1, 64)
}
