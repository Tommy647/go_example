// Package main a client to make requests against our servers
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

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

	opts := []grpc.DialOption{
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
	// Create generic request
	var request grpcclient.Requester
	// BasicGreeter request with options
	request = grpcclient.BasicGreeter{
		RequestOpts: grpcclient.RequestOpts{
			Context: ctx,
			Names:   []string{"Tom"},
		},
	}
	// Send the BasicGreeter request
	request.Request(c)
	// CustomGreeter request with options
	request = grpcclient.CustomGreeter{
		RequestOpts: grpcclient.RequestOpts{
			Context:  ctx,
			Names:    []string{"Tom", "Jimmy"},
			Greeting: "Welcome",
		},
	}
	// Send the CustomGreeter request
	request.Request(c)
}
