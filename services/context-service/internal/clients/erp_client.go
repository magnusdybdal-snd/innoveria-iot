package clients

import (
	"context"
	"fmt"
	"net/http"

	"innoveria-iot/context-service/internal/clients/dto"
	"innoveria-iot/context-service/internal/clients/mappers"
	"innoveria-iot/context-service/internal/domain"
	"innoveria-iot/pkg/httpclient"
)

// ERPClient is an HTTP client for the ERP service.
type ERPClient struct {
	baseURL string
	client  *httpclient.Client
}

// NewERPClient creates a new ERPClient targeting the given base URL.
func NewERPClient(baseURL string) *ERPClient {
	return &ERPClient{
		baseURL: baseURL,
		client:  httpclient.New(),
	}
}

// GetOrders fetches all orders for the given company from the ERP service,
// each enriched with their operations and production resources.
//
// TODO: the route /api/v1/erp/orders does not exist on the erp-service yet.
// Update the URL and query parameters once the GET endpoint is implemented.
func (c *ERPClient) GetOrders(ctx context.Context, companyID string) ([]domain.ERPOrder, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders?company_id=%s", c.baseURL, companyID)

	resp, err := httpclient.DoRequest[[]dto.ERPOrderResponse](
		c.client, ctx, url, http.MethodGet, nil, nil,
	)
	if err != nil {
		return nil, err
	}

	orders := make([]domain.ERPOrder, len(resp))
	for i, o := range resp {
		orders[i] = mappers.ToERPOrder(o)
	}
	return orders, nil
}

// GetOrderByID fetches a single order by ID from the ERP service, enriched with
// its operations and production resources.
//
// TODO: the route /api/v1/erp/orders/{id} does not exist on the erp-service yet.
// Update the URL once the GET endpoint is implemented.
func (c *ERPClient) GetOrderByID(ctx context.Context, companyID string, orderID int64) (*domain.ERPOrder, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders/%d?company_id=%s", c.baseURL, orderID, companyID)

	resp, err := httpclient.DoRequest[dto.ERPOrderResponse](
		c.client, ctx, url, http.MethodGet, nil, nil,
	)
	if err != nil {
		return nil, err
	}

	order := mappers.ToERPOrder(resp)
	return &order, nil
}
