package read_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

const (
	url = "/object/asd123"
)

func TestHandle(t *testing.T) {
	c := context.Background()

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	req, err := http.NewRequestWithContext(c, http.MethodGet, url, nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
