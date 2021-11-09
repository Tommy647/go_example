package grpcserver

import (
	"context"

	"github.com/Tommy647/go_example"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ go_example.HelloServiceServer = &HelloServer{}

// GreetProvider something that greets
type GreetProvider interface {
	Greet(context.Context, string) string
}

// HelloServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type HelloServer struct {
	greeter GreetProvider
}

// Bye responds to the Bye gRPC call
func (h HelloServer) Bye(ctx context.Context, request *go_example.ByeRequest) (*go_example.ByeResponse, error) {
	return &go_example.ByeResponse{Response: h.greeter.HelloGreet(ctx, request.GetName())}, nil
}

// New instance of our gRPC service
func New(g GreetProvider) *HelloServer {
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

	return &go_example.HelloResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}
