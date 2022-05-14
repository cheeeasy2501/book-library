package auth

import (
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Authorization struct {
	UserRepo *user.UserRepo
}

func NewAuthorization() *Authorization {
	return &Authorization{
		UserRepo: &user.UserRepo{},
	}
}
