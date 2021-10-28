package grpcserver

import (
	"context"

	grpc "github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/greeter"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ grpc.HelloWorldServiceServer = &HelloWorldServer{}

// HelloWorldServer provides the implementation of our gRPC service
type HelloWorldServer struct{}

// New instance of our gRPC service
func New() *HelloWorldServer {
	return &HelloWorldServer{}
}

// HelloWorld responds to the HelloWorld gRPC call
func (h HelloWorldServer) Hello(ctx context.Context, request *grpc.HelloRequest) (*grpc.HelloResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &grpc.HelloResponse{Response: greeter.HelloGreet(request.GetName())}, nil
}
