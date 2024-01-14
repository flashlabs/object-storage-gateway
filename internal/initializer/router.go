package initializer

import (
	"github.com/gorilla/mux"
)

func Router() (*mux.Router, error) {
	return mux.NewRouter(), nil
}
