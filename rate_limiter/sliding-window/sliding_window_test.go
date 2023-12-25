package sliding_window

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTradeLimitMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		userId         string
		numberOfTrades int
		sleepDuration  time.Duration
		expectedStatus int
	}{
		{"UnderLimit", "USER1", 4, 0, http.StatusOK},
		{"AtLimit", "USER2", 5, 0, http.StatusOK},
		{"OverLimit", "USER3", 6, 0, http.StatusTooManyRequests},
		{"ResetAfterOneMinute", "USER4", 6, 10 * time.Second, http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new TradeWindow for each test to ensure isolation
			// If not, these trades are not isolated, then some test cases work for single, not work for together.
			req := httptest.NewRequest("GET", "/getDetails", nil)
			req.Header.Set("X-User-Id", tt.userId)
			handler := TradeLimitMiddleware(http.HandlerFunc(TradeHandler))

			for i := 1; i <= tt.numberOfTrades; i++ {
				if tt.sleepDuration > 0 {
					time.Sleep(tt.sleepDuration)
				}

				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)

				if tt.name == "OverLimit" && i > 5 {
					if status := rr.Code; status != tt.expectedStatus { // assertions
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}

				if tt.name == "ResetAfterOneMinute" {
					if status := rr.Code; status != tt.expectedStatus { // assertions
						t.Errorf("handler returned wrong status code: got %v want %v",
							status, tt.expectedStatus)
					}
				}
			}
		})
	}
}

func TestAllowTrade(t *testing.T) {
	// Initialize userActivities for testing
	userActivities = make(map[string]*TradeTracker)

	userId := "USER1"
	time1 := time.Now()

	t.Run("AllowTrade", func(t *testing.T) {
		AllowTrade(userId, time1)
		activity, exists := userActivities[userId]
		if !exists {
			t.Fatal("userActivity not found for user", userId)
		}
		if len(activity.TradeTimes) != 1 || !activity.TradeTimes[0].Equal(time1) {
			t.Error("AllowTrade did not update TradeTimes correctly")
		}
	})
}
