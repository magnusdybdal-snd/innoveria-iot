// Package repository provides data access methods for aggregation rules in the context service.
package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/dbutil"
)

const (
	getByCompanyID = `
	SELECT rule_id, company_id, name, context_type, measurement_type, aggregation_method, time_bucket_minutes, is_active, created_at, updated_at
	FROM context.aggregation_rule
	WHERE company_id = $1
`

	getByID = `
	SELECT rule_id, company_id, name, context_type, measurement_type, aggregation_method, time_bucket_minutes, is_active, created_at, updated_at
	FROM context.aggregation_rule
	WHERE rule_id = $1
`

	createRule = `
	INSERT INTO context.aggregation_rule (company_id, name, context_type, measurement_type, aggregation_method, time_bucket_minutes, is_active)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING rule_id
`
)

// RuleRepository provides methods to interact with the aggregation_rule table in the database.
type RuleRepository struct {
	db *dbutil.DB
}

// NewRuleRepository creates a new instance of RuleRepository with the given database connection.
func NewRuleRepository(db *dbutil.DB) *RuleRepository {
	return &RuleRepository{db: db}
}

// GetByCompanyID retrieves all aggregation rules for a given company ID.
func (r *RuleRepository) GetByCompanyID(ctx context.Context, companyID string) ([]domain.AggregationRule, error) {
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

// GetByID retrieves a single aggregation rule by its ID.
// Returns domain.ErrNotFound if no rule exists with that ID.
func (r *RuleRepository) GetByID(ctx context.Context, ruleID string) (domain.AggregationRule, error) {
	var rule domain.AggregationRule
	err := r.db.Pool.QueryRow(ctx, getByID, ruleID).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.AggregationRule{}, domain.ErrNotFound
		}
		return domain.AggregationRule{}, err
	}
	return rule, nil
}

// Create inserts a new aggregation rule into the database and returns the generated rule ID.
func (r *RuleRepository) Create(ctx context.Context, rule domain.AggregationRule) (string, error) {
	var ruleID string
	// Execute the insert query and scan the returned rule_id into the ruleID variable
	err := r.db.Pool.QueryRow(ctx, createRule,
		rule.CompanyID,
		rule.Name,
		rule.ContextType,
		rule.MeasurementType,
		rule.AggregationMethod,
		rule.TimeBucketMinutes,
		rule.IsActive,
	).Scan(&ruleID)
	if err != nil {
		var pgErr *pgconn.PgError
		// PostgreSQL error code 23505 is "unique_violation", raised when inserting a duplicate
		// (company_id, context_type) pair, which violates the unique constraint on the table.
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", domain.ErrConflict
		}
		return "", err
	}
	return ruleID, nil
}
