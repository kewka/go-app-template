package httphandler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kewka/go-app-template/pkg/app"
)

// @Summary Get items.
// @Tags items
// @Success 200 {array} app.ItemsListResponse
// @Router /items [get]
func handleItems(deps *Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := deps.ItemsService.List(r.Context(), &app.ItemsListRequest{
			Pagination: Pagination(r),
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		JSON(w, http.StatusOK, res)
	}
}

// @Summary Get item by id.
// @Tags items
// @Param id path int true "ID"
// @Success 200 {object} app.ItemModel
// @Success 404 {object} ErrorResponse
// @Router /items/{id} [get]
func handleItem(deps *Deps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		res, err := deps.ItemsService.GetByID(r.Context(), id)
		if err != nil {
			if err == app.ErrNotFound {
				Error(w, http.StatusNotFound, err)
			} else {
				Error(w, http.StatusInternalServerError, err)
			}
			return
		}
		JSON(w, http.StatusOK, res)
	}
}
