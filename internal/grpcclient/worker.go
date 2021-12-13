package grpcclient

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"sync"

	grpc "github.com/Tommy647/go_example"
)

type Message struct {
	Msg1 *grpc.HelloRequest
	Msg2 *grpc.CustomGreeterRequest
}

// requestWorker handles making requests to the grpc grpcServer
// @todo: figure out a generic type for the chan

// requestWorker handles making requests to the grpc grpcServer
// @todo: figure out a generic type for the chan
func (c *Client) requestWorkerStruct(ctx context.Context, wg *sync.WaitGroup, queue <-chan Message) {
	defer wg.Done()
	for { // forever!
		select {
		case request, ok := <-queue:
			if !ok {
				// channel has been closed, queue is empty, so we exit here
				return
			}
			if request.Msg1 != nil {
				resp, err := c.helloClient.Hello(ctx, request.Msg1)
				if err != nil {
					log.Println("error messaging grpcServer", err.Error())
					continue
				}
				log.Println("Message:", resp.GetResponse())
			}
			if request.Msg1 != nil {
				resp, err := c.greetingClient.CustomGreeter(ctx, request.Msg2)
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

// requestWorker handles making requests to the grpc grpcServer
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
