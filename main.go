package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

const (
	appInitializationFailedCode = 1
	port                        = 3000
)

func main() {
	log.Println("Spacelift Homework Object Storage")

	c := context.Background()

	if err := initApp(c); err != nil {
		fmt.Println("error while executing initApp: %w", err)

		os.Exit(appInitializationFailedCode)
	}

	log.Println("App done")
}

func initApp(c context.Context) error {
	err := initializer.Shards(c)
	if err != nil {
		return fmt.Errorf("error while initializing initializer.Shards: %w", err)
	}

	r, err := initializer.Router()
	if err != nil {
		return fmt.Errorf("error while executing initializer.Router: %w", err)
	}

	err = initializer.Handler(r)
	if err != nil {
		return fmt.Errorf("error while executing initializer.Handler: %w", err)
	}

	srv, err := initializer.Server(r, port)
	if err != nil {
		return fmt.Errorf("error while executing initializer.Server: %w", err)
	}

	err = srv.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error while executing srv.ListenAndServe: %w", err)
	}

	return nil
}
