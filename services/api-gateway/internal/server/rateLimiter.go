package server

import (
	"log/slog"
	"net"
	"net/http"
	"net/netip"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"innoveria-iot/pkg/json"
)

// rateLimitPolicy defines the policy which the rate limit middleware follows
type rateLimitPolicy struct {
	// Limit is the token refill rate used by x/time/rate.
	Limit rate.Limit
	// Burst is the maximum number of tokens that may be consumed at once.
	Burst int
	// Window is a human-readable duration used when describing policy intent.
	Window time.Duration
}

// rateLimitEntry stores limiter state per client key (for example per IP or user).
type rateLimitEntry struct {
	// limiter is the token bucket for one key.
	limiter *rate.Limiter
	// lastSeen tracks the latest request time for stale entry cleanup.
	lastSeen time.Time
}

// rateLimitStore is an in-memory map of limiter entries guarded by a mutex.
type rateLimitStore struct {
	// mu protects entries from concurrent reads/writes.
	mu sync.Mutex
	// closeOnce ensures cleanup goroutine shutdown is triggered once.
	closeOnce sync.Once
	// entries maps a rate-limit key to its limiter state.
	entries map[string]*rateLimitEntry
	// ttl is intended to control when inactive entries are evicted.
	ttl time.Duration
	// stopCh signals the cleanup goroutine to exit.
	stopCh chan struct{}
	// doneCh closes when the cleanup goroutine has stopped.
	doneCh chan struct{}
}

func newRateLimitStore(ttl time.Duration) *rateLimitStore {
	store := &rateLimitStore{
		entries: make(map[string]*rateLimitEntry),
		ttl:     ttl,
		stopCh:  make(chan struct{}),
		doneCh:  make(chan struct{}),
	}
	go store.cleanUpLoop(1 * time.Minute)
	return store
}

func (s *rateLimitStore) cleanUpLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer func() {
		ticker.Stop()
		close(s.doneCh)
	}()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			s.mu.Lock()
			for key, entry := range s.entries {
				if now.Sub(entry.lastSeen) > s.ttl {
					delete(s.entries, key)
				}
			}
			s.mu.Unlock()
		case <-s.stopCh:
			return
		}
	}
}

func (s *rateLimitStore) Close() {
	s.closeOnce.Do(func() {
		close(s.stopCh)
		<-s.doneCh
	})
}

func (s *rateLimitStore) allow(key string, policy rateLimitPolicy) bool {
	now := time.Now()

	s.mu.Lock()
	entry, ok := s.entries[key]
	if !ok {
		entry = &rateLimitEntry{
			limiter:  rate.NewLimiter(policy.Limit, policy.Burst),
			lastSeen: now,
		}
		s.entries[key] = entry
	}
	entry.lastSeen = now
	limiter := entry.limiter
	s.mu.Unlock()
	return limiter.Allow()
}

func perMinute(requests int) rate.Limit {
	if requests <= 0 {
		return rate.Inf
	}

	return rate.Every(time.Minute / time.Duration(requests))
}

func bruteForcePolicy(r *http.Request) (rateLimitPolicy, bool) {
	path := r.URL.Path

	switch {
	case r.Method == http.MethodPost && path == AUTHENTICATION_ROUTE+"/login":
		return rateLimitPolicy{Limit: perMinute(5), Burst: 5, Window: time.Minute}, true
	case r.Method == http.MethodPost && path == AUTHENTICATION_ROUTE+"/refresh":
		return rateLimitPolicy{Limit: perMinute(20), Burst: 10, Window: time.Minute}, true
	default:
		return rateLimitPolicy{}, false
	}
}

func rateLimiterMiddleware(store *rateLimitStore, next http.Handler) http.Handler {
	// denyResult holds the structure for how the rateLimiterMiddleware handles errors
	type denyResult struct {
		code       int
		message    string
		retryAfter int64
		logMsg     string
		logFields  []any
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isBruteForceProtectedPath(r) {
			next.ServeHTTP(w, r)
			return
		}

		// Only hadle policy for /refresh and /login
		policy, ok := bruteForcePolicy(r)
		if !ok {
			next.ServeHTTP(w, r)
			return
		}

		var deny *denyResult

		ip, ok := clientIP(r)
		if !ok {
			deny = &denyResult{
				code:    http.StatusBadRequest,
				message: "invalid client address",
				logMsg:  "rejecting request with invalid client address",
				logFields: []any{
					"method", r.Method,
					"path", r.URL.Path,
					"remote_addr", r.RemoteAddr,
				},
			}
		} else {
			key := r.Method + ":" + r.URL.Path + ":" + ip
			if !store.allow(key, policy) {
				deny = &denyResult{
					code:       http.StatusTooManyRequests,
					message:    "too many requests",
					retryAfter: int64(policy.Window / time.Second),
					logMsg:     "rate limit exceeded",
					logFields: []any{
						"method", r.Method,
						"path", r.URL.Path,
						"client_ip", ip,
					},
				}
			}
		}
		if deny != nil {
			if deny.retryAfter > 0 {
				w.Header().Set("Retry-After", strconv.FormatInt(deny.retryAfter, 10))
			}
			slog.Warn(deny.logMsg, deny.logFields...)
			resp := json.ErrorResponse{
				Error: json.ErrorDetail{
					Code:    deny.code,
					Message: deny.message,
				},
			}
			if err := json.Encode(w, deny.code, resp); err != nil {
				slog.Error("failed to write rate limit response", "error", err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isBruteForceProtectedPath(r *http.Request) bool {
	path := r.URL.Path

	if r.Method != http.MethodPost {
		return false
	}

	return path == AUTHENTICATION_ROUTE+"/login" || path == AUTHENTICATION_ROUTE+"/refresh"
}

func clientIP(r *http.Request) (string, bool) {
	remoteIP, ok := parseRemoteIP(r.RemoteAddr)
	if !ok {
		return "", false
	}

	// checks to see if the proxy ip has a trused header
	if isTrustedProxyIP(remoteIP) {
		if xffIP, ok := firstXForwardedFor(r.Header.Get("X-Forwarded-For")); ok {
			return xffIP.String(), true
		}
	}

	return remoteIP.String(), true
}

func parseRemoteIP(remoteAddr string) (netip.Addr, bool) {
	host, _, err := net.SplitHostPort(strings.TrimSpace(remoteAddr))
	if err != nil {
		host = strings.TrimSpace(remoteAddr)
	}

	ip, err := netip.ParseAddr(host)
	if err != nil {
		return netip.Addr{}, false
	}

	return ip.Unmap(), true
}

func firstXForwardedFor(raw string) (netip.Addr, bool) {
	parts := strings.Split(raw, ",")
	for _, part := range parts {
		candidate := strings.TrimSpace(part)
		if candidate == "" {
			continue
		}

		ip, err := netip.ParseAddr(candidate)
		if err == nil {
			return ip.Unmap(), true
		}
	}

	return netip.Addr{}, false
}

// Todo: This is now trusted by caddy or another reverse proxy
// but should be handled here for an extra layer of reinforcement
func isTrustedProxyIP(ip netip.Addr) bool {
	return ip.IsLoopback() || ip.IsPrivate()
}
