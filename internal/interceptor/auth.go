package interceptor

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/Tommy647/go_example/internal/jwt"
	"github.com/Tommy647/go_example/internal/logger"
)

// authorizationHeader for passing JWT token between requests
const authorizationHeader = `authorization`

// WithAuth check for the presence of a JWY in incoming requests
func WithAuth() grpc.UnaryServerInterceptor { return withAuth }

// AttachAuth interceptor to add our jwt token to all grpc requests
func AttachAuth() grpc.UnaryClientInterceptor { return attachAuth }

// withAuth check for the presence of a JWY in incoming requests and add it to the context
func withAuth(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (
	resp interface{},
	err error,
) {
	logger.Info(ctx, "with auth interceptor", zap.String("method", info.FullMethod))
	token, err := getMetadata(ctx, authorizationHeader)
	if err != nil {
		return nil, errors.New("missing authentication token")
	}

	claims, err := jwt.GetClaims(token)
	if err != nil {
		return nil, errors.New("invalid authentication token")
	}

	ctx = jwt.WithUser(ctx, claims)
	return handler(ctx, req)
}

// attachAuth jwt token to outgoing grpc requests
func attachAuth(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	logger.Info(ctx, "attach auth interceptor", zap.String("method", method))
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", os.Getenv("TOKEN"))
	return invoker(ctx, method, req, reply, cc, opts...)
}
