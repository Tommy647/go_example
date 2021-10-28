// main - client makes request against our gRPC server
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	_grpc "github.com/Tommy647/grpc"
	"github.com/Tommy647/grpc/internal/client"
)

func main() {
	log.Println("client starting") // prove the client is up
	// create a new context that expires in 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// close the context when we leave this function
	defer cancel()

	opts := []grpc.DialOption{
		grpc.WithInsecure(), // just for development
	}
	// make a grpc connection to the server
	conn, err := grpc.Dial("localhost:9090", opts...)
	if err != nil {
		panic(err.Error()) // we can not continue
	}
	// create a new client on our connection
	hwClient := _grpc.NewHelloWorldServiceClient(conn)

	// create a new instance of our client application
	c := client.New(
		client.WithHelloWorldClient(hwClient), // give it our gRPC client
	)
	// start our client running with no input
	c.Run(ctx)
	// reuse the client and add some names
	c.Run(ctx, "Tom", "Orson", "Kurt")
}
