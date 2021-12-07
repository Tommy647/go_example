package trace

import "context"

// ctxTrace to namespace our trace ID on the context
const ctxTrace ctxTraceID = `trace-id`

type ctxTraceID string

// WithTraceID adds the trace ID to the context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, ctxTrace, traceID)
}

// GetTraceID from the context
func GetTraceID(ctx context.Context) string {
	if traceID, ok := ctx.Value(ctxTrace).(string); ok {
		return traceID
	}
	return ""
}
