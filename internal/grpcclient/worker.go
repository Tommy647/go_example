package grpcclient

import (
	"context"
	"log"
	"sync"

	"google.golang.org/protobuf/proto"

	grpc "github.com/Tommy647/go_example"
)

// requestWorkerInterface handles making requests to the grpc grpcServer
// Uses generic proto.Message interface type to allow the method to be more generic
func (c *Client) requestWorkerInterface(ctx context.Context, wg *sync.WaitGroup, queue <-chan proto.Message) {
	defer wg.Done()
	for { // forever!
		select {
		case request, ok := <-queue:
			if !ok {
				// channel has been closed, queue is empty, so we exit here
				return
			}
			switch concrete := request.(type) {
			case *grpc.HelloRequest:
				resp, err := c.helloClient.Hello(ctx, concrete)
				if err != nil {
					log.Println("error messaging grpcServer", err.Error())
					continue
				}
				log.Println("Message:", resp.GetResponse())
			case *grpc.CustomGreeterRequest:
				resp, err := c.greetingClient.CustomGreeter(ctx, concrete)
				if err != nil {
					log.Println("error messaging grpcServer", err.Error())
					continue
				}
				log.Println("Message:", resp.GetResponse())
			}

		case <-ctx.Done():
			// we timed out
			return
		}
	}
}
