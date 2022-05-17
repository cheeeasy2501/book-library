package apperrors

type Unauthorized string

func (e Unauthorized) Error() string {
	return string(e)
}
