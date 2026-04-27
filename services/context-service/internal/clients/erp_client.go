package clients

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/clients/mappers"
	"innoveria-iot/context-service/internal/domain"
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
//
// TODO: the route /api/v1/erp/orders does not exist on the erp-service yet.
// Update the URL and query parameters once the GET endpoint is implemented.
func (c *erpClientImpl) GetOrders(ctx context.Context, companyID string) ([]domain.ERPOrderSummary, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders", c.baseURL)
	headers := map[string]string{"X-Auth-Company-Id": companyID}

	resp, err := httpclient.DoRequest[[]dto.ERPOrderSummaryResponse](
		c.client, ctx, url, http.MethodGet, nil, headers,
	)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.ERPOrderSummary, len(resp))
	for i, o := range resp {
		orders[i] = mappers.ToERPOrderSummary(o)
	}
	return orders, nil
}

// GetOrderByID fetches a single order by ID from the ERP service, enriched with
// its operations and production resources.
//
// TODO: the route /api/v1/erp/orders/{id} does not exist on the erp-service yet.
// Update the URL once the GET endpoint is implemented.
func (c *erpClientImpl) GetOrderByID(ctx context.Context, companyID string, orderID int64) (*domain.ERPOrder, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders/%d", c.baseURL, orderID)
	headers := map[string]string{"X-Auth-Company-Id": companyID}

	resp, err := httpclient.DoRequest[dto.ERPOrderResponse](
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
