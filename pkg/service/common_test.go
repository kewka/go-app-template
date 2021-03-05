package service_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kewka/go-app-template/pkg/service"
	"github.com/kewka/go-app-template/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestHandleIndex(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	svc := service.New(dbpool)
	handler := svc.HTTPHandler()
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	handler.ServeHTTP(w, r)
	res := w.Result()
	defer res.Body.Close()
	rawBody, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello world", string(rawBody))
}
