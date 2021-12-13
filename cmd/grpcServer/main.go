package main

import (
	"database/sql"
	"fmt"
	_greeter "github.com/Tommy647/go_example/internal/greeter"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq" // special: we need to include this package here to ensure the drivers load, but we do not need the code
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Tommy647/go_example"
	"github.com/Tommy647/go_example/internal/dbgreeter"
	"github.com/Tommy647/go_example/internal/grpcserver"
)

const (
	// environment variable names
	envGreeter = `GREETER`     // which greeter to use
	dbHost     = `DB_HOST`     // database host
	dbPort     = `DB_PORT`     // database port
	dbUser     = `DB_USER`     // database user
	dbPassword = `DB_PASSWORD` // database password
	dbDbname   = `DB_DBNAME`   // database name
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
	var greeter grpcserver.GreetProvider
	switch strings.ToLower(os.Getenv(envGreeter)) {
	case "db":
		log.Print("database selected")
		db, err := sql.Open("postgres", getPostgresConnection())
		if err != nil {
			panic("database" + err.Error())
		}
		greeter = dbgreeter.New(db)

		// 'register' our gRPC service with the newly created gRPC server
		go_example.RegisterHelloServiceServer(gRPCServer, grpcserver.New(greeter))
		// enable reflection for development, allows us to see the gRPC schema
		reflection.Register(gRPCServer)
		// let the user know we got this far
		log.Print("starting grpcServer")
		// serve the grpc server on the tcp listener - this blocks until told to close
		log.Fatal(gRPCServer.Serve(listener))
	default:
		greeter = _greeter.New()
	}

	go_example.RegisterCustomGreeterServiceServer(gRPCServer, grpcserver.NewGreeter(greeter))
}

// getPostgresConnection string we need to open the connection
// gets connection details from the environment for now @todo: replace with viper
func getPostgresConnection() string {
	host := os.Getenv(dbHost)
	port, err := strconv.Atoi(os.Getenv(dbPort))
	if err != nil {
		panic("port must be a number " + err.Error())
	}
	user := os.Getenv(dbUser)
	password := os.Getenv(dbPassword)
	dbname := os.Getenv(dbDbname)

	return fmt.Sprintf(
		`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname,
	)
}
