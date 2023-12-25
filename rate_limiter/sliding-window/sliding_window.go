package sliding_window

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

/*
	Assume a stock trading platform limits the number of trades a user can execute within any rolling 1-minute window
	to maintain market integrity and prevent rapid trading that could lead to market manipulation.
*/

const GLOBAL_MAX_TRADES = 5

type TradeTracker struct {
	TradeTimes []time.Time
	Mutex      sync.Mutex
}

var (
	userActivities = make(map[string]*TradeTracker)
)

func TradeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trade executed successfully")
	// Do exact logic after trade executed.
}

func TradeLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("X-User-Id")
		currTime := time.Now()

		if !CheckIfTradeAllowed(userId, currTime) {
			http.Error(w, "Trade limit exceeded", http.StatusTooManyRequests)
			return
		}

		AllowTrade(userId, currTime)
		next.ServeHTTP(w, r)
	})
}

func CheckIfTradeAllowed(userId string, currTime time.Time) bool {
	_, exists := userActivities[userId]
	if !exists {
		userActivities[userId] = &TradeTracker{
			TradeTimes: nil,
			Mutex:      sync.Mutex{},
		}
	}

	activity := userActivities[userId]
	activity.Mutex.Lock()
	defer activity.Mutex.Unlock()

	tenSecondsAgo := currTime.Add(-10 * time.Second)

	// Check if trades
	var validTradeTimes []time.Time // nil slice declaration
	for _, tradeTime := range activity.TradeTimes {
		if tradeTime.After(tenSecondsAgo) {
			validTradeTimes = append(validTradeTimes, tradeTime)
		}
	}
	activity.TradeTimes = validTradeTimes

	return len(activity.TradeTimes) < GLOBAL_MAX_TRADES
}

func AllowTrade(userId string, currTime time.Time) {
	_, exists := userActivities[userId]
	if !exists {
		userActivities[userId] = &TradeTracker{
			TradeTimes: nil,
			Mutex:      sync.Mutex{},
		}
	}

	activity := userActivities[userId]
	activity.Mutex.Lock()
	defer activity.Mutex.Unlock()

	activity.TradeTimes = append(activity.TradeTimes, currTime)
}
