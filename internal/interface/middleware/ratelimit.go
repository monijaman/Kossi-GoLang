package middleware

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// tokenBucket is a simple per-visitor token bucket used to throttle request rate.
type tokenBucket struct {
	mu         sync.Mutex
	tokens     float64
	capacity   float64
	refillRate float64 // tokens per second
	lastRefill time.Time
}

func newTokenBucket(capacity, refillRate float64) *tokenBucket {
	return &tokenBucket{
		tokens:     capacity,
		capacity:   capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}
}

// allow reports whether a request may proceed, consuming one token if so.
func (b *tokenBucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastRefill).Seconds()
	b.lastRefill = now

	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}

	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

type visitorEntry struct {
	bucket   *tokenBucket
	lastSeen time.Time
}

// rateLimiter tracks a token bucket per client IP, evicting idle entries so
// memory does not grow unbounded when a scraper cycles through many IPs.
type rateLimiter struct {
	mu         sync.Mutex
	visitors   map[string]*visitorEntry
	capacity   float64
	refillRate float64
}

func newRateLimiter(requestsPerMinute, burst int) *rateLimiter {
	rl := &rateLimiter{
		visitors:   make(map[string]*visitorEntry),
		capacity:   float64(burst),
		refillRate: float64(requestsPerMinute) / 60.0,
	}
	go rl.cleanupLoop()
	return rl
}

func (rl *rateLimiter) cleanupLoop() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		cutoff := time.Now().Add(-10 * time.Minute)
		rl.mu.Lock()
		for ip, v := range rl.visitors {
			if v.lastSeen.Before(cutoff) {
				delete(rl.visitors, ip)
			}
		}
		rl.mu.Unlock()
	}
}

func (rl *rateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	v, exists := rl.visitors[ip]
	if !exists {
		v = &visitorEntry{bucket: newTokenBucket(rl.capacity, rl.refillRate)}
		rl.visitors[ip] = v
	}
	v.lastSeen = time.Now()
	rl.mu.Unlock()

	return v.bucket.allow()
}

// envInt reads an integer environment variable, falling back to def if unset or invalid.
func envInt(key string, def int) int {
	if raw := os.Getenv(key); raw != "" {
		if v, err := strconv.Atoi(raw); err == nil && v > 0 {
			return v
		}
	}
	return def
}

// clientIP extracts the originating client IP, preferring the first
// X-Forwarded-For entry since the service typically sits behind a
// platform proxy (Railway) or the Next.js proxy route, both of which
// forward this header from the original browser/client request.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		if ip := strings.TrimSpace(parts[0]); ip != "" {
			return ip
		}
	}
	if xrip := r.Header.Get("X-Real-IP"); xrip != "" {
		return strings.TrimSpace(xrip)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// Default limits: generous enough for normal browsing/app usage (product
// listing + images loading in parallel) while capping sustained bulk
// scraping. Override with RATE_LIMIT_RPM / RATE_LIMIT_BURST env vars.
var defaultLimiter = newRateLimiter(
	envInt("RATE_LIMIT_RPM", 300),
	envInt("RATE_LIMIT_BURST", 60),
)

// exemptPaths bypasses rate limiting for infrastructure endpoints and
// already-uploaded static assets, which are not the scraping target and
// whose throttling could break platform health checks.
func exempt(path string) bool {
	return path == "/health" || strings.HasPrefix(path, "/uploads/")
}

// RateLimitMiddleware throttles requests per client IP using a token
// bucket, returning 429 once the bucket is exhausted. This is a
// best-effort, single-instance defense against scraping/abuse; IP
// rotation still bypasses it, so pair it with edge-level bot protection
// (e.g. Cloudflare) for stronger coverage.
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if exempt(r.URL.Path) {
			next.ServeHTTP(w, r)
			return
		}

		ip := clientIP(r)
		if !defaultLimiter.allow(ip) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Retry-After", "5")
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "rate limit exceeded, please slow down",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
