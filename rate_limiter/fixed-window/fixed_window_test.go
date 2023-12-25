package fixed_window

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

/*
		go test ./...
	    go test -v -coverpkg=./... -coverprofile=profile.cov ./...     --> creates profile.cov

		View coverage(profile.cov):
		go tool cover -html=profile.cov
*/
func TestRateLimitMiddleware(t *testing.T) {
	tests := []struct {
		name                  string
		userId                string
		currentRequestCount   int
		sendAdditionalRequest bool
		expectedStatus        int
		resetAfter            time.Duration
	}{
		{"UnderLimit", "USER1", 5, true, http.StatusOK, 0},
		{"OverLimit", "USER2", 10, true, 429, 0},
		{"TryAfterReset", "USER3", 10, true, http.StatusOK, time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// mock request
			req := httptest.NewRequest("GET", "/getDetails", nil)
			req.Header.Set("X-User-Id", tt.userId)
			handler := RateLimitMiddleware(http.HandlerFunc(GetDetailsHandler))

			for i := 0; i < tt.currentRequestCount; i++ {
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)
			}

			// wait for reset interval if needed
			if tt.resetAfter > 0 {
				time.Sleep(tt.resetAfter)
			}

			// Send an additional request if required
			if tt.sendAdditionalRequest {
				rr := httptest.NewRecorder()
				handler.ServeHTTP(rr, req)

				// Assert for the additional request
				if status := rr.Code; status != tt.expectedStatus {
					t.Errorf("handler returned wrong status code: got %v want %v",
						status, tt.expectedStatus)
				}
			}
		})
	}
}
