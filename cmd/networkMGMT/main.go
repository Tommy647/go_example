package main

import (
	"log"

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

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
