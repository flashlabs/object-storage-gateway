package read

import (
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"

	process "github.com/spacelift-io/homework-object-storage/internal/process/object/resource/read"
	"github.com/spacelift-io/homework-object-storage/pkg"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP read resource handler started")

	vars := mux.Vars(r)
	log.Printf("Processing object with ID %q\n", vars["id"])

	out, err := process.Execute(r.Context(), process.Input{
		Vars: vars,
	})
	if err != nil && !errors.Is(err, pkg.ErrObjectNotExists) {
		log.Println("Read process execution ended with error", err)
	}

	if err != nil && errors.Is(err, pkg.ErrObjectNotExists) {
		log.Println("Object not found")
		w.WriteHeader(http.StatusNotFound)
	}

	entity := out.Entity
	if entity != nil {
		defer func(o *minio.Object) {
			ce := o.Close()
			if ce != nil {
				log.Println("error while closing the object: %w", ce)
			}
		}(entity)

		_, err = io.Copy(w, entity)
		if err != nil {
			log.Println("Copying entity into the writer ended with error", err)

			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	log.Println("HTTP read resource handler ended")
}
