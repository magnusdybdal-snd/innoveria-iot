package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

// New construct a resuable HTTP Client with global request timeout
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

// HTTPError handles HTTP Error handling
type HTTPError struct {
	StatusCode int
	Status     string
	URL        string
	Method     string
	Body       []byte
}

// Error handling when server responds with non-2xx status code
func (e *HTTPError) Error() string {
	return fmt.Sprintf("%s %s: %s", e.Method, e.URL, e.Status)
}

// DoRequest sends an HTTP request and decodes a JSON response into T.
//
// Behavior:
//   - If body != nil, it is JSON-marshaled and sent as the request body.
//   - headers are applied as-is.
//   - Any non-2xx response returns an *HTTPError containing StatusCode/Status and
//     a truncated response body.
//   - On success, the response body is JSON-decoded into T. If the response body
//     is empty (common for 204), it returns the zero value of T and nil error.
func DoRequest[T any](
	client *Client,
	ctx context.Context,
	url string,
	method string,
	body any,
	headers map[string]string,
) (T, error) {
	var zero T

	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return zero, fmt.Errorf("marshall request body %w", err)
		}
		reqBody = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
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
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		const maxErrBody = 8 << 10 // 8KB
		b, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrBody))
		return zero, &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			URL:        url,
			Method:     method,
			Body:       b,
		}
	}

	var data T
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		if errors.Is(err, io.EOF) {
			return zero, nil
		}
		return zero, fmt.Errorf("decode response: %w", err)
	}

	return data, nil
}

// DoRaw returns the request to handle responses which only returns an status code
func DoRaw(
	client *Client,
	ctx context.Context,
	url string,
	method string,
	body any,
	headers map[string]string,
) (*http.Response, error) {
	var reqBody io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshall request body %w", err)
		}
		reqBody = bytes.NewReader(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := client.http.Do(req) // Defer body where the function is used
	if err != nil {
		return nil, fmt.Errorf("error doing the request: %w", err)
	}

	// error handling for status codes
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		const maxErrBody = 8 << 10 // 8KB
		b, _ := io.ReadAll(io.LimitReader(resp.Body, maxErrBody))
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			URL:        url,
			Method:     method,
			Body:       b,
		}
	}

	return resp, nil
}
