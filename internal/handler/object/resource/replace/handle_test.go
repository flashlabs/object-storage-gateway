package replace_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spacelift-io/homework-object-storage/internal/initializer"
)

const (
	url = "/object/asd123"
)

func TestHandle(t *testing.T) {
	r, err := initializer.Router()
	assert.Nil(t, err)

	err = initializer.Handler(r)
	assert.Nil(t, err)

	req, err := http.NewRequest(http.MethodPut, url, nil)
	assert.Nil(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}
