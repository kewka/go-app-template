package app_test

import (
	"context"
	"testing"

	"github.com/kewka/go-app-template/pkg/app"
	"github.com/kewka/go-app-template/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewItemsService(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	assert.NotNil(t, app.NewItemsService(dbpool))
}

func TestItemsList(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	testhelpers.ClearItems(dbpool)
	created := map[int]app.ItemModel{}
	for _, name := range []string{"first", "second", "third"} {
		item := testhelpers.CreateItem(dbpool, name)
		created[item.ID] = item
	}
	itemsService := app.NewItemsService(dbpool)
	res, err := itemsService.List(context.Background(), &app.ItemsListRequest{})
	assert.Nil(t, err)
	assert.Equal(t, len(created), res.Count)
	assert.Equal(t, len(created), len(res.Items))
	for _, i := range res.Items {
		assert.Equal(t, created[i.ID].Name, i.Name)
	}
}

func TestItemsGetByID(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	testhelpers.ClearItems(dbpool)
	itemsService := app.NewItemsService(dbpool)

	_, err := itemsService.GetByID(context.Background(), 123)
	assert.Equal(t, app.ErrNotFound, err)

	created := testhelpers.CreateItem(dbpool, "item")
	res, err := itemsService.GetByID(context.Background(), created.ID)
	assert.Nil(t, err)
	assert.Equal(t, created.ID, res.ID)
	assert.Equal(t, created.Name, res.Name)
}
