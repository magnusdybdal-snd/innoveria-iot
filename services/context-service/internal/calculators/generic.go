package calculators

import (
	"math"
	"time"

	"innoveria-iot/context-service/internal/domain"
)

// GenericAggregationCalculator handles context types that extract a single payload field
// and apply a standard aggregation method (AVG, SUM, MIN, MAX) per time bucket.
// It is used as a fallback when no specialised Calculator is registered for a context type.
type GenericAggregationCalculator struct{}

// Calculate implements Calculator using rule.AggregationMethod and rule.MeasurementType.
func (c *GenericAggregationCalculator) Calculate(input Input) (domain.ContextData, error) {
	buckets := buildBuckets(input.From, input.To, input.BucketMinutes)
	lastIdx := len(buckets) - 1

	for i := range buckets {
		inBucket := readingsInBucket(input.Readings, buckets[i].PeriodStart, buckets[i].PeriodEnd, i == lastIdx)
		buckets[i].Value = aggregate(inBucket, input.Rule.MeasurementType, input.Rule.AggregationMethod)
	}

	total := 0.0
	for _, b := range buckets {
		total += b.Value
	}

	return domain.ContextData{
		CompanyID:    input.Rule.CompanyID,
		ContextType:  input.Rule.ContextType,
		Value:        total,
		Unit:         input.Rule.MeasurementType,
		PeriodStart:  input.From,
		PeriodEnd:    input.To,
		Buckets:      buckets,
		CalculatedAt: time.Now().UTC(),
	}, nil
}

// aggregate applies the given method to the extracted values from readings.
// Returns 0.0 for empty slices.
func aggregate(readings []domain.MeasurementReading, key, method string) float64 {
	if len(readings) == 0 {
		return 0.0
	}

	values := make([]float64, len(readings))
	for i, r := range readings {
		values[i] = extractFloat(r, key)
	}

	switch method {
	case "AVG":
		sum := 0.0
		for _, v := range values {
			sum += v
		}
		return sum / float64(len(values))
	case "SUM":
		sum := 0.0
		for _, v := range values {
			sum += v
		}
		return sum
	case "MIN":
		min := math.MaxFloat64
		for _, v := range values {
			if v < min {
				min = v
			}
		}
		return min
	case "MAX":
		max := -math.MaxFloat64
		for _, v := range values {
			if v > max {
				max = v
			}
		}
		return max
	default:
		return 0.0
	}
}
