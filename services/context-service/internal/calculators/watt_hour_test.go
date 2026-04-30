package calculators

import (
	"testing"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// wattReading is a test helper that builds a MeasurementReading with a "current" payload value.
func wattReading(amps float64, t time.Time) domain.MeasurementReading {
	return domain.MeasurementReading{
		DeviceEUI: "dev1",
		Timestamp: t,
		Payload:   map[string]any{"current": amps},
	}
}

// TestTrapezoidalWh_TwoReadings verifies the basic trapezoid calculation for two readings.
// With 1A at t=0 and 1A at t=1h at 230V: energy = (230 + 230) / 2 * 1h = 230 Wh.
func TestTrapezoidalWh_TwoReadings(t *testing.T) {
	t0 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(1 * time.Hour)

	readings := []domain.MeasurementReading{
		wattReading(1.0, t0),
		wattReading(1.0, t1),
	}

	got, err := trapezoidalWh(readings, "current", 230.0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 230.0 {
		t.Errorf("expected 230.0 Wh, got %v", got)
	}
}

// TestTrapezoidalWh_VaryingAmps verifies trapezoidal integration with different start and end currents.
// With 2A at t=0 and 4A at t=1h at 230V: energy = (460 + 920) / 2 * 1h = 690 Wh.
func TestTrapezoidalWh_VaryingAmps(t *testing.T) {
	t0 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(1 * time.Hour)

	readings := []domain.MeasurementReading{
		wattReading(2.0, t0),
		wattReading(4.0, t1),
	}

	got, err := trapezoidalWh(readings, "current", 230.0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 690.0 {
		t.Errorf("expected 690.0 Wh, got %v", got)
	}
}

// TestTrapezoidalWh_400V verifies that 400V line voltage produces the correct result.
// With 1A at t=0 and 1A at t=1h at 400V: energy = (400 + 400) / 2 * 1h = 400 Wh.
func TestTrapezoidalWh_400V(t *testing.T) {
	t0 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(1 * time.Hour)

	readings := []domain.MeasurementReading{
		wattReading(1.0, t0),
		wattReading(1.0, t1),
	}

	got, err := trapezoidalWh(readings, "current", 400.0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 400.0 {
		t.Errorf("expected 400.0 Wh, got %v", got)
	}
}

// TestTrapezoidalWh_InsufficientReadings verifies that fewer than two readings returns 0 Wh.
func TestTrapezoidalWh_InsufficientReadings(t *testing.T) {
	readings := []domain.MeasurementReading{
		wattReading(1.0, time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)),
	}

	got, err := trapezoidalWh(readings, "current", 230.0)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0.0 {
		t.Errorf("expected 0.0 Wh, got %v", got)
	}
}
