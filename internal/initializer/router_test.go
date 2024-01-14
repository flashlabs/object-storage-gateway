package initializer_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

func TestRouter(t *testing.T) {
	r, err := initializer.Router()

	require.NoError(t, err)
	assert.Equal(t, mux.NewRouter(), r)
}
