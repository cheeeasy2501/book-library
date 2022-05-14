package errors

type HandlerValidateError string
type HandlerNotFoundError string

func (e HandlerValidateError) Error() string {
	return string(e)
}

func (e HandlerNotFoundError) Error() string {
	return string(e)
}
