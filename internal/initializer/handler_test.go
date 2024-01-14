package initializer_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

func TestHandler(t *testing.T) {
	require.NoError(t, initializer.Handler(mux.NewRouter()))
}
