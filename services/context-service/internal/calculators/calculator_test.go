package calculators

import (
	"testing"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// TestBuildBuckets_DividesWindowIntoEqualSlices verifies that a window divisible by bucketMins
// produces the correct number of equal buckets with aligned start and end times.
func TestBuildBuckets_DividesWindowIntoEqualSlices(t *testing.T) {
	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC) // 60-minute window

	buckets := buildBuckets(from, to, 20) // expect 3 buckets of 20 min each

	if len(buckets) != 3 {
		t.Fatalf("expected 3 buckets, got %d", len(buckets))
	}
	if !buckets[0].PeriodStart.Equal(from) {
		t.Errorf("first bucket start: got %v, want %v", buckets[0].PeriodStart, from)
	}
	if !buckets[2].PeriodEnd.Equal(to) {
		t.Errorf("last bucket end: got %v, want %v", buckets[2].PeriodEnd, to)
	}
}

// TestReadingsInBucket_ReturnsReadingsWithinRange verifies that only readings within [start, end)
// are returned, and that a reading exactly at end is included when isLast is true.
func TestReadingsInBucket_ReturnsReadingsWithinRange(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 1, 0, 30, 0, 0, time.UTC)

	readings := []domain.MeasurementReading{
		{DeviceEUI: "a", Timestamp: start},                       // at start — included
		{DeviceEUI: "b", Timestamp: start.Add(15 * time.Minute)}, // inside — included
		{DeviceEUI: "c", Timestamp: end},                         // at end, isLast=true — included
		{DeviceEUI: "d", Timestamp: end.Add(1 * time.Minute)},    // after end — excluded
		{DeviceEUI: "e", Timestamp: start.Add(-1 * time.Minute)}, // before start — excluded
	}

	got := readingsInBucket(readings, start, end, true)

	if len(got) != 3 {
		t.Fatalf("expected 3 readings, got %d", len(got))
	}
}

// TestExtractFloat_ReturnsValueForKnownKey verifies that a float64 value is correctly extracted
// from a measurement payload by key.
func TestExtractFloat_ReturnsValueForKnownKey(t *testing.T) {
	r := domain.MeasurementReading{
		DeviceEUI: "dev1",
		Payload:   map[string]any{"power": float64(42.5)},
	}

	val, err := extractFloat(r, "power")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 42.5 {
		t.Errorf("expected 42.5, got %v", val)
	}
}

// TestBuildBuckets_ZeroBucketMins verifies that a non-positive bucketMins returns a single bucket
// spanning the entire window.
func TestBuildBuckets_ZeroBucketMins(t *testing.T) {
	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 1, 1, 1, 0, 0, 0, time.UTC)

	buckets := buildBuckets(from, to, 0)

	if len(buckets) != 1 {
		t.Fatalf("expected 1 bucket, got %d", len(buckets))
	}
	if !buckets[0].PeriodStart.Equal(from) || !buckets[0].PeriodEnd.Equal(to) {
		t.Errorf("bucket should span full window: got [%v, %v]", buckets[0].PeriodStart, buckets[0].PeriodEnd)
	}
}

// TestBuildBuckets_UnevenWindow verifies that the last bucket is clamped to `to` when the window
// does not divide evenly by bucketMins.
func TestBuildBuckets_UnevenWindow(t *testing.T) {
	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 1, 1, 0, 50, 0, 0, time.UTC) // 50-minute window, 20-min buckets → 3 buckets

	buckets := buildBuckets(from, to, 20)

	if len(buckets) != 3 {
		t.Fatalf("expected 3 buckets, got %d", len(buckets))
	}
	if !buckets[2].PeriodEnd.Equal(to) {
		t.Errorf("last bucket end should be clamped to %v, got %v", to, buckets[2].PeriodEnd)
	}
}

// TestReadingsInBucket_ExcludesEndWhenNotLast verifies that a reading exactly at the bucket end
// is excluded when isLast is false.
func TestReadingsInBucket_ExcludesEndWhenNotLast(t *testing.T) {
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 1, 0, 30, 0, 0, time.UTC)

	readings := []domain.MeasurementReading{
		{DeviceEUI: "a", Timestamp: end}, // exactly at end, isLast=false — excluded
	}

	got := readingsInBucket(readings, start, end, false)

	if len(got) != 0 {
		t.Fatalf("expected 0 readings, got %d", len(got))
	}
}

// TestExtractFloat_MissingKey verifies that an error is returned when the payload key does not exist.
func TestExtractFloat_MissingKey(t *testing.T) {
	r := domain.MeasurementReading{
		DeviceEUI: "dev1",
		Payload:   map[string]any{},
	}

	_, err := extractFloat(r, "power")

	if err == nil {
		t.Fatal("expected error for missing key, got nil")
	}
}

// TestExtractFloat_IntValue verifies that an int payload value is correctly converted to float64.
func TestExtractFloat_IntValue(t *testing.T) {
	r := domain.MeasurementReading{
		DeviceEUI: "dev1",
		Payload:   map[string]any{"power": int(10)},
	}

	val, err := extractFloat(r, "power")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 10.0 {
		t.Errorf("expected 10.0, got %v", val)
	}
}

// TestExtractFloat_UnsupportedType verifies that an error is returned for a non-numeric payload value.
func TestExtractFloat_UnsupportedType(t *testing.T) {
	r := domain.MeasurementReading{
		DeviceEUI: "dev1",
		Payload:   map[string]any{"power": "not-a-number"},
	}

	_, err := extractFloat(r, "power")

	if err == nil {
		t.Fatal("expected error for unsupported type, got nil")
	}
}
