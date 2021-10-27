package main

import (
	"log"
	"net"

	"github.com/Tommy647/grpc/internal/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_grpc "github.com/Tommy647/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err.Error())
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	// grpcServer.GracefulStop()
	_ = _grpc.File_grpc_proto

	_grpc.RegisterHelloWorldServiceServer(grpcServer, server.New())
	reflection.Register(grpcServer)
	log.Print("starting server")
	log.Fatal(grpcServer.Serve(listener))
}
