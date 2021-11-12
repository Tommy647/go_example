package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	vault "github.com/hashicorp/vault/api"
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
	// connect to vault?
	_, _ = getSecretCert()

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

func getSecretCert() (string, error) {
	config := vault.DefaultConfig()

	client, err := vault.NewClient(config)
	if err != nil {
		return "", err
	}

	jwt, err := os.ReadFile("/etc/vault/service_account")
	if err != nil {
		return "", err
	}
	params := map[string]interface{}{
		"jwt":  string(jwt),
		"role": "dev-role-k8s", // the name of the role in Vault that was created with this app's Kubernetes service account bound to it
	}

	// log in to Vault's Kubernetes auth method
	resp, err := client.Logical().Write("auth/kubernetes/login", params)
	if err != nil {
		return "", fmt.Errorf("unable to log in with Kubernetes auth: %w", err)
	}
	if resp == nil || resp.Auth == nil || resp.Auth.ClientToken == "" {
		return "", fmt.Errorf("login response did not return client token")
	}
	// now you will use the resulting Vault token for making all future calls to Vault
	client.SetToken(resp.Auth.ClientToken)

	// get secret from Vault
	secret, err := client.Logical().Read("kv-v2/data/creds")
	if err != nil {
		return "", fmt.Errorf("unable to read secret: %w", err)
	}

	data, ok := secret.Data["data"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("data type assertion failed: %T %#v", secret.Data["data"], secret.Data["data"])
	}

	// data map can contain more than one key-value pair, in this case we're just grabbing one of them
	key := "password"
	value, ok := data[key].(string)
	if !ok {
		return "", fmt.Errorf("value type assertion failed: %T %#v", data[key], data[key])
	}

	return value, nil
}
