package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Common   *CommonConfig
	Auth     *AuthConfig
	Api      *ApiConfig
	Database *DatabaseConfig
}

func NewConfig() *Config {
	return &Config{}
}

type CommonConfig struct {
}

type AuthConfig struct {
	jwtKey string `envconfig:"JWT_KEY"`
}

func (ac *AuthConfig) GetJWTKey() string {
	return ac.jwtKey
}

type ApiConfig struct {
	Port string `envconfig:"API_PORT" default:"8080"`
}

type DatabaseConfig struct {
	Driver   string `envconfig:"DB_DRIVER" default:"postgres"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	Address  string `envconfig:"DB_ADDRESS" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"books-library"`

	MaxOpenConnection             string        `envconfig:"DB_MAX_OPEN_CONNECTION" default:"5"`
	MaxOpenConnectionLifetime     time.Duration `envconfig:"DB_MAX_OPEN_CONNECTION_LIFETIME" default:"10m"`
	MaxOpenIdleConnection         string        `envconfig:"DB_MAX_IDLE_CONNECTION" default:"1"`
	MaxOpenIdleConnectionLifetime time.Duration `envconfig:"DB_MAX_IDLE_CONNECTION_LIFETIME" default:"60m"`
}

// GetConnectionString postgres://postgres:123456@127.0.0.1:5432/dummy
func (dc *DatabaseConfig) GetConnectionString() (string, error) {
	return fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable", dc.Port, dc.Address, dc.User, dc.Password, dc.Name), nil
}

func (c *Config) LoadEnv() error {
	err := envconfig.Process("", c)
	if err != nil {
		return err
	}

	return nil
}
