package config

import (
	"cheeeasy2501/book-library/internal/database"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Api       *ApiConfig
	Databases map[int16]*DatabaseConfig
}

func (c *Config) LoadEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}

	return nil
}

// postgres://postgres:123456@127.0.0.1:5432/dummy
func (c *Config) GetConnectionString() (string, error) {
	if _, ok := c.Databases[0]; !ok {
		return "", database.DatabaseConfigNotFound
	}
	return fmt.Sprintf("%s:%s@%s:%s/%s", c.Databases[0].User, c.Databases[0].Password, c.Databases[0].Address, c.Databases[0].Port, c.Databases[0].Name), nil
}

func NewConfig() *Config {
	return &Config{}
}

type ApiConfig struct {
	Port string `envconfig:"API_PORT" default:"8080"`
}

type DatabaseConfig struct {
	Driver   string `envconfig:"DB_DRIVER" default:"postgres"`
	User     string `envconfig:"DB_USER" default:"root"`
	Password string `envconfig:"DB_PASSWORD" default:"root"`
	Address  string `envconfig:"DB_ADDRESS" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"book-library"`

	MaxOpenConnection             string        `envconfig:"DB_MAX_OPEN_CONNECTION" default:"5"`
	MaxOpenConnectionLifetime     time.Duration `envconfig:"DB_MAX_OPEN_CONNECTION_LIFETIME" default:"10m"`
	MaxOpenIdleConnection         string        `envconfig:"DB_MAX_IDLE_CONNECTION" default:"1"`
	MaxOpenIdleConnectionLifetime time.Duration `envconfig:"DB_MAX_IDLE_CONNECTION_LIFETIME" default:"60m"`
}
