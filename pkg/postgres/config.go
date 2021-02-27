package postgres

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Database string `envconfig:"POSTGRES_DB" required:"true"`
	SSLMode  string `envconfig:"POSTGRES_SSLMODE" default:"disable"`
}

// ConnectionURL ...
func (c *Config) ConnectionURL() string {
	return fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SSLMode,
	)
}

// LoadConfig ...
func LoadConfig() (*Config, error) {
	c := Config{}
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}
	return &c, nil
}
