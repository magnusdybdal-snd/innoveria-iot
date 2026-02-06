package proxy

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
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
	rp := httputil.NewSingleHostReverseProxy(base)

	// Transport for configuring own tranpsport. So we dont use the default transport config
	rp.Transport = DefaultReverseProxyTransport()

	// Flush repsonse periodically instead of buffering
	rp.FlushInterval = 100 * time.Millisecond
	return rp
}
