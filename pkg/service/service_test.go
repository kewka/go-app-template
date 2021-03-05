package service_test

import (
	"testing"

	"github.com/kewka/go-app-template/pkg/service"
	"github.com/kewka/go-app-template/pkg/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testhelpers.Long(t)
	dbpool := testhelpers.SetupPostgres()
	defer dbpool.Close()
	svc := service.New(dbpool)
	assert.NotNil(t, svc)
}
