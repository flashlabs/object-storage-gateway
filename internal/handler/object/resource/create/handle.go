package create

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	process "github.com/spacelift-io/homework-object-storage/internal/process/object/resource/create"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP create resource handler started")

	vars := mux.Vars(r)
	log.Printf("Processing object with ID %q\n", vars["id"])

	_, err := process.Execute(r.Context(), process.Input{
		Vars:    vars,
		Payload: r.Body,
	})
	if err != nil {
		handleFailure(w, "Create process ended with error", http.StatusInternalServerError, err)

		return
	}

	// 200, 204 for successful update, 201 for content creation
	w.WriteHeader(http.StatusNoContent)

	log.Println("HTTP create resource handler ended")
}

func handleFailure(w http.ResponseWriter, message string, status int, err error) {
	log.Println(message, err)

	w.WriteHeader(status)

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Println("Error while writing to http.ResponseWriter", err)
	}
}
