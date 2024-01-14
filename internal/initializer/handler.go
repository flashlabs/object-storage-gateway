package initializer

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	objectResource "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource"
	objectResourceRead "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource/read"
	objectResourceReplace "github.com/spacelift-io/homework-object-storage/internal/handler/object/resource/replace"
)

func Handler(r *mux.Router) error {
	log.Println("Initializing handlers")

	r.HandleFunc(objectResource.PatternPUT, objectResourceReplace.Handle).Methods(http.MethodPut)
	r.HandleFunc(objectResource.PatternGET, objectResourceRead.Handle).Methods(http.MethodGet)

	return nil
}
