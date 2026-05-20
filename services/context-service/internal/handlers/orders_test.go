package handlers_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/context-service/internal/handlers"
)

type mockContextService struct {
	t                *testing.T
	getOrdersFunc    func(ctx context.Context, companyID, userID, role string) ([]domain.ERPOrderSummary, error)
	getOrderByIDFunc func(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.ERPOrder, error)
	getOrderCtxFunc  func(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.OrderContext, error)
}

func (m *mockContextService) GetOrders(ctx context.Context, companyID, userID, role string) ([]domain.ERPOrderSummary, error) {
	if m.getOrdersFunc == nil {
		m.t.Fatal("unexpected call to GetOrders")
	}
	return m.getOrdersFunc(ctx, companyID, userID, role)
}

func (m *mockContextService) GetOrderByID(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.ERPOrder, error) {
	if m.getOrderByIDFunc == nil {
		m.t.Fatal("unexpected call to GetOrderByID")
	}
	return m.getOrderByIDFunc(ctx, companyID, userID, role, orderID)
}

func (m *mockContextService) GetOrderContext(ctx context.Context, companyID, userID, role string, orderID int64) (*domain.OrderContext, error) {
	if m.getOrderCtxFunc == nil {
		m.t.Fatal("unexpected call to GetOrderContext")
	}
	return m.getOrderCtxFunc(ctx, companyID, userID, role, orderID)
}

func (m *mockContextService) GetContextData(
	_ context.Context, _, _, _ string, _ []string, _ string, _, _ time.Time, _ int,
) ([]domain.ContextData, error) {
	m.t.Fatal("unexpected call to GetContextData")
	return nil, nil
}

func (m *mockContextService) GetProductionResources(_ context.Context, _ string) ([]domain.ERPProductionResource, error) {
	m.t.Fatal("unexpected call to GetProductionResources")
	return nil, nil
}

// TestGetOrders_Success_Returns200 verifies that a successful service call returns 200 with the order list.
func TestGetOrders_Success_Returns200(t *testing.T) {
	svc := &mockContextService{
		t: t,
		getOrdersFunc: func(_ context.Context, _, _, _ string) ([]domain.ERPOrderSummary, error) {
			return []domain.ERPOrderSummary{
				{ID: 1, Name: "MO-2026-001"},
				{ID: 2, Name: "MO-2026-002"},
			}, nil
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodGet, "/api/v1/context/orders", nil))
	rec := httptest.NewRecorder()

	handlers.GetOrders(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body []struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(body) != 2 {
		t.Fatalf("expected 2 orders, got %d", len(body))
	}
	if body[0].ID != 1 || body[0].Name != "MO-2026-001" {
		t.Errorf("unexpected first order: %+v", body[0])
	}
	if body[1].ID != 2 || body[1].Name != "MO-2026-002" {
		t.Errorf("unexpected second order: %+v", body[1])
	}
}

// TestGetOrders_EmptyList_Returns200 verifies that an empty order list returns 200 with an empty array.
func TestGetOrders_EmptyList_Returns200(t *testing.T) {
	svc := &mockContextService{
		t: t,
		getOrdersFunc: func(_ context.Context, _, _, _ string) ([]domain.ERPOrderSummary, error) {
			return []domain.ERPOrderSummary{}, nil
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodGet, "/api/v1/context/orders", nil))
	rec := httptest.NewRecorder()

	handlers.GetOrders(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var body []any
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response body: %v", err)
	}
	if len(body) != 0 {
		t.Errorf("expected empty array, got %d items", len(body))
	}
}

// TestGetOrders_ServiceError_Returns500 verifies that a service error returns 500 Internal Server Error.
func TestGetOrders_ServiceError_Returns500(t *testing.T) {
	svc := &mockContextService{
		t: t,
		getOrdersFunc: func(_ context.Context, _, _, _ string) ([]domain.ERPOrderSummary, error) {
			return nil, errors.New("erp unavailable")
		},
	}

	req := withAuthHeaders(httptest.NewRequest(http.MethodGet, "/api/v1/context/orders", nil))
	rec := httptest.NewRecorder()

	handlers.GetOrders(svc).ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
