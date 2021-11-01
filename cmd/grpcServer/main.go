package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/grpcserver"
)

func main() {
	// create a tcp listener for your rGPC service
	listener, err := net.Listen("tcp", "0.0.0.0:9090") //nolint:gosec // development use only
	if err != nil {
		panic(err.Error())
	}

	// define any server options we want to apply
	var opts []grpc.ServerOption
	// create a new gRPC server
	gRPCServer := grpc.NewServer(opts...)
	// @todo: this grpcServer.GracefulStop()
	// 'register' our gRPC service with the newly created gRPC server
	go_example.RegisterHelloServiceServer(gRPCServer, grpcserver.New())
	// enable reflection for development, allows us to see the gRPC schema
	reflection.Register(gRPCServer)
	// let the user know we got this far
	log.Print("starting grpcServer")
	// serve the grpc server on the tcp listener - this blocks until told to close
	log.Fatal(gRPCServer.Serve(listener))
}
