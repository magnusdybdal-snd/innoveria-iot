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

// GetOrders fetches the list of orders for the given company from the ERP service.
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

// GetOrderDetail fetches the full detail of a single order, including its reportings.
func (c *ERPClient) GetOrderDetail(ctx context.Context, orderID string) (*domain.ERPOrderDetail, error) {
	url := fmt.Sprintf("%s/api/v1/erp/orders/%s", c.baseURL, orderID)

	resp, err := httpclient.DoRequest[dto.ERPOrderDetailResponse](
		c.client, ctx, url, http.MethodGet, nil, nil,
	)
	if err != nil {
		return nil, err
	}

	detail := mappers.ToERPOrderDetail(resp)
	return &detail, nil
}
