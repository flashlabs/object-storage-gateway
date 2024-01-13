package initializer

import (
	"context"
	"log"

	"github.com/gorilla/mux"
)

func Handler(c context.Context, r *mux.Router) error {
	log.Println("Initializing handlers")

	return nil
}
