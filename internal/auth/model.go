package auth

import (
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/cheeeasy2501/book-library/internal/user"
)

type Authorization struct {
	UserRepo   *user.UserRepo
	AuthConfig *config.AuthConfig
}

func NewAuthorization(cnf *config.AuthConfig) *Authorization {
	return &Authorization{
		UserRepo:   &user.UserRepo{},
		AuthConfig: cnf,
	}
}
