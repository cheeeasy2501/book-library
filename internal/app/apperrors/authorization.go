package apperrors

const (
	InvalidSigningMethod       = Unauthorized("Token singing method is invalid.")
	EmptyAuthorizationHeader   = Unauthorized("Authorization header is empty.")
	InvalidAuthorizationHeader = Unauthorized("Authorization header is invalid")
)
