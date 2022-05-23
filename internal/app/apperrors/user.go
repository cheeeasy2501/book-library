package apperrors

const (
	UserNotFound = NotFoundError("User isn't found")
)

const (
	UserIsNotAuthorize   = Unauthorized("Unauthorized user!")
	InvalidCredentionals = Unauthorized("Invalid username or password!")
)

const (
	EmptyUserName = ValidateError("Username is empty!")
)
