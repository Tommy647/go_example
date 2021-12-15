package grpcclient

import (
	"github.com/Tommy647/go_example"
	"google.golang.org/protobuf/proto"
	"sync"
)

// Run sends a request to the grpcServer and logs the response
func (c Client) Run(requestType string, opts RequestOpts) {
	// if we have not been initialised correctly, just exit
	if c.helloClient == nil {
		return
	}
	// wait group, so we can wait for concurrent threads to finish
	wg := &sync.WaitGroup{}
	names := opts.Names

	switch requestType {
	case "BasicGreeter":
		// queue to hold the inputs
		queue := make(chan proto.Message)

		// create the workers as go routines
		for i := 0; i < c.workers; i++ {
			wg.Add(1)
			// create a worker to handle the requests concurrently
			go c.requestWorkerInterface(opts.Context, wg, queue)
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
	case "CustomGreeter":
		// queue to hold the inputs
		queue := make(chan proto.Message)

		// create the workers as go routines
		for i := 0; i < c.workers; i++ {
			wg.Add(1)
			// create a worker to handle the requests concurrently
			go c.requestWorkerInterface(opts.Context, wg, queue)
		}

		// send a request off for each name
		if len(names) == 0 {
			queue <- &go_example.CustomGreeterRequest{}
		}

		for i := range names {
			queue <- &go_example.CustomGreeterRequest{Name: names[i], Greeting: opts.Greeting}
		}
		// close the queue - we have successfully added all the work
		close(queue)
	}

	// wait for the wait group to unlock
	wg.Wait()
}
