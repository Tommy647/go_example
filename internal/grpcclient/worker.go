package grpcclient

import (
	"context"
	"log"
	"sync"

	grpc "github.com/Tommy647/go_example"
)

// requestWorker handles making requests to the grpc grpcServer
func (c *Client) requestHelloWorker(ctx context.Context, wg *sync.WaitGroup, queue <-chan *grpc.HelloRequest) {
	defer wg.Done()
	for { // forever!
		select {
		case request, ok := <-queue:
			if !ok {
				// channel has been closed, queue is empty, so we exit here
				return
			}
			resp, err := c.helloClient.Hello(ctx, request)
			if err != nil {
				log.Println("error messaging grpcServer", err.Error())
				continue
			}
			log.Println("Message:", resp.GetResponse())
		case <-ctx.Done():
			// we timed out
			return
		}
	}
}

// requestWorker handles making requests to the grpc grpcServer
func (c *Client) requestCustomGreeterWorker(ctx context.Context, wg *sync.WaitGroup, queue <-chan *grpc.CustomGreeterRequest) {
	defer wg.Done()
	for { // forever!
		select {
		case request, ok := <-queue:
			if !ok {
				// channel has been closed, queue is empty, so we exit here
				return
			}
			resp, err := c.greetingClient.CustomGreeter(ctx, request)
			if err != nil {
				log.Println("error messaging grpcServer", err.Error())
				continue
			}
			log.Println("Message:", resp.GetResponse())
		case <-ctx.Done():
			// we timed out
			return
		}
	}
}
