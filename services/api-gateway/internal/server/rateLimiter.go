package server

import (
	"fmt"
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
	// entries maps a rate-limit key to its limiter state.
	entries map[string]*rateLimitEntry
	// ttl is intended to control when inactive entries are evicted.
	ttl time.Duration
}

var bruteForceRateLimitStore = newRateLimitStore(10 * time.Minute)

func newRateLimitStore(ttl time.Duration) *rateLimitStore {
	store := &rateLimitStore{
		entries: make(map[string]*rateLimitEntry),
		ttl:     ttl,
	}
	go store.cleanUpLoop(1 * time.Minute)
	return store
}

func (s *rateLimitStore) cleanUpLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		now := time.Now()
		s.mu.Lock()
		for key, entry := range s.entries {
			if now.Sub(entry.lastSeen) > s.ttl {
				delete(s.entries, key)
			}
		}
		s.mu.Unlock()
	}
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

func bruteForcePolicy(r *http.Request) rateLimitPolicy {
	path := r.URL.Path

	switch {
	case r.Method == http.MethodPost && path == AUTHENTICATION_ROUTE+"/login":
		return rateLimitPolicy{Limit: perMinute(5), Burst: 5, Window: time.Minute}
	case r.Method == http.MethodPost && path == AUTHENTICATION_ROUTE+"/refresh":
		return rateLimitPolicy{Limit: perMinute(20), Burst: 10, Window: time.Minute}
	default:
		return rateLimitPolicy{Limit: perMinute(60), Burst: 20, Window: time.Minute}
	}
}

func bruteForceKey(r *http.Request) string {
	return r.Method + ":" + r.URL.Path + ":" + clientIP(r)
}

func rateLimiterMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !isBruteForceProtectedPath(r) {
			next.ServeHTTP(w, r)
			return
		}

		policy := bruteForcePolicy(r)
		key := bruteForceKey(r)

		if bruteForceRateLimitStore.allow(key, policy) {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Retry-After", strconv.FormatInt(int64(policy.Window/time.Second), 10))
		json.HandleError(w, http.StatusTooManyRequests, fmt.Errorf("rate limit exceeded"), "too many requests")
	})
}

func isBruteForceProtectedPath(r *http.Request) bool {
	path := r.URL.Path

	if r.Method != http.MethodPost {
		return false
	}

	return path == AUTHENTICATION_ROUTE+"/login" || path == AUTHENTICATION_ROUTE+"/refresh"
}

func clientIP(r *http.Request) string {
	remoteIP, ok := parseRemoteIP(r.RemoteAddr)
	if !ok {
		return "unknown"
	}

	if isTrustedProxyIP(remoteIP) {
		if xffIP, ok := firstXForwardedFor(r.Header.Get("X-Forwarded-For")); ok {
			return xffIP.String()
		}
	}

	return remoteIP.String()
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

func isTrustedProxyIP(ip netip.Addr) bool {
	return ip.IsLoopback() || ip.IsPrivate()
}
