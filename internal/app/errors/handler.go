package errors

type ValidateError string
type NotFoundError string

func (e ValidateError) Error() string {
	return string(e)
}

func (e NotFoundError) Error() string {
	return string(e)
}
