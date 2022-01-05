package main

import (
	"log"

	"github.com/Tommy647/go_example/internal/grpcserver"

	"github.com/Tommy647/go_example/internal/router"

	"github.com/Tommy647/go_example/internal/db"
)

func Run() error {
	// responsible for initializing and starting
	// our gRPC server
	deviceStore, err := db.New()
	if err != nil {
		return err
	}

	err = deviceStore.Migrate()
	if err != nil {
		log.Println("cannot run migrations")
		return err
	}
	routerService := router.New(deviceStore)

	deviceHandler := grpcserver.NewRtrService(routerService)

	if err := deviceHandler.Serve(); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
