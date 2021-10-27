package client

import (
	"context"
	"log"

	"github.com/Tommy647/grpc"
)

// Client to handle making the gRPC request to the server
type Client struct {
	// client connected to the server
	client grpc.HelloWorldServiceClient
}

// option for variadic configuration of our client
type option func(*Client)

// WithHelloWorldClient set the HelloWorldClient
func WithHelloWorldClient(hwc grpc.HelloWorldServiceClient) option {
	return func(c *Client) {
		c.client = hwc
	}
}

// New instance of our Client, accepts variadic options
func New(opts ...option) *Client {
	c := &Client{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Run sends a request to the server and logs the response
func (c Client) Run(ctx context.Context, names ...string) {
	if c.client == nil {
		return
	}

	if len(names) == 0 {
		// make a single blank request and exit
		resp, err := c.client.HelloWorld(ctx, &grpc.HelloRequest{})
		if err != nil {
			log.Println("error messaging server", err.Error())
			return
		}
		log.Println("Message: ", resp.Response)
		return
	}

	for i := range names {
		// send a request off for each name
		resp, err := c.client.HelloWorld(ctx, &grpc.HelloRequest{Name: names[i]})
		if err != nil {
			log.Println("error messaging server", err.Error())
			continue
		}
		log.Printf("Message: Given: %q %s\n", names[i], resp.Response)
	}
}
