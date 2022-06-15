package apperrors

const (
	BooksNotFound = NotFoundError("Books aren't found!")
	BookNotFound  = NotFoundError("Book isn't found!")
)

type AppError struct {
	HTTPError     error
	InternalError error
}

func NewAppError(httpErr error, internalErr error) AppError {
	return AppError{
		HTTPError:     httpErr,
		InternalError: internalErr,
	}
}

func (err AppError) Error() string {
	return err.HTTPError.Error()
}
