package calculators

import (
	"testing"

	"innoveria-iot/context-service/internal/domain"
)

// reading is a test helper that builds a MeasurementReading with a single payload key/value.
func reading(key string, val float64) domain.MeasurementReading {
	return domain.MeasurementReading{
		DeviceEUI: "dev1",
		Payload:   map[string]any{key: val},
	}
}

// TestAggregate_AVG verifies that AVG returns the mean of all reading values.
func TestAggregate_AVG(t *testing.T) {
	readings := []domain.MeasurementReading{reading("v", 2.0), reading("v", 4.0)}

	got, err := aggregate(readings, "v", "AVG")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 3.0 {
		t.Errorf("expected 3.0, got %v", got)
	}
}

// TestAggregate_SUM verifies that SUM returns the total of all reading values.
func TestAggregate_SUM(t *testing.T) {
	readings := []domain.MeasurementReading{reading("v", 1.0), reading("v", 2.0), reading("v", 3.0)}

	got, err := aggregate(readings, "v", "SUM")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 6.0 {
		t.Errorf("expected 6.0, got %v", got)
	}
}

// TestAggregate_MIN verifies that MIN returns the smallest reading value.
func TestAggregate_MIN(t *testing.T) {
	readings := []domain.MeasurementReading{reading("v", 3.0), reading("v", 1.0), reading("v", 2.0)}

	got, err := aggregate(readings, "v", "MIN")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 1.0 {
		t.Errorf("expected 1.0, got %v", got)
	}
}

// TestAggregate_MAX verifies that MAX returns the largest reading value.
func TestAggregate_MAX(t *testing.T) {
	readings := []domain.MeasurementReading{reading("v", 3.0), reading("v", 1.0), reading("v", 2.0)}

	got, err := aggregate(readings, "v", "MAX")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 3.0 {
		t.Errorf("expected 3.0, got %v", got)
	}
}

// TestAggregate_EmptyReadings verifies that an empty slice returns 0.0 without error.
func TestAggregate_EmptyReadings(t *testing.T) {
	got, err := aggregate(nil, "v", "SUM")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 0.0 {
		t.Errorf("expected 0.0, got %v", got)
	}
}

// TestAggregate_UnknownMethod verifies that an unsupported aggregation method returns an error.
func TestAggregate_UnknownMethod(t *testing.T) {
	readings := []domain.MeasurementReading{reading("v", 1.0)}

	_, err := aggregate(readings, "v", "MEDIAN")

	if err == nil {
		t.Fatal("expected error for unsupported method, got nil")
	}
}
