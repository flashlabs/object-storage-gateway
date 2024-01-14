package replace_test

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
	url    = "/object/asd123"
	url404 = "/object/thisIsParamWithLengthMoreThan32Chars"
)

func TestHandle(t *testing.T) {
	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	c := context.Background()

	req, err := http.NewRequestWithContext(c, http.MethodPut, url, nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestHandleRouting(t *testing.T) {
	tt := []struct {
		routeVariable  string
		expectedStatus int
	}{
		{url, http.StatusInternalServerError},
		{url404, http.StatusNotFound},
	}

	r, err := initializer.Router()
	require.NoError(t, err)

	err = initializer.Handler(r)
	require.NoError(t, err)

	c := context.Background()
	for _, s := range tt {
		req, err := http.NewRequestWithContext(c, http.MethodPut, s.routeVariable, nil)
		require.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, s.expectedStatus, rr.Code)
	}
}
