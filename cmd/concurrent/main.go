package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

const maxWorkers = 1

var devices = map[int]device{
	1: {id: 1},
	2: {id: 2},
	3: {id: 3},
	4: {id: 4},
}

var randomSource = rand.NewSource(647)
var randomGen = rand.New(randomSource)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	workload := make(chan int)
	prechecks := make(chan int)
	configchange := make(chan int)
	postchecks := make(chan int)
	done := make(chan int)
	go worker(ctx, workload, prechecks, noop)
	go worker(ctx, prechecks, configchange, preCheck)
	go worker(ctx, configchange, postchecks, configChange)
	go worker(ctx, postchecks, done, postCheck)

LOOP:
	for id, _ := range devices {
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
	id                int
	pre, config, post bool
}

func (d device) Print() {
	log.Printf("%d %#v", d.id, d)
}

type work func(context.Context, int) bool

// preCheck a device by ID
func preCheck(ctx context.Context, id int) bool {
	select {
	case <-ctx.Done():
		return false
	default: // nothing to do
	}
	time.Sleep(time.Duration(randomGen.Int63n(1000)) * time.Millisecond)
	log.Println("Precheck", id)
	if d, ok := devices[id]; ok {
		d.pre = true
		return true
	}
	return false
}

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
	time.Sleep(time.Duration(randomGen.Int63n(5000)) * time.Millisecond)
	log.Println("Config change", id)
	if d, ok := devices[id]; ok {
		d.config = true
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
	time.Sleep(time.Duration(randomGen.Int63n(1000)) * time.Millisecond)
	log.Println("Postcheck", id)
	if d, ok := devices[id]; ok {
		d.post = true
		return true
	}
	return false
}

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
