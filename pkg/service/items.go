package service

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Masterminds/squirrel"
	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4"
)

// ItemModel ...
type ItemModel struct {
	ID   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// ItemsListRequest ...
type ItemsListRequest struct {
	Pagination
}

// ItemsListResponse ...
type ItemsListResponse struct {
	Items []ItemModel `json:"items" binding:"required"`
	Count int         `json:"count" binding:"required"`
}

func (s *service) ListItems(ctx context.Context, req *ItemsListRequest) (*ItemsListResponse, error) {
	ret := &ItemsListResponse{
		Items: []ItemModel{},
	}
	q := squirrel.Select("id", "name").
		From("items").
		Limit(req.GetLimit()).
		Offset(req.GetOffset()).
		PlaceholderFormat(squirrel.Dollar)
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.dbpool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		i := ItemModel{}
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		ret.Items = append(ret.Items, i)
	}
	if err := s.dbpool.QueryRow(ctx, "SELECT COUNT(*) FROM items").Scan(&ret.Count); err != nil {
		return nil, err
	}
	return ret, nil
}

func (s *service) GetItemByID(ctx context.Context, id int) (*ItemModel, error) {
	ret := &ItemModel{}
	err := s.dbpool.QueryRow(
		ctx,
		"SELECT id, name FROM items WHERE id = $1",
		id,
	).Scan(&ret.ID, &ret.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return ret, nil
}

// @Summary Get items.
// @Tags items
// @Success 200 {object} ItemsListResponse
// @Router /items [get]
func (s *service) handleItemsList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := s.ListItems(r.Context(), &ItemsListRequest{
			Pagination: RequestPagination(r),
		})
		if err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
		JSON(w, http.StatusOK, res)
	}
}

// @Summary Get item by id.
// @Tags items
// @Param id path int true "ID"
// @Success 200 {object} ItemModel
// @Success 404 {object} ErrorResponse
// @Router /items/{id} [get]
func (s *service) handleItemByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		res, err := s.GetItemByID(r.Context(), id)
		if err != nil {
			if err == ErrNotFound {
				Error(w, http.StatusNotFound, err)
			} else {
				Error(w, http.StatusInternalServerError, err)
			}
			return
		}
		JSON(w, http.StatusOK, res)
	}
}
