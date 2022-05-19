package apperrors

import "fmt"

type ConfigDatabaseError string

func (err ConfigDatabaseError) Error() string {
	return fmt.Sprintf("Database package error - %s", err)
}
