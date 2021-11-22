// Package main a client to make requests against our servers
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/Tommy647/go_example/internal/grpcclient"
	"github.com/Tommy647/go_example/internal/tls"
)

const (
	// defaultRunTimeout for context
	defaultRunTimeout = 10 * time.Second
	envPort           = `PORT`         // envPort to make requests to
	address           = `127.0.0.1:%s` // address we expect gRPC services
)

func main() {
	log.Println("client starting") // prove the client is up
	// create a new context that expires in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), defaultRunTimeout)
	// close the context when we leave this function
	defer cancel()

	// convert our jwt token to a gRPC compatible format
	jwt, err := oauth.NewJWTAccessFromKey([]byte("I'm a JWT!")) // @todo: this
	if err != nil {
		panic("token error " + err.Error())
	}
	_ = jwt
	tlsConf, err := tls.LoadCertificates()
	if err != nil {
		panic(err.Error())
	}

	opts := []grpc.DialOption{
		//grpc.WithPerRPCCredentials(jwt),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)),
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
