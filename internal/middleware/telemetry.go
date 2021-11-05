package middleware

import (
	"log"
	"net/http"
	"time"
)

// WithBasicTelemetry to track request timing
func WithBasicTelemetry(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("tracking request timings")
		next.ServeHTTP(w, r)
		log.Printf("call to %s %s took %s", r.Method, r.URL.Path, time.Since(start))
	})
}
