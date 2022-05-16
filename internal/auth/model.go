package auth

import (
	"github.com/cheeeasy2501/book-library/internal/user"
)

type Authorization struct {
	UserRepo *user.UserRepo
}

func NewAuthorization() *Authorization {
	return &Authorization{
		UserRepo: &user.UserRepo{},
	}
}
