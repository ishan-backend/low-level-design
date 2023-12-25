package token_bucket

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
	Two use-cases:

	A CDN needs to limit the bandwidth used by individual users to prevent any single user from consuming excessive resources,
    which could lead to network congestion and degraded service for others.

	An API service needs to limit the number of requests a user can make per second to prevent overuse or abuse,
    ensuring fair access and stable performance for all users.
*/

// code is written for one user

const (
	capacity   = 5
	refillRate = 1 // number of tokens added per second
)

type TokenBucket struct {
	Capacity          int
	CurrentLeftTokens int
	RefillRate        int
	LastRefillTime    time.Time
	Mutex             sync.Mutex
}

func NewTokenBucket() *TokenBucket {
	return &TokenBucket{
		Capacity:          capacity,
		CurrentLeftTokens: capacity,
		RefillRate:        refillRate,
		LastRefillTime:    time.Now(),
	}
}

func (tb *TokenBucket) Refill() {
	tb.Mutex.Lock()
	defer tb.Mutex.Unlock()

	currTime := time.Now()
	elapsedTimeInSeconds := currTime.Sub(tb.LastRefillTime).Seconds()
	tokenCountToRefill := int(elapsedTimeInSeconds) * tb.RefillRate

	tb.CurrentLeftTokens += tokenCountToRefill
	if tb.CurrentLeftTokens > tb.Capacity {
		tb.CurrentLeftTokens = tb.Capacity
	}
	tb.LastRefillTime = currTime
}

func (tb *TokenBucket) AllowRequest() bool {
	tb.Refill()

	if tb.CurrentLeftTokens > 0 {
		tb.CurrentLeftTokens--
		return true
	}

	return false
}

func RateLimitMiddleware(tb *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !tb.AllowRequest() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func APIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API request successful")
	// Do some real logic
}
