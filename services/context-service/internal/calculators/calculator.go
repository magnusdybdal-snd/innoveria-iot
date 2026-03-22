// Package calculators provides context-type-specific computation logic for the context service.
package calculators

import (
	"log/slog"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// Input holds all data a Calculator needs to compute context data for a single device.
type Input struct {
	Rule          domain.AggregationRule
	Readings      []domain.MeasurementReading
	BucketMinutes int
	From          time.Time
	To            time.Time
}

// Calculator computes a ContextData result from raw sensor readings.
// Each context type that requires specialised logic gets its own implementation.
// Types without a registered calculator fall back to GenericAggregationCalculator.
type Calculator interface {
	Calculate(input Input) (domain.ContextData, error)
}

// NewRegistry returns the map of explicitly registered context-type calculators.
// Add new entries here when a context type needs specialised computation logic.
func NewRegistry() map[string]Calculator {
	return map[string]Calculator{
		"watt_over_time": &WattHourCalculator{},
	}
}

// buildBuckets divides the time window [from, to] into equal slices of bucketMins minutes.
// The last bucket's end is clamped to `to`. Returns a single bucket spanning the full
// window when bucketMins is zero or negative.
func buildBuckets(from, to time.Time, bucketMins int) []domain.ContextBucket {
	if bucketMins <= 0 {
		return []domain.ContextBucket{{PeriodStart: from, PeriodEnd: to}}
	}

	step := time.Duration(bucketMins) * time.Minute
	var buckets []domain.ContextBucket
	start := from
	for start.Before(to) {
		end := start.Add(step)
		if end.After(to) {
			end = to
		}
		buckets = append(buckets, domain.ContextBucket{PeriodStart: start, PeriodEnd: end})
		start = end
	}
	return buckets
}

// readingsInBucket returns all readings whose timestamp falls within [start, end).
// The final bucket uses an inclusive end so no reading at exactly `to` is dropped.
func readingsInBucket(readings []domain.MeasurementReading, start, end time.Time, isLast bool) []domain.MeasurementReading {
	var out []domain.MeasurementReading
	for _, r := range readings {
		afterStart := !r.Timestamp.Before(start)
		beforeEnd := r.Timestamp.Before(end)
		atEnd := r.Timestamp.Equal(end)
		if afterStart && (beforeEnd || (isLast && atEnd)) {
			out = append(out, r)
		}
	}
	return out
}

// extractFloat retrieves a numeric value from a measurement payload by key.
// Returns 0.0 and logs a warning if the key is missing or has an unexpected type.
func extractFloat(r domain.MeasurementReading, key string) float64 {
	raw, ok := r.Payload[key]
	if !ok {
		slog.Warn("payload key not found", "device_eui", r.DeviceEUI, "key", key)
		return 0.0
	}
	switch v := raw.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	default:
		slog.Warn("unexpected payload value type", "device_eui", r.DeviceEUI, "key", key, "type", v)
		return 0.0
	}
}
