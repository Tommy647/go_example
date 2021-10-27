package server

import (
	"context"
	"fmt"

	"github.com/Tommy647/grpc"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ grpc.HelloWorldServiceServer = &HelloWorldServer{}

// HelloWorldServer provides the implementation of our gRPC service
type HelloWorldServer struct{}

// HelloWorld responds to the HelloWorld gRPC call
func (h HelloWorldServer) HelloWorld(ctx context.Context, request *grpc.HelloRequest) (*grpc.HelloResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	response := "World" // default value if there is no name in the response
	if request != nil && request.GetName() != "" {
		response = request.GetName()
	}
	return &grpc.HelloResponse{Response: fmt.Sprintf("Hello, %s!", response)}, nil
}

// New instance of our gRPC service
func New() *HelloWorldServer {
	return &HelloWorldServer{}
}
