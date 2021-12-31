package grpcserver

import (
	"context"

	"github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/logger"
)

// ensure our client implements the interface - this breaks compilation if it fails
// Here we guarantee that the structs on the right implements the method defined by the
// interface on the left.
// The left side has been implemented by the generated grpc code after compilation.
// The right side is implemented just some lines below in this file. If we forget to implement
// the required method to satisfy the interface on the left this var definitions will break compilation!
var _ go_example.HelloServiceServer = &HelloServer{}
var _ go_example.CoffeeServiceServer = &CoffeeServer{}

// GreetProvider something that greets
type GreetProvider interface {
	Greet(context.Context, string) string
}

// NewHS instance of our gRPC service
func NewHS(g GreetProvider) *HelloServer {
	return &HelloServer{
		greeter: g,
	}
}

// HelloServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type HelloServer struct {
	greeter GreetProvider
	// greeter interface {Greet( context.Context, string) string}
}

// Hello responds to the Hello gRPC call
func (h HelloServer) Hello(ctx context.Context, request *go_example.HelloRequest) (*go_example.HelloResponse, error) {
	logger.Info(ctx, "call to Hello") // , zap.String("name", request.Name))
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &go_example.HelloResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}

type CoffeeProvider interface {
	CoffeeGreet(context.Context, string) string
}

type CoffeeServer struct {
	coffeer CoffeeProvider
}

func NewCS(c CoffeeProvider) *CoffeeServer {
	return &CoffeeServer{
		coffeer: c,
	}
}

// Coffee responds to the Coffee gRPC call
func (c CoffeeServer) Coffee(ctx context.Context, request *go_example.CoffeeRequest) (*go_example.CoffeeResponse, error) {
	logger.Info(ctx, "call to Coffee")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally left blank
	}
	return &go_example.CoffeeResponse{Price: c.coffeer.CoffeeGreet(ctx,
		request.GetType())}, nil
}
