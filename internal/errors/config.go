package errors

import "fmt"

const (
	ConfigDatabaseNotFound = ConfigDatabaseError("Item in database config map not found!")
)

type ConfigDatabaseError string

func (err ConfigDatabaseError) Error() string {
	return fmt.Sprintf("Database package error - %s", err)
}
