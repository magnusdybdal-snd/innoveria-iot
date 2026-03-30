package monitor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"

	"innoveria-iot/erp-agent-service/internal/monitor/dto"
	"innoveria-iot/pkg/httpclient"
)

type Client struct {
	host, lang, company string
	forceRelogin        bool
	username, password  string
	httpClient          *httpclient.Client
	sessionID           string

	mu            sync.Mutex
	sessiontSetAt time.Time
}

// base returns the base url for accessing monitor erp
func (c *Client) base() string {
	return fmt.Sprintf("https://%s:8001/%s/%s", c.host, c.lang, c.company)
}

// loginUrl returns the url for authenticating with monitor erp
// it will return a session id which is used throughout the api calls
func (c *Client) loginUrl() string {
	return c.base() + "/login"
}

// apiUrl returns the endpoint for
func (c *Client) apiUrl(path string) string {
	return c.base() + "/api/v1/" + path
}

// ensureSession will start a new session with monitor erp
// it will extract a session id which is used in all api calls
func (c *Client) ensureSession(ctx context.Context) error {
	c.mu.Lock()
	if c.sessionID != "" {
		c.mu.Unlock()
		return nil
	}
	c.mu.Unlock()

	// auth request body for monitor erp
	body := dto.AuthBody{
		Username:     c.username,
		Password:     c.password,
		ForceRelogin: c.forceRelogin,
	}

	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		c.loginUrl(),
		http.MethodPost,
		body,
		map[string]string{
			"Accept": "application/json",
		},
	)
	if err != nil {
		return err
	}

	// Extracting session id
	sid := resp.Header.Get("X-Monitor-SessionId")
	if sid == "" {
		return fmt.Errorf("login succeeded but missing session id")
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}

	c.mu.Lock()
	c.sessionID = sid
	c.sessiontSetAt = time.Now()
	c.mu.Unlock()

	return nil
}

func (c *Client) queryOnce(ctx context.Context, u string, out any) error {
	// Extract the session id
	c.mu.Lock()
	sid := c.sessionID
	c.mu.Unlock()

	resp, err := httpclient.DoRaw(
		c.httpClient,
		ctx,
		u,
		http.MethodGet,
		nil,
		map[string]string{
			"Accept":              "application/json",
			"X-Monitor-SessionId": sid,
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if out == nil {
		io.Copy(io.Discard, resp.Body)
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		if errors.Is(err, io.EOF) {
			return nil
		}
		return fmt.Errorf("decode monitor query response: %w", err)
	}
	return nil
}

func (c *Client) Query(ctx context.Context, path string, opts url.Values, out any) error {
	if err := c.ensureSession(ctx); err != nil {
		return err
	}

	// Create a api route with queries
	u := c.apiUrl(path)
	if len(opts) > 0 {
		u += "?" + opts.Encode()
	}

	err := c.queryOnce(ctx, u, out)
	if err == nil {
		return nil
	}

	// if the session is stale, empty the session id
	c.mu.Lock()
	c.sessionID = ""
	c.mu.Unlock()

	return nil
}
