package initializer_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

func TestHandler(t *testing.T) {
	assert.Nil(t, initializer.Handler(mux.NewRouter()))
}
