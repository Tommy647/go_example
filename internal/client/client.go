package client

import (
	"context"
	"log"
	"sync"

	"github.com/Tommy647/grpc"
)

// Client to handle making the gRPC request to the server
type Client struct {
	// client connected to the server
	client  grpc.HelloWorldServiceClient // client for gRPC requests
	workers int                          // workers to create for processing requests
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
	c := &Client{
		workers: 4, // default value
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Run sends a request to the server and logs the response
func (c Client) Run(ctx context.Context, names ...string) {
	// if we have not been initialised correctly, just exit
	if c.client == nil {
		return
	}
	// queue to hold the inputs
	queue := make(chan *grpc.HelloRequest)

	// wait group, so we can wait for concurrent threads to finish
	wg := &sync.WaitGroup{}

	// create the workers as go routines
	for i := 0; i < c.workers; i++ {
		wg.Add(1)
		// create a worker to handle the requests concurrently
		go c.requestWorker(ctx, wg, queue)
	}

	// send a request off for each name
	if len(names) == 0 {
		queue <- &grpc.HelloRequest{}
	}

	for i := range names {
		queue <- &grpc.HelloRequest{Name: names[i]}
	}
	// close the queue - we have successfully added all the work
	close(queue)
	// wait for the wait group to unlock
	wg.Wait()
}

func (c *Client) requestWorker(ctx context.Context, wg *sync.WaitGroup, queue <-chan *grpc.HelloRequest) {
	defer wg.Done()
	for { // forever!
		select {
		case request, ok := <-queue:
			if !ok {
				// channel has been closed, queue is empty, so we exit here
				return
			}
			resp, err := c.client.HelloWorld(ctx, request)
			if err != nil {
				log.Println("error messaging server", err.Error())
				continue
			}
			log.Println("Message:", resp.GetResponse())
		case <-ctx.Done():
			// we timed out
			return
		}
	}
}
