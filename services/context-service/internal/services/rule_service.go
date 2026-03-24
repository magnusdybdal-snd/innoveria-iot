package services

import (
	"context"
	"innoveria-iot/context-service/internal/domain"
)

// RuleServiceImpl implements business logic related to aggregation rules in the context service.
type RuleServiceImpl struct {
	repo domain.RuleRepository
}

// NewRuleServiceImpl creates a new instance of RuleServiceImpl with the given repository.
func NewRuleServiceImpl(repo domain.RuleRepository) *RuleServiceImpl {
	return &RuleServiceImpl{repo: repo}
}

// GetRules retrieves all aggregation rules for a given company ID.
func (s *RuleServiceImpl) GetRules(ctx context.Context, companyID string) ([]domain.AggregationRule, error) {
	return s.repo.GetByCompanyID(ctx, companyID)
}
