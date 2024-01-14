package initializer

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Server(r *mux.Router, port int) (*http.Server, error) {
	return &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}, nil
}
