package app

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

// ItemModel ...
type ItemModel struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ItemsService ...
type ItemsService interface {
	List(ctx context.Context, req *ItemsListRequest) (*ItemsListResponse, error)
	GetByID(ctx context.Context, id int) (*ItemModel, error)
}

type itemsService struct {
	dbpool *pgxpool.Pool
}

// NewItemsService ...
func NewItemsService(dbpool *pgxpool.Pool) ItemsService {
	return &itemsService{
		dbpool: dbpool,
	}
}

// ItemsListRequest ...
type ItemsListRequest struct {
	Pagination
}

// ItemsListResponse ...
type ItemsListResponse struct {
	Items []ItemModel `json:"items"`
	Count int         `json:"count"`
}

func (s *itemsService) List(ctx context.Context, req *ItemsListRequest) (*ItemsListResponse, error) {
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

func (s *itemsService) GetByID(ctx context.Context, id int) (*ItemModel, error) {
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
