package replace

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	process "github.com/spacelift-io/homework-object-storage/internal/process/object/resource/replace"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP create resource handler started")

	s := http.StatusCreated
	vars := mux.Vars(r)
	log.Printf("Processing object with ID %q\n", vars["id"])

	out, err := process.Execute(r.Context(), process.Input{
		Vars:          vars,
		Payload:       r.Body,
		ContentLength: r.ContentLength,
	})
	if err != nil {
		log.Println("Create process ended with error", err)

		s = http.StatusInternalServerError
	}

	if out.Replaced {
		s = http.StatusNoContent
	}

	w.WriteHeader(s)

	log.Println("HTTP create resource handler ended")
}
