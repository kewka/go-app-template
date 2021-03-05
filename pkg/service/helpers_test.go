package service_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kewka/go-app-template/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	w := httptest.NewRecorder()
	err := errors.New("server error")
	service.Error(w, http.StatusInternalServerError, err)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("content-type"))
	expected := service.ErrorResponse{
		Message: err.Error(),
	}
	actual := service.ErrorResponse{}
	assert.Nil(t, json.NewDecoder(res.Body).Decode(&actual))
	assert.Equal(t, expected, actual)
}

func TestJSON(t *testing.T) {
	w := httptest.NewRecorder()
	data := map[string]string{
		"data": "hello world",
	}
	service.JSON(w, http.StatusCreated, data)
	res := w.Result()
	defer res.Body.Close()
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.Equal(t, "application/json", res.Header.Get("content-type"))
	actual := map[string]string{}
	assert.Nil(t, json.NewDecoder(res.Body).Decode(&actual))
	assert.Equal(t, data, actual)
}

func TestBindJSON(t *testing.T) {
	v := map[string]string{}
	service.BindJSON(strings.NewReader(`{"message": "hello world"}`), &v)
	assert.Equal(t, "hello world", v["message"])
}

func TestPagination(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	p := service.RequestPagination(r)
	assert.Equal(t, uint64(0), p.Limit)
	assert.Equal(t, uint64(0), p.Offset)

	r = httptest.NewRequest(http.MethodGet, "/?limit=123&offset=321", nil)
	p = service.RequestPagination(r)
	assert.Equal(t, uint64(123), p.Limit)
	assert.Equal(t, uint64(321), p.Offset)
}
