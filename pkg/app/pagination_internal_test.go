package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination(t *testing.T) {
	p := Pagination{}
	assert.Equal(t, defaultPaginationLimit, p.GetLimit())
	assert.Equal(t, uint64(0), p.GetOffset())

	p.Limit = maxPaginationLimit + 1
	p.Offset = 123
	assert.Equal(t, maxPaginationLimit, p.GetLimit())
	assert.Equal(t, uint64(123), p.GetOffset())

	p.Limit = 9999
	assert.Equal(t, uint64(1000), p.GetCustomLimit(123, 1000))
	p.Limit = 0
	assert.Equal(t, uint64(123), p.GetCustomLimit(123, 1000))
}
