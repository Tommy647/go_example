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

// CustomGreetProvider something that greets
type CustomGreetProvider interface {
	Greet(context.Context, string, string) string
}

// CustomGreetServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type CustomGreetServer struct {
	greeter CustomGreetProvider
}

// NewGreeter instance of our gRPC service
func NewGreeter(g CustomGreetProvider) *CustomGreetServer {
	return &CustomGreetServer{
		greeter: g,
	}
}

// CustomGreeter responds to the Hello gRPC call
func (h CustomGreetServer) CustomGreeter(ctx context.Context, request *go_example.CustomGreeterRequest) (*go_example.CustomGreeterResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &go_example.CustomGreeterResponse{Response: h.greeter.Greet(ctx, request.Greeting, request.Name)}, nil
}
