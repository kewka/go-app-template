package service

import (
	"context"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Service ...
type Service interface {
	HTTPHandler() http.Handler
	ListItems(ctx context.Context, req *ItemsListRequest) (*ItemsListResponse, error)
	GetItemByID(ctx context.Context, id int) (*ItemModel, error)
}

type service struct {
	dbpool *pgxpool.Pool
}

// New ...
func New(dbpool *pgxpool.Pool) Service {
	return &service{
		dbpool: dbpool,
	}
}
