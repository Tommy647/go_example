package grpcserver

import (
	"context"
	"strings"

	"github.com/Tommy647/go_example/internal/greeter"

	"github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/logger"
)

// ensure our client implements the interface - this breaks compilation if it fails
var _ go_example.HelloServiceServer = &HelloServer{}
var _ go_example.CoffeeServiceServer = &CoffeeServer{}

// GreetProvider something that greets
type GreetProvider interface {
	Greet(context.Context, string) string
}

type CoffeeProvider interface {
	CoffeeGreet(context.Context, string) string
}

// HelloServer provides the implementation of our gRPC service
// has to meet the go_example.HelloServiceServer interface
type HelloServer struct {
	greeter GreetProvider
}

type CoffeeServer struct {
	coffeer CoffeeProvider
}

// NewHS instance of our gRPC service
func NewHS(g GreetProvider) *HelloServer {
	return &HelloServer{
		greeter: g,
	}
}

func NewCS(c CoffeeProvider) *CoffeeServer {
	return &CoffeeServer{
		coffeer: c,
	}
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

func (c CoffeeServer) Coffee(ctx context.Context, request *go_example.CoffeeRequest) (*go_example.CoffeeResponse, error) {
	logger.Info(ctx, "call to Coffee")
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally left blank
	}
	// If we receive "db" in the request we check for espresso or macchiato in the Type
	// if found we retrieve the price from the db, otherwise we provide the basic CoffeeGreeter
	if strings.EqualFold(request.Source, "db") &&
		(strings.EqualFold(request.Type, "espresso") || strings.EqualFold(request.Type, "macchiato")) {
		return &go_example.CoffeeResponse{Price: c.coffeer.CoffeeGreet(ctx, strings.Title(strings.ToLower(request.Type)))}, nil
	}
	return &go_example.CoffeeResponse{Price: greeter.New().CoffeeGreet(ctx, request.Type)}, nil
}
