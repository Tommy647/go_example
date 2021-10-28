package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_grpc "github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/grpcserver"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090") //nolint:gosec // development use only
	if err != nil {
		panic(err.Error())
	}

	var opts []grpc.ServerOption
	gRPCServer := grpc.NewServer(opts...)
	// @todo: this grpcServer.GracefulStop()

	_grpc.RegisterHelloWorldServiceServer(gRPCServer, grpcserver.New())
	reflection.Register(gRPCServer)
	log.Print("starting grpcServer")
	log.Fatal(gRPCServer.Serve(listener))
}
