package database

import (
	"cheeeasy2501/book-library/internal/config"
	"github.com/tsenart/nap"
	"time"
)

var Instance *Database

type Database struct {
	Conn *nap.DB
}

func NewDatabaseInstance() {
	Instance = &Database{}
}

func (db *Database) OpenConnection(config *config.Config) error {
	connectionString, err := config.GetConnectionString()
	if err != nil {
		return err
	}
	//TODO: check how set multiple connections :range by map
	connection, err := nap.Open(config.Databases[0].Driver, connectionString)
	if err != nil {
		return err
	}
	connection.SetMaxOpenConns(int(config.Databases[0].MaxOpenConnectionLifetime * time.Minute))
	connection.SetMaxIdleConns(int(config.Databases[0].MaxOpenIdleConnectionLifetime * time.Minute))
	Instance.Conn = connection

	return nil
}

func (db *Database) CloseConnection() {
	err := Instance.Conn.Close()
	if err != nil {
		return
	}
}
