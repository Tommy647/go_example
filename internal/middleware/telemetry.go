package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/Tommy647/go_example/internal/logger"
)

// WithBasicTelemetry to track request timing
func WithBasicTelemetry(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info(r.Context(), "tracking request timings")

		next.ServeHTTP(w, r) // Calls next handler (Auth -> HandleHello)

		logger.Info(
			r.Context(),
			"call completed",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Duration("duration", time.Since(start)),
		)
	})
}
