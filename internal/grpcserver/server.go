package grpcserver

import (
	"context"
	"strings"

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

// Welcome responds to the welcome gRPC call
func (h HelloServer) Welcome(ctx context.Context, request *go_example.WelcomeRequest) (*go_example.WelcomeResponse, error) {
	// ensure our context is still valid
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default: // intentionally blank
	}

	return &go_example.WelcomeResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}

// Espresso responds to the
func (h HelloServer) Espresso(ctx context.Context, request *go_example.EspressoRequest) (*go_example.EspressoResponse, error) {
	select {
	case <- ctx.Done():
		return nil, ctx.Err()
	default:
		}
	if strings.EqualFold(request.Source, "db") {

	}
	return &go_example.EspressoResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}

func (h HelloServer) Macchiato(ctx context.Context, request *go_example.MacchiatoRequest) (*go_example.MacchiatoResponse, error) {
	select {
	case <- ctx.Done():
		return nil, ctx.Err()
	default:
		}
		return &go_example.MacchiatoResponse{Response: h.greeter.Greet(ctx, request.GetName())}, nil
}
