package grpcclient

import (
	"context"
	"sync"

	"github.com/Tommy647/go_example"
)

// Run sends a request to the grpcServer and logs the response
func (c Client) Run(ctx context.Context, names ...string) {
	// if we have not been initialised correctly, just exit
	if c.client == nil {
		return
	}
	// queue to hold the inputs
	queue := make(chan *go_example.HelloRequest)

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
		queue <- &go_example.HelloRequest{}
	}

	for i := range names {
		queue <- &go_example.HelloRequest{Name: names[i]}
	}
	// close the queue - we have successfully added all the work
	close(queue)
	// wait for the wait group to unlock
	wg.Wait()
}
