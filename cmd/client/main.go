// main - client makes request against our gRPC server
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	_grpc "github.com/Tommy647/grpc"
	"github.com/Tommy647/grpc/internal/client"
)

func main() {
	log.Println("client starting") // prove the client is up

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
	// start our client running
	c.Run(context.Background())
}
