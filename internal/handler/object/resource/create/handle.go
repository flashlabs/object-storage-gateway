package create

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP create resource handler started")

	vars := mux.Vars(r)

	log.Printf("Processing object with ID %q\n", vars["id"])

	// 200,204 for successful update, 201 for content creation
	w.WriteHeader(http.StatusNoContent)

	log.Println("HTTP create resource handler ended")
}
