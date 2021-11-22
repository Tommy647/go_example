package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	_db "github.com/Tommy647/go_example/internal/db"
	"github.com/Tommy647/go_example/internal/dbgreeter"
	_greeter "github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/grpcserver"
	"github.com/Tommy647/go_example/internal/tls"
)

const (
	// environment variable names
	envGreeter = `GREETER` // which greeter to use
	envPort    = `PORT`    // port to listen on

	// address as we need to listen on all interfaces for docker
	address = `0.0.0.0:%s`
)

func main() {
	port := os.Getenv(envPort)
	if port == "" {
		panic("missing port, provide PORT env var")
	}

	// create a tcp listener for our gRPC service
	listener, err := net.Listen("tcp", fmt.Sprintf(address, port))
	if err != nil {
		panic(err.Error())
	}

	// define any server options we want to apply
	var opts = []grpc.ServerOption(nil)

	tlsConfig, err := tls.GetCertificates()
	if err != nil {
		panic("tls certificate error " + err.Error())
	}

	if tlsConfig != nil {
		opts = append(opts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	// create a new gRPC server
	gRPCServer := grpc.NewServer(opts...)
	// @todo: this grpcServer.GracefulStop()

	// 'register' our gRPC service with the newly created gRPC server
	go_example.RegisterHelloServiceServer(gRPCServer, grpcserver.New(getGreeter()))
	// enable reflection for development, allows us to see the gRPC schema
	reflection.Register(gRPCServer)
	// let the user know we got this far
	log.Print("starting grpcServer on ", port)
	// serve the grpc server on the tcp listener - this blocks until told to close
	log.Fatal(gRPCServer.Serve(listener))
}

// getGreeter decide which greeter service to use
func getGreeter() grpcserver.GreetProvider {
	if strings.EqualFold(os.Getenv(envGreeter), "db") { // picked up by the linter, this is func ignores case
		db, err := _db.NewConnection()
		if err != nil {
			panic("database" + err.Error())
		}
		return dbgreeter.New(db)
	}
	return _greeter.New()
}
