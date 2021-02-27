package testhelpers

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kewka/go-app-template/pkg/app"
)

// ClearItems ...
func ClearItems(dbpool *pgxpool.Pool) {
	_, err := dbpool.Exec(context.Background(), "DELETE FROM items")
	if err != nil {
		panic(err)
	}
}

// CreateItem ...
func CreateItem(dbpool *pgxpool.Pool, name string) app.ItemModel {
	ret := app.ItemModel{}
	err := dbpool.QueryRow(
		context.Background(),
		"INSERT INTO items (name) VALUES ($1) RETURNING id, name",
		name,
	).Scan(&ret.ID, &ret.Name)
	if err != nil {
		panic(err)
	}
	return ret
}
