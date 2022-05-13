package database

import "fmt"

const (
	DatabaseConfigNotFound = DatabaseError("Item in database config map not found!")
)

type DatabaseError string

func (err DatabaseError) Error() string {
	return fmt.Sprintf("Database package error - %s", err)
}
