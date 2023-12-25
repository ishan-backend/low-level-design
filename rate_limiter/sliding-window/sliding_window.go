package sliding_window

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type UserActivityTracker struct {
	CountOfRequests int
	NextResetTime   time.Time
}

var (
	userActivities = make(map[string]*UserActivityTracker)
	mutex          = &sync.Mutex{}
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		userId := request.Header.Get("X-User-Id")

		mutex.Lock()
		currTime := time.Now()
		activity, exists := userActivities[userId]
		if !exists || currTime.After(activity.NextResetTime) {
			// this step of initialising when userActivities is nil
			userActivities[userId] = &UserActivityTracker{
				CountOfRequests: 1,
				NextResetTime:   currTime.Add(1 * time.Minute),
			}
			mutex.Unlock()
			next.ServeHTTP(writer, request)
			return
		}

		if activity.CountOfRequests >= 10 {
			http.Error(writer, "Rate limit exceeded", http.StatusTooManyRequests)
			mutex.Unlock()
			return
		}

		activity.CountOfRequests++
		mutex.Unlock()
		next.ServeHTTP(writer, request)
	})
}

func GetDetailsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Details fetched successfully")
	// exact logic flow inside this handler
}
