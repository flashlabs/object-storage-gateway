package initializer_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

func TestServer(t *testing.T) {
	r := mux.NewRouter()
	srv, err := initializer.Server(r, 3000)

	require.NoError(t, err)
	assert.Equal(t, &http.Server{
		Addr:              fmt.Sprintf(":%d", 3000),
		Handler:           r,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}, srv)
}
