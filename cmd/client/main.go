// Package main a client to make requests against our servers
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"

	"github.com/Tommy647/go_example"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/Tommy647/go_example/internal/interceptor"
	"github.com/Tommy647/go_example/internal/logger"
	"github.com/Tommy647/go_example/internal/tls"
	"github.com/Tommy647/go_example/internal/trace"
)

const (
	// defaultRunTimeout for context
	defaultRunTimeout = 10 * time.Second
	envPort           = `PORT`         // envPort to make requests to
	address           = `127.0.0.1:%s` // address we expect gRPC services
)

func main() {
	// create a new context that expires in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), defaultRunTimeout)
	// close the context when we leave this function
	defer cancel()
	ctx = trace.WithTraceID(ctx, `system`)
	if err := logger.New(`go_example_http`); err != nil {
		panic(err.Error())
	}
	logger.Info(ctx, "client starting") // prove the client is up

	tlsConf, err := tls.LoadCertificates()
	if err != nil {
		panic(err.Error())
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)),
		grpc.WithChainUnaryInterceptor(
			interceptor.AttachAuth(),
			interceptor.AttachTrace(),
		),
	}

	port := os.Getenv(envPort)
	if port == "" {
		panic("need the PORT env var")
	}

	// implement a grpc client

	conn, err := grpc.Dial(fmt.Sprintf(address, port), opts...)
	if err != nil {
		panic(err)
	}

	helloClient := go_example.NewHelloServiceClient(conn)

	hello, err := helloClient.Hello(ctx, &go_example.HelloRequest{})
	if err != nil {
		logger.Error(ctx, "helloservice hello", zap.Error(err))
		return
	}

	fmt.Println(hello.Response)

	coffeeClient := go_example.NewCoffeeServiceClient(conn)
	coffee, err := coffeeClient.Coffee(ctx, &go_example.CoffeeRequest{
		Type:   "espresso",
		Source: "DB",
	})

	fmt.Println(coffee)
}
