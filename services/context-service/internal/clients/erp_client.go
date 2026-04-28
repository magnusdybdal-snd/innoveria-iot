package clients

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/clients/mappers"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/erp/dto"
	"innoveria-iot/pkg/httpclient"
)

// erpClientImpl is the HTTP implementation of domain.ERPClient.
type erpClientImpl struct {
	baseURL string
	client  *httpclient.Client
}

// NewERPClient creates a new erpClientImpl targeting the given base URL.
func NewERPClient(baseURL string) domain.ERPClient {
	return &erpClientImpl{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// GetOrders fetches the slim order list for the given company from the ERP service.
func (c *erpClientImpl) GetOrders(ctx context.Context, companyID string) ([]domain.ERPOrderSummary, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders", c.baseURL)
	headers := map[string]string{"X-Auth-Company-Id": companyID}

	resp, err := httpclient.DoRequest[[]dto.OrderSummary](
		c.client, ctx, url, http.MethodGet, nil, headers,
	)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusUnauthorized {
			return nil, domain.ErrUnauthorized
		}
		return nil, err
	}

	orders := make([]domain.ERPOrderSummary, len(resp))
	for i, o := range resp {
		orders[i] = mappers.ToERPOrderSummary(o)
	}
	return orders, nil
}

// GetProductionResources fetches all production resources for the given company from the ERP service.
func (c *erpClientImpl) GetProductionResources(ctx context.Context, companyID string) ([]domain.ERPProductionResource, error) {
	url := fmt.Sprintf("%s/api/v1/erp/production-resources", c.baseURL)
	headers := map[string]string{"X-Auth-Company-Id": companyID}

	resp, err := httpclient.DoRequest[[]dto.ERPProductionResourceResponse](
		c.client, ctx, url, http.MethodGet, nil, headers,
	)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusUnauthorized {
			return nil, domain.ErrUnauthorized
		}
		return nil, err
	}

	resources := make([]domain.ERPProductionResource, len(resp))
	for i, r := range resp {
		resources[i] = mappers.ToERPProductionResource(r)
	}
	return resources, nil
}

// GetOrderByID fetches a single order by ID from the ERP service, enriched with
// its operations and production resources.
func (c *erpClientImpl) GetOrderByID(ctx context.Context, companyID string, orderID int64) (*domain.ERPOrder, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders/%d", c.baseURL, orderID)
	headers := map[string]string{"X-Auth-Company-Id": companyID}

	resp, err := httpclient.DoRequest[dto.OrderAggregate](
		c.client, ctx, url, http.MethodGet, nil, headers,
	)
	if err != nil {
		var httpErr *httpclient.HTTPError
		if errors.As(err, &httpErr) && httpErr.StatusCode == http.StatusNotFound {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	order := mappers.ToERPOrder(resp)
	return &order, nil
}
