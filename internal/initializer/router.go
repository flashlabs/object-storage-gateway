package initializer

import (
	"context"

	"github.com/gorilla/mux"
)

func Router(c context.Context) (*mux.Router, error) {
	r := mux.NewRouter()

	return r, nil
}
