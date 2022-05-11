package database

import (
	"cheeeasy2501/book-library/internal/config"
	"database/sql"
	_ "database/sql/driver"
	"time"
)

var Instance *Database

type Database struct {
	Conn *sql.DB
}

func NewDatabaseInstance() {
	Instance = &Database{}
}

func (db *Database) OpenConnection(config *config.Config) {
	connectionString := config.GetConnectionString()

	connection, err := sql.Open(config.Database.Driver, connectionString)
	if err != nil {
		return
	}
	connection.SetMaxOpenConns(int(config.Database.MaxOpenConnectionLifetime * time.Minute))
	connection.SetMaxIdleConns(int(config.Database.MaxOpenIdleConnectionLifetime * time.Minute))
	Instance.Conn = connection
}

func (db *Database) CloseConnection() {
	err := Instance.Conn.Close()
	if err != nil {
		return
	}
}
