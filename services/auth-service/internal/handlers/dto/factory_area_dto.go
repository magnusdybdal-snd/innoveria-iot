package dto

import (
	"time"

	"innoveria-iot/auth-service/internal/domain"
)

// FactoryAreaResponse is the response payload for factory area operations.
type FactoryAreaResponse struct {
	ID          string  `json:"id"`
	FactoryID   string  `json:"factory_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

// FactoryAreaListResponse wraps a list of factory areas with total count.
type FactoryAreaListResponse struct {
	TotalCount   int                   `json:"total_count"`
	FactoryAreas []FactoryAreaResponse `json:"factory_areas"`
}

// CreateNewFactoryArea is the request payload for creating a factory area.
type CreateNewFactoryArea struct {
	FactoryID   string  `json:"factory_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// MapCreateFactoryAreaToDomain maps a create-factory-area DTO to domain FactoryArea.
func MapCreateFactoryAreaToDomain(from CreateNewFactoryArea) domain.FactoryArea {
	return domain.FactoryArea{
		FactoryID:   from.FactoryID,
		Name:        from.Name,
		Description: from.Description,
	}
}

// MapFactoryAreaFromDomain maps domain FactoryArea to a FactoryAreaResponse DTO.
func MapFactoryAreaFromDomain(from domain.FactoryArea) FactoryAreaResponse {
	return FactoryAreaResponse{
		ID:          from.ID,
		FactoryID:   from.FactoryID,
		Name:        from.Name,
		Description: from.Description,
		CreatedAt:   from.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   from.UpdatedAt.Format(time.RFC3339),
	}
}

// MapFactoryAreaListFromDomain maps domain factory areas to list response DTO.
func MapFactoryAreaListFromDomain(from []domain.FactoryArea) FactoryAreaListResponse {
	areas := make([]FactoryAreaResponse, 0, len(from))
	for _, area := range from {
		areas = append(areas, MapFactoryAreaFromDomain(area))
	}

	return FactoryAreaListResponse{
		TotalCount:   len(areas),
		FactoryAreas: areas,
	}
}
