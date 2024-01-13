package initializer

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	objectResource "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource"
	objectResourceCreate "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource/create"
	objectResourceRead "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource/read"
)

func Handler(c context.Context, r *mux.Router) error {
	log.Println("Initializing handlers")

	r.HandleFunc(objectResource.PatternPUT, objectResourceCreate.Handle).Methods(http.MethodPut)
	r.HandleFunc(objectResource.PatternGET, objectResourceRead.Handle).Methods(http.MethodGet)

	return nil
}
