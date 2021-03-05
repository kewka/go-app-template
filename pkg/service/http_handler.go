package service

import (
	"net/http"

	"github.com/go-chi/chi"
)

// @title Service HTTP API.
// @version 1.0
func (s *service) HTTPHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", s.handleIndex())
	r.Get("/items", s.handleItemsList())
	r.Get("/items/{id}", s.handleItemByID())
	return r
}
