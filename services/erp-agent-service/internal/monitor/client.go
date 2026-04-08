// Package monitor provides a minimal client for Monitor ERP authentication and queries.
package monitor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"innoveria-iot/erp-agent-service/internal/config"
	"innoveria-iot/erp-agent-service/internal/monitor/dto"
	"innoveria-iot/pkg/httpclient"
)

type MonitorERPEndpoint string

const (
	Base            MonitorERPEndpoint = "Manufacturing/"
	Orders          MonitorERPEndpoint = Base + "ManufacturingOrders"
	OrderReportings MonitorERPEndpoint = Base + "ManufacturingOrderOperationReportings"
	Workcenters     MonitorERPEndpoint = Base + "WorkCenters"
)

// Client manages Monitor ERP session lifecycle and authenticated API calls.
type Client struct {
	host, port, lang, company string
	forceRelogin              bool
	username, password        string
	httpClient                *httpclient.Client
	sessionID                 string

	// Used at runtime
	mu           sync.Mutex
	sessionSetAt time.Time
}

// New is the constructor for the monitor erp client
func New(cfg config.Config) *Client {
	return &Client{
		host:         cfg.MonitorERPHost,
		port:         cfg.MonitorERPPort,
		lang:         "en",
		company:      strconv.Itoa(cfg.MonitorERPCompanyNumber),
		forceRelogin: cfg.MonitorERPForceRelogin,
		username:     cfg.MonitorERPUsername,
		password:     cfg.MonitorERPPassword,
		httpClient:   httpclient.New(),
		sessionID:    "",
	}
}

// base returns the base url for accessing monitor erp
func (c *Client) base() string {
	return fmt.Sprintf("https://%s:%s/%s/%s", c.host, c.port, c.lang, c.company)
}

// loginUrl returns the url for authenticating with monitor erp
// it will return a session id which is used throughout the api calls
func (c *Client) loginUrl() string {
	return c.base() + "/login"
}

// apiUrl returns the target endpoint for monitor erp
func (c *Client) apiUrl(path MonitorERPEndpoint) string {
	return c.base() + "/api/v1/" + string(path)
}

// ensureSession will start a new session with monitor erp
// it will extract a session id which is used in all api calls
// this is a synchronized access to sessionID
func (c *Client) ensureSession(ctx context.Context) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.sessionID != "" {
		return nil
	}

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
			"Accept":        "application/json",
			"Cache-Control": "no-cache",
		},
	)
	if err != nil {
		return fmt.Errorf("monitor login failed: %w", err)
	}

	// catching response body error
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close monitor response body: %w", cerr)
		}
	}()

	// Extracting session id
	sid := resp.Header.Get("X-Monitor-SessionId")
	_, _ = io.Copy(io.Discard, resp.Body) // Drain body so conn can be reused
	if sid == "" {
		slog.Warn(
			"monitor login response missing session id; endpoint may be misconfigured",
			"url", c.loginUrl(),
			"status", resp.StatusCode,
		)
		return fmt.Errorf("monitor login succeeded but missing session id")
	}

	c.sessionID = sid
	c.sessionSetAt = time.Now()

	return nil
}

// queryOnce performs a single authenticated GET request to Monitor.
//
// It does not handle session refresh/retry logic; callers are expected to
// decide if and when a failed request should be retried.
func (c *Client) queryOnce(ctx context.Context, u, sid string, out any) (err error) {
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
	// catching response body error
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && err == nil {
			err = fmt.Errorf("close monitor response body: %w", cerr)
		}
	}()

	if out == nil {
		if _, err := io.Copy(io.Discard, resp.Body); err != nil {
			return fmt.Errorf("discard monitor response body: %w", err)
		}
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

// Query executes a Monitor query against /api/v1/{path}.
//
// Flow:
//  1. Ensures a session exists.
//  2. Executes one request.
//  3. On 401/403, clears the cached session, re-authenticates and retries once.
func (c *Client) Query(ctx context.Context, path MonitorERPEndpoint, opts url.Values, out any) error {
	if err := c.ensureSession(ctx); err != nil {
		return err
	}

	// Create a api route with queries
	u := c.apiUrl(path)
	if len(opts) > 0 {
		u += "?" + opts.Encode()
	}

	c.mu.Lock()
	sid := c.sessionID
	c.mu.Unlock()

	err := c.queryOnce(ctx, u, sid, out)
	if err == nil {
		return nil
	}

	var httpErr *httpclient.HTTPError
	if !errors.As(err, &httpErr) {
		return err
	}

	if httpErr.StatusCode == http.StatusNotFound {
		slog.Warn(
			"monitor query returned 404; route may be misconfigured",
			"path", path,
			"url", u,
		)
		return err
	}

	if httpErr.StatusCode != http.StatusUnauthorized && httpErr.StatusCode != http.StatusForbidden {
		return err
	}

	// if the session is stale, empty the session id
	// and try new query
	c.mu.Lock()
	if c.sessionID == sid {
		c.sessionID = ""
	}
	c.mu.Unlock()

	if err := c.ensureSession(ctx); err != nil {
		return fmt.Errorf("monitor relogin failed: %w", err)
	}

	c.mu.Lock()
	sid = c.sessionID
	c.mu.Unlock()

	return c.queryOnce(ctx, u, sid, out)
}
