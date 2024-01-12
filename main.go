package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

const (
	appInitializationFailedCode = 1
)

func main() {
	log.Println("Spacelift Homework Object Storage")

	if err := initApp(); err != nil {
		fmt.Println("error while executing initApp: %w", err)

		os.Exit(appInitializationFailedCode)
	}

	log.Println("App done")
}

func initApp() error {
	err := initializer.Shards()
	if err != nil {
		return fmt.Errorf("error while initializing initializer.Shards: %w", err)
	}

	err = initializer.Handler()
	if err != nil {
		return fmt.Errorf("error while executing initializer.Handler: %w", err)
	}

	return nil
}
