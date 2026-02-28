package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

func New() *Client {
	return &Client{
		http: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:       10,
				IdleConnTimeout:    30 * time.Second,
				DisableCompression: true,
			},
			Timeout: 10 * time.Second,
		},
	}
}

func DoRequest[T any](
	client *Client,
	ctx context.Context,
	url string,
	method string,
	body any,
	headers map[string]string,
) (T, error) {
	var zero T

	// TODO: Handle body for PUT, POST, PATCH requests
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return zero, fmt.Errorf("request error: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return zero, fmt.Errorf("error doing the request: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck // TODO(vinjar): handle or wrap Body.Close error properly

	// error handling for status codes
	if resp.StatusCode >= http.StatusBadRequest {
		return zero, err
	}
	var data T
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return zero, err
	}

	return data, nil
}
