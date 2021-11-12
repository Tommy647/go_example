// Package main an example of concurrency pipelining

package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

// maxWorkers number of workers we will create in go routines
const maxWorkers = 4

// oneSecond value to generate a random number up to one second in length
const oneSecond = 1000
const fiveSecond = 5000

// timeout duration for the entire process
const timeout = 10 * time.Second

// randomSeed for our number generator
const randomSeed = 647

// devices dummy data to use
var devices = map[int]device{
	1: {id: 1},
	2: {id: 2},
	3: {id: 3},
	4: {id: 4},
}

// randomGen for random numbers
var randomGen = rand.New(rand.NewSource(randomSeed)) //nolint:gosec // basic random numbers not for crypto

func main() {
	// context so we can handle timeouts
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// channels for passing work through our pipeline
	workload := make(chan int)
	prechecks := make(chan int)
	configchange := make(chan int)
	postchecks := make(chan int)
	done := make(chan int)

	// create workers for each stage
	go worker(ctx, workload, prechecks, noop)
	go worker(ctx, prechecks, configchange, preCheck)
	go worker(ctx, configchange, postchecks, configChange)
	go worker(ctx, postchecks, done, postCheck)

LOOP:
	for id := range devices {
		log.Println("adding", id)
		select {
		case workload <- id: // blocks until worker is ready to accept
		case <-ctx.Done():
			log.Println("context in", ctx.Err().Error())
			break LOOP
		}
	}
	close(workload)

	log.Println("waiting")

	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err().Error())
			return
		case id, open := <-done:
			if !open {
				log.Println("done")
				return
			}
			log.Printf("%d finished", id)
		}
	}
}

type device struct {
	id int
}

// Print helper function for logging
func (d device) Print() { log.Printf("%d %#v", d.id, d) }

// work func type so we can change the function called by workers
type work func(context.Context, int) bool

// preCheck a device by ID
func preCheck(ctx context.Context, id int) bool {
	t := time.NewTimer(time.Duration(randomGen.Int63n(oneSecond)) * time.Millisecond)
	select {
	case <-ctx.Done():
		return false
	case <-t.C: // wait for the timer to trigger, alternative to the time.Sleep
		t.Stop()
	}

	log.Println("Precheck", id)
	if _, ok := devices[id]; ok {
		return true
	}
	return false
}

// noop a no-op work function
func noop(context.Context, int) bool {
	return true
}

// configChange a device by ID
func configChange(ctx context.Context, id int) bool {
	select {
	case <-ctx.Done():
		return false
	default: // nothing to do
	}
	time.Sleep(time.Duration(randomGen.Int63n(fiveSecond)) * time.Millisecond)
	log.Println("Config change", id)
	if _, ok := devices[id]; ok {
		return true
	}
	return false
}

// postCheck a device by ID
func postCheck(ctx context.Context, id int) bool {
	select {
	case <-ctx.Done():
		return false
	default: // nothing to do
	}
	time.Sleep(time.Duration(randomGen.Int63n(oneSecond)) * time.Millisecond)
	log.Println("Postcheck", id)
	if _, ok := devices[id]; ok {
		return true
	}
	return false
}

// worker our generic worker function, accepts a channel to watch for input,
// a channel to push completed work to and a work func (f) to perform on the work unit
func worker(ctx context.Context, input <-chan int, output chan<- int, f work) {
	wg := sync.WaitGroup{}
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					log.Println("context", ctx.Err().Error())
					return
				case id, open := <-input: // blocks until channel presents
					if !open {
						return
					}
					_ = f(ctx, id)
					output <- id
				}
			}
		}()
	}
	wg.Wait()
	close(output)
}
