package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	_db "github.com/Tommy647/go_example/internal/db"
	"github.com/Tommy647/go_example/internal/dbgreeter"
	_greeter "github.com/Tommy647/go_example/internal/greeter"
	"github.com/Tommy647/go_example/internal/grpcserver"
	"github.com/Tommy647/go_example/internal/interceptor"
	"github.com/Tommy647/go_example/internal/logger"
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
	// get the zap structured logger
	if err = logger.New(`go_example_grpc`); err != nil {
		panic(err.Error())
	}
	defer logger.Close()
	logger.Debug(context.Background(), "application logger started")

	// define any server options we want to apply
	var opts = []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.WithTrace(),
			interceptor.WithAuth(),
		),
	}

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

	// We grab an instance of our fruitServer
	fS := grpcserver.NewFS()
	// 'register' our gRPC service with the newly created gRPC server
	go_example.RegisterHelloServiceServer(gRPCServer, grpcserver.NewHS(getGreeter()))
	go_example.RegisterCoffeeServiceServer(gRPCServer, grpcserver.NewCS(getCoffeeGreeter()))
	go_example.RegisterFruitServiceServer(gRPCServer, fS)

	// enable reflection for development, allows us to see the gRPC schema
	reflection.Register(gRPCServer)
	// let the user know we got this far
	logger.Info(context.Background(), "starting grpcServer", zap.String("port", port))

	// serve the grpc server on the tcp listener - this blocks until told to close
	if err := gRPCServer.Serve(listener); err != nil {
		logger.Fatal(context.Background(), "grpc service failed", zap.Error(err))
	}
}

// getGreeter decide which greeter service to use
func getGreeter() grpcserver.GreetProvider {
	if strings.EqualFold(os.Getenv(envGreeter), "db") { // picked up by the linter, this func ignores case
		logger.Info(context.Background(), "using database greeter")
		db, err := _db.NewConnection()
		if err != nil {
			panic("database" + err.Error())
		}
		return dbgreeter.New(db)
	}
	logger.Info(context.Background(), "using string greeter")
	return _greeter.New()
}

// getCoffeeGreeter assumes the use of a DB, however during the preparation of the DBGreeter
// the operation can revert to a basicGreeter which serves drinks from strings
func getCoffeeGreeter() grpcserver.CoffeeProvider {
	logger.Info(context.Background(), "using database greeter")
	db, err := _db.NewConnection()
	if err != nil {
		panic("database" + err.Error())
	}
	return dbgreeter.New(db)
}
