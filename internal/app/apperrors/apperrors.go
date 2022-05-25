package apperrors

import "fmt"

type (
	ValidateError       string
	NotFoundError       string
	ConfigDatabaseError string
	Unauthorized        string
)

func (e ValidateError) Error() string {
	return string(e)
}

func (e NotFoundError) Error() string {
	return string(e)
}

func (e ConfigDatabaseError) Error() string {
	return fmt.Sprintf("Database package error - %s", e)
}

func (e Unauthorized) Error() string {
	return string(e)
}
