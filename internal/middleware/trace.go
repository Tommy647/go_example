package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/Tommy647/go_example/internal/trace"
)

// traceHeader from envoy docs
const traceHeader = `x-request-d`

// WithTrace gets a trace ID from the incoming header, or generates
// a new ID, and adds this to the context for following system calls
func WithTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := r.Header.Get(traceHeader)
		if traceID == "" { // if it is blank, create a new uuid
			traceID = uuid.New().String()
		}
		// attach to the ctx to make it available everywhere else
		ctx := trace.WithTraceID(r.Context(), traceID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
