package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	for _, v := range []string{
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
		"POSTGRES_SSLMODE",
	} {
		defer os.Setenv(v, os.Getenv(v))
		os.Unsetenv(v)
	}
	c, err := LoadConfig()
	assert.Equal(t, "required key POSTGRES_HOST missing value", err.Error())
	assert.Empty(t, c)

	os.Setenv("POSTGRES_PORT", "42")
	os.Setenv("POSTGRES_HOST", "POSTGRES_HOST")
	os.Setenv("POSTGRES_USER", "POSTGRES_USER")
	os.Setenv("POSTGRES_PASSWORD", "POSTGRES_PASSWORD")
	os.Setenv("POSTGRES_DB", "POSTGRES_DB")
	os.Setenv("POSTGRES_SSLMODE", "POSTGRES_SSLMODE")

	c, err = LoadConfig()
	assert.Nil(t, err)
	assert.Equal(t, 42, c.Port)
	assert.Equal(t, "POSTGRES_HOST", c.Host)
	assert.Equal(t, "POSTGRES_USER", c.User)
	assert.Equal(t, "POSTGRES_PASSWORD", c.Password)
	assert.Equal(t, "POSTGRES_DB", c.Database)
	assert.Equal(t, "POSTGRES_SSLMODE", c.SSLMode)
}
