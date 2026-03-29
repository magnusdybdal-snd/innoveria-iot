package dto

import "innoveria-iot/context-service/internal/domain"

// MapContextDataDomainToDTO maps a slice of domain ContextData to a slice of ContextDataResponse.
func MapContextDataDomainToDTO(from []domain.ContextData) []ContextDataResponse {
	out := make([]ContextDataResponse, len(from))
	for i, d := range from {
		out[i] = mapContextData(d)
	}
	return out
}

func mapContextData(from domain.ContextData) ContextDataResponse {
	buckets := make([]BucketResponse, len(from.Buckets))
	for i, b := range from.Buckets {
		buckets[i] = BucketResponse{
			PeriodStart: b.PeriodStart,
			PeriodEnd:   b.PeriodEnd,
			Value:       b.Value,
		}
	}
	return ContextDataResponse{
		DeviceEUI:    from.DeviceEUI,
		CompanyID:    from.CompanyID,
		ContextType:  from.ContextType,
		Unit:         from.Unit,
		PeriodStart:  from.PeriodStart,
		PeriodEnd:    from.PeriodEnd,
		TotalValue:   from.Value,
		Buckets:      buckets,
		CalculatedAt: from.CalculatedAt,
	}
}
