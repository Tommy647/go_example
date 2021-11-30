package interceptor

import (
	"context"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Tommy647/go_example/internal/logger"
	"github.com/Tommy647/go_example/internal/trace"
)

// traceHeader for passing trace ID between requests
const traceHeader = `trace-id`

// AttachTrace ID to each outgoing grpc request
func AttachTrace() grpc.UnaryClientInterceptor { return attachTrace }

// WithTrace gets a trace ID from an incoming request and attaches it to the context
func WithTrace() grpc.UnaryServerInterceptor { return withTrace } // @todo: complete

// withTrace ID expected in the request metadata, generates a new one if missing
func withTrace(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	resp interface{},
	err error,
) {
	logger.Info(ctx, "with trace interceptor", zap.String("method", info.FullMethod))
	traceID, err := getMetadata(ctx, traceHeader)
	if err != nil {
		traceID = uuid.NewString()
	}
	ctx = trace.WithTraceID(ctx, traceID)
	return handler(ctx, req)
}

// attachTrace ID to outgoing grpc requests
func attachTrace(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	logger.Info(ctx, "attach trace ID interceptor", zap.String("method", method))
	// traceID should be on the context
	traceID := trace.GetTraceID(ctx)
	// attach to the grpc metadata via the context (stores it in a different place to how we use it)
	ctx = metadata.AppendToOutgoingContext(ctx, traceHeader, traceID)
	return invoker(ctx, method, req, reply, cc, opts...)
}
