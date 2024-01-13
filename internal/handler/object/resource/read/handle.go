package read

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.Println("HTTP read resource handler started")

	vars := mux.Vars(r)

	log.Printf("Processing object with ID %q\n", vars["id"])

	// 200 for ok, 404 for missing
	w.WriteHeader(http.StatusOK)

	log.Println("HTTP read resource handler ended")
}
