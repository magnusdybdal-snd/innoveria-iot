// Package repository provides data access methods for aggregation rules in the context service.
package repository

import (
	"context"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	getByCompanyID = `
	SELECT rule_id, company_id, name, context_type, measurement_type, aggregation_method, time_bucket_minutes, is_active, created_at, updated_at
	FROM context.aggregation_rule 
	WHERE company_id = $1
`
)

// RuleMappingRepository provides methods to interact with the aggregation_rule table in the database.
type RuleMappingRepository struct {
	db *dbutil.DB
}

// NewRuleMappingRepository creates a new instance of RuleMappingRepository with the given database connection.
func NewRuleMappingRepository(db *dbutil.DB) *RuleMappingRepository {
	return &RuleMappingRepository{db: db}
}

// GetByCompanyID retrieves all aggregation rules for a given company ID.
func (r *RuleMappingRepository) GetByCompanyID(ctx context.Context, companyID string) ([]domain.AggregationRule, error) {
	rows, err := r.db.Pool.Query(ctx, getByCompanyID, companyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rules := []domain.AggregationRule{}
	for rows.Next() {
		var rule domain.AggregationRule
		err := rows.Scan(
			&rule.ID,
			&rule.CompanyID,
			&rule.Name,
			&rule.ContextType,
			&rule.MeasurementType,
			&rule.AggregationMethod,
			&rule.TimeBucketMinutes,
			&rule.IsActive,
			&rule.CreatedAt,
			&rule.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}
