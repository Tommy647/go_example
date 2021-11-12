// Package main a client to make requests against our servers
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/oauth"

	"github.com/Tommy647/go_example/internal/grpcclient"
)

// defaultRunTimeout for context
const defaultRunTimeout = 10 * time.Second

func main() {
	log.Println("client starting") // prove the client is up
	// create a new context that expires in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), defaultRunTimeout)

	// close the context when we leave this function
	defer cancel()

	// convert our jwt token to a gRPC compatible format
	jwt, err := oauth.NewJWTAccessFromKey([]byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJyb2xlcyI6WyJ1c2VyIiwiaGVsbG8iXSwiaXNzIjoidGVzdCIsInN1YiI6InNvbWVib2R5IiwiYXVkIjpbInNvbWVib2R5X2Vsc2UiXSwiZXhwIjoxNjM2NzMxNzAzLCJuYmYiOjE2MzY3MjgxMDMsImlhdCI6MTYzNjcyODEwMywianRpIjoiMSJ9.R9FQidUi3WJ2KvPKB00UVF7FyKi4lPFrvYyHipQ4em8"))
	if err != nil {
		panic("token error " + err.Error())
	}

	opts := []grpc.DialOption{
		grpc.WithPerRPCCredentials(jwt),
		grpc.WithInsecure(), // just for development
	}

	// create a new instance of our client application
	c, err := grpcclient.New(
		grpcclient.WithHost("localhost:9090"), // @todo: env var/parameter
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
