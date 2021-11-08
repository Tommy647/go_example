package grpcserver

import (
	"context"

	"github.com/Tommy647/go_example"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ go_example.HelloServiceServer = &HelloServer{}

// Greeter something that greets
type Greeter interface {
	HelloGreet(context.Context, string) string
}

// HelloServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type HelloServer struct {
	greeter Greeter
}

// New instance of our gRPC service
func New(g Greeter) *HelloServer {
	return &HelloServer{
		greeter: g,
	}
}

// Hello responds to the Hello gRPC call
func (h HelloServer) Hello(ctx context.Context, request *go_example.HelloRequest) (*go_example.HelloResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &go_example.HelloResponse{Response: h.greeter.HelloGreet(ctx, request.GetName())}, nil
}
