package service_test

import (
	"context"
	"testing"

	"github.com/kewka/go-app-template/pkg/service"
	"github.com/kewka/go-app-template/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestListItems(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	svc := service.New(dbpool)
	testhelpers.ClearItems(dbpool)
	created := map[int]service.ItemModel{}
	for _, name := range []string{"first", "second", "third"} {
		item := testhelpers.CreateItem(dbpool, name)
		created[item.ID] = item
	}
	res, err := svc.ListItems(context.Background(), &service.ItemsListRequest{})
	assert.Nil(t, err)
	assert.Equal(t, len(created), res.Count)
	assert.Equal(t, len(created), len(res.Items))
	for _, i := range res.Items {
		assert.Equal(t, created[i.ID].Name, i.Name)
	}
}

func TestGetItemByID(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	testhelpers.ClearItems(dbpool)
	svc := service.New(dbpool)

	_, err := svc.GetItemByID(context.Background(), 123)
	assert.Equal(t, service.ErrNotFound, err)

	created := testhelpers.CreateItem(dbpool, "item")
	res, err := svc.GetItemByID(context.Background(), created.ID)
	assert.Nil(t, err)
	assert.Equal(t, created.ID, res.ID)
	assert.Equal(t, created.Name, res.Name)
}
