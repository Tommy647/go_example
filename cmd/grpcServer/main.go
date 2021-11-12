package main

import (
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	_db "github.com/Tommy647/go_example/internal/db"
	"github.com/Tommy647/go_example/internal/dbgreeter"
	_greeter "github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/grpcserver"
)

const (
	// environment variable names
	envGreeter = `GREETER` // which greeter to use
)

func main() {
	// create a tcp listener for your gRPC service
	listener, err := net.Listen("tcp", "0.0.0.0:9090") //nolint:gosec // development use only
	if err != nil {
		panic(err.Error())
	}

	// define any server options we want to apply
	var opts []grpc.ServerOption
	// create a new gRPC server
	gRPCServer := grpc.NewServer(opts...)
	// @todo: this grpcServer.GracefulStop()

	// decide which function to run
	var greeter grpcserver.GreetProvider = _greeter.New()
	if strings.EqualFold(os.Getenv(envGreeter), "db") { // picked up by the linter, this is func ignores case
		db, err := _db.NewConnection()
		if err != nil {
			panic("database" + err.Error())
		}
		greeter = dbgreeter.New(db)
	}

	// 'register' our gRPC service with the newly created gRPC server
	go_example.RegisterHelloServiceServer(gRPCServer, grpcserver.New(greeter))
	// enable reflection for development, allows us to see the gRPC schema
	reflection.Register(gRPCServer)
	// let the user know we got this far
	log.Print("starting grpcServer")
	// serve the grpc server on the tcp listener - this blocks until told to close
	log.Fatal(gRPCServer.Serve(listener))
}
