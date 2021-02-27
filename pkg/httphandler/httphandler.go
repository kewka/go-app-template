package httphandler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kewka/go-app-template/pkg/app"
)

// Deps ...
type Deps struct {
	ItemsService app.ItemsService
}

// New ...
// @title App API
// @version 1.0
func New(deps *Deps) http.Handler {
	r := chi.NewRouter()
	r.Get("/", handleIndex())
	r.Get("/items", handleItems(deps))
	r.Get("/items/{id}", handleItem(deps))
	return r
}
