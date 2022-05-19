package apperrors

const (
	UserNotFound = NotFoundError("User isn't found")
)

const (
	UserIsNotAuthorize = Unauthorized("Unauthorized user!")
)

const (
	EmptyUserName = ValidateError("Username is empty!")
)
