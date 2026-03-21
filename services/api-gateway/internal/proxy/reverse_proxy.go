// Package proxy TODO(@vinjar): add proper documentation.
package proxy

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"innoveria-iot/pkg/json"
)

// DefaultReverseProxyTransport returns a shared HTTP transport
// configured for gateway / reverse-proxy usage.
//
// Why a custom transport?
//   - The default transport has very permissive timeouts
//   - Gateways must fail fast when upstreams are slow or unhealthy
//   - Connection reuse and pooling are critical for performance
//
// This transport is safe to share across multiple reverse proxies
// and upstream services.
func DefaultReverseProxyTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          256,
		MaxIdleConnsPerHost:   64,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ResponseHeaderTimeout: 30 * time.Second,
	}
}

// NewReverseProxy constructs a reverse proxy targeting a single upstream host.
//
// It wraps httputil.NewSingleHostReverseProxy and applies gateway-level
// defaults such as:
//   - shared transport
//   - streaming-friendly flush behavior
func NewReverseProxy(base *url.URL) *httputil.ReverseProxy {
	// Standard reverse proxy to target host
	rp := &httputil.ReverseProxy{
		// customized director to remove untrusted IP forwarder
		Rewrite: func(pr *httputil.ProxyRequest) {
			// Set upstream target URL
			pr.SetURL(base)

			// Remove any client-supplied forwarding headers
			pr.Out.Header.Del("X-Forwarded-For")
			pr.Out.Header.Del("X-Forwarded-Host")
			pr.Out.Header.Del("X-Forwarded-Proto")
			pr.Out.Header.Del("X-Real-IP")
			pr.Out.Header.Del("Forwarded")

			// Set trusted forwarding headers from gateway request
			pr.SetXForwarded()

		},

		// Transport for configuring own tranpsport. So we dont use the default transport config
		Transport: DefaultReverseProxyTransport(),

		// Flush repsonse periodically instead of buffering
		FlushInterval: 100 * time.Millisecond,

		// Error handling for to big requests and upstream errors
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err error) {
			var mbe *http.MaxBytesError
			if errors.As(err, &mbe) {
				json.HandleError(w, http.StatusRequestEntityTooLarge, err, "request to large")
			}

			json.HandleError(w, http.StatusBadGateway, err, "bad gateway")
		},
	}

	return rp
}
