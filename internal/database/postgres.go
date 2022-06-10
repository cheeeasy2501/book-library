package database

import (
	"context"
	"database/sql"
	"github.com/cheeeasy2501/book-library/internal/config"
	_ "github.com/lib/pq"
	"github.com/tsenart/nap"
	"time"
)

type Database struct {
	cnf  *config.DatabaseConfig
	Conn *nap.DB
}

func NewDatabase(cnf *config.DatabaseConfig) *Database {
	return &Database{
		cnf: cnf,
	}
}

func (db *Database) OpenConnection() error {
	connectionString, err := db.cnf.GetConnectionString()
	if err != nil {
		return err
	}
	connection, err := nap.Open(db.cnf.Driver, connectionString)
	if err != nil {
		return err
	}
	connection.SetMaxOpenConns(int(db.cnf.MaxOpenConnectionLifetime * time.Minute))
	connection.SetMaxIdleConns(int(db.cnf.MaxOpenIdleConnectionLifetime * time.Minute))
	err = connection.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Реализуем обертки, чтобы вернуть Session интерфейс
//func (db *Database) Prepare(query string) (nap.Stmt, error) {
//	return db.Conn.Prepare(query)
//}
//func (db *Database) PrepareContext(ctx context.Context, query string) (nap.Stmt, error) {
//	return db.Conn.PrepareContext(ctx, query)
//}
func (db *Database) Exec(query string, args ...any) (sql.Result, error) {
	return db.Conn.Exec(query, args...)
}
func (db *Database) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.Conn.ExecContext(ctx, query, args...)
}
func (db *Database) Query(query string, args ...any) (*sql.Rows, error) {
	return db.Conn.Query(query, args...)
}
func (db *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return db.Conn.QueryContext(ctx, query, args...)
}
func (db *Database) QueryRow(query string, args ...any) *sql.Row {
	return db.Conn.QueryRow(query, args...)
}
func (db *Database) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return db.Conn.QueryRowContext(ctx, query, args...)
}

const txKey = "tx"

func (db *Database) TxSession(ctx context.Context) (context.Context, func(error), error) {
	if value := ctx.Value(txKey); value != nil {
		if tx, ok := value.(*sql.Tx); ok {
			return ctx, save(tx), nil
		}
	}
	tx, err := db.Conn.BeginTx(ctx, nil)
	if err != nil {
		return ctx, save(nil), err
	}
	ctx = context.WithValue(ctx, txKey, tx)

	return ctx, save(tx), nil
}

func (db *Database) GetTxSession(ctx context.Context) Session {
	if value := ctx.Value(txKey); value != nil {
		if v, ok := value.(*sql.Tx); ok {
			return v
		}
	}

	return db
}

// Use function after all transactions
func save(tx *sql.Tx) func(err error) {
	return func(err error) {
		if tx == nil {
			return
		}

		switch err {
		case nil:
			tx.Commit()
		default:
			tx.Rollback()
		}
	}
}

type Session interface {
	//Prepare(query string) (nap.Stmt, error)
	//PrepareContext(ctx context.Context, query string) (nap.Stmt, error)
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}
