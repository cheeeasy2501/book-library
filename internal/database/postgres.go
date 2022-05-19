package database

import (
	"github.com/cheeeasy2501/book-library/internal/config"
	_ "github.com/lib/pq"
	"github.com/tsenart/nap"
	"time"
)

type Database struct{}

func NewDatabaseConnection(cnf *config.DatabaseConfig) (*nap.DB, error) {
	db := &Database{}
	connection, err := db.OpenConnection(cnf)
	if err != nil {
		return nil, err
	}

	return connection, nil
}
func (db *Database) OpenConnection(cnf *config.DatabaseConfig) (*nap.DB, error) {
	connectionString, err := cnf.GetConnectionString()
	if err != nil {
		return nil, err
	}
	connection, err := nap.Open(cnf.Driver, connectionString)
	if err != nil {
		return nil, err
	}
	connection.SetMaxOpenConns(int(cnf.MaxOpenConnectionLifetime * time.Minute))
	connection.SetMaxIdleConns(int(cnf.MaxOpenIdleConnectionLifetime * time.Minute))
	err = connection.Ping()
	if err != nil {
		return nil, err
	}

	return connection, nil
}
