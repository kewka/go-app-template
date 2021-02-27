package postgres_test

import (
	"context"
	"testing"

	"github.com/kewka/go-app-template/pkg/postgres"
	"github.com/kewka/go-app-template/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNewPool(t *testing.T) {
	testhelpers.Long(t)
	cfg, err := postgres.LoadConfig()
	assert.Nil(t, err)
	pool, err := postgres.NewPool(context.Background(), cfg)
	assert.Nil(t, err)
	assert.NotNil(t, pool)
	pool.Close()
}
