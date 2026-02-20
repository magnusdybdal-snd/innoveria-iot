package chirpstackrest

import (
	"context"
	"net/http"

	"innoveria-iot/device-service/internal/config"

	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	baseURL    string
	token      string
	httpClient *httpclient.Client
}

func New(cfg config.Config) *Client {
	return &Client{
		baseURL:    cfg.ChirpstackURL,
		token:      cfg.ChirpstackSecret,
		httpClient: httpclient.New(),
	}
}

func (c *Client) GetAllGatewayStatus(ctx context.Context, devEUI string, limit int) (ChirpstackGatewayList, error) {
	resp, err := httpclient.DoRequest[ChirpstackGatewayList](
		c.httpClient,
		ctx,
		c.baseURL+"/api/gateways?limit=1",
		http.MethodGet,
		nil,
		map[string]string{
			"Authorization": "Bearer " + c.token,
		},
	)
	if err != nil {
		return ChirpstackGatewayList{}, err
	}

	return resp, nil
}
