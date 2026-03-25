package calculators

import (
	"log/slog"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// voltageV is the assumed line voltage in volts.
// TODO: replace with per-device voltage from device-service when available.
const voltageV = 230.0

// WattHourCalculator computes energy consumption in Wh from instantaneous current readings.
// It expects payload[rule.MeasurementType] to be an RMS current value in amps (e.g. Milesight CT "current").
// Watts are derived as: watts = amps × voltageV.
// Energy per bucket is calculated using trapezoidal integration over actual reading timestamps.
type WattHourCalculator struct{}

// Calculate implements Calculator for context_type = "watt_over_time".
func (c *WattHourCalculator) Calculate(input Input) (domain.ContextData, error) {
	buckets := buildBuckets(input.From, input.To, input.BucketMinutes)
	lastIdx := len(buckets) - 1

	for i := range buckets {
		inBucket := readingsInBucket(input.Readings, buckets[i].PeriodStart, buckets[i].PeriodEnd, i == lastIdx)
		wh, err := trapezoidalWh(inBucket, input.Rule.MeasurementType)
		if err != nil {
			return domain.ContextData{}, err
		}
		buckets[i].Value = wh
	}

	total := 0.0
	for _, b := range buckets {
		total += b.Value
	}

	return domain.ContextData{
		CompanyID:    input.Rule.CompanyID,
		ContextType:  input.Rule.ContextType,
		Value:        total,
		Unit:         "Wh",
		PeriodStart:  input.From,
		PeriodEnd:    input.To,
		Buckets:      buckets,
		CalculatedAt: time.Now().UTC(),
	}, nil
}

// trapezoidalWh integrates instantaneous watt readings over time using the trapezoid rule.
// Returns 0 when fewer than two readings are available (no interval to integrate over).
func trapezoidalWh(readings []domain.MeasurementReading, measurementType string) (float64, error) {
	if len(readings) < 2 {
		slog.Warn("insufficient readings for trapezoidal integration, bucket yields 0 Wh", "count", len(readings), "measurement_type", measurementType)
		return 0.0, nil
	}

	energy := 0.0
	for i := 1; i < len(readings); i++ {
		raw0, err := extractFloat(readings[i-1], measurementType)
		if err != nil {
			return 0.0, err
		}
		raw1, err := extractFloat(readings[i], measurementType)
		if err != nil {
			return 0.0, err
		}
		w0 := raw0 * voltageV
		w1 := raw1 * voltageV
		dt := readings[i].Timestamp.Sub(readings[i-1].Timestamp).Hours()
		energy += (w0 + w1) / 2 * dt
	}
	return energy, nil
}
