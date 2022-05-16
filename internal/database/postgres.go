package database

import (
	"github.com/cheeeasy2501/book-library/internal/config"
	_ "github.com/lib/pq"
	"github.com/tsenart/nap"
	"time"
)

var Instance *Database

type Database struct {
	Conn *nap.DB
}

func SetNewDatabaseInstance() {
	Instance = &Database{}
}

func (db *Database) OpenConnection(config *config.Config) error {
	//TODO: config is empty!
	connectionString, err := config.GetConnectionString()
	if err != nil {
		return err
	}
	//TODO: check how set multiple connections :range by map
	connection, err := nap.Open(config.Database.Driver, connectionString)
	if err != nil {
		return err
	}
	connection.SetMaxOpenConns(int(config.Database.MaxOpenConnectionLifetime * time.Minute))
	connection.SetMaxIdleConns(int(config.Database.MaxOpenIdleConnectionLifetime * time.Minute))
	err = connection.Ping()
	if err != nil {
		return err
	}

	Instance.Conn = connection

	return nil
}

func (db *Database) CloseConnection() {
	err := Instance.Conn.Close()
	if err != nil {
		return
	}
}
