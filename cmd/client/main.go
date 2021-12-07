// Package main a client to make requests against our servers
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/Tommy647/go_example/internal/grpcclient"
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
		grpc.WithUnaryInterceptor(interceptor.AttachAuth()),
		grpc.WithUnaryInterceptor(interceptor.AttachTrace()),
	}

	port := os.Getenv(envPort)
	if port == "" {
		panic("need the PORT env var")
	}

	// create a new instance of our client application
	c, err := grpcclient.New(
		grpcclient.WithHost(fmt.Sprintf(address, port)),
		grpcclient.WithDialOptions(opts...),
	)
	if err != nil {
		panic(err.Error())
	}

	defer func() { _ = c.Close() }()
	// start our client running with no input
	c.Run(ctx)
	// reuse the client and add some names
	c.Run(ctx, "Tom", "Orson", "Kurt")
}
