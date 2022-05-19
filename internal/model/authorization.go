package model

import (
	"github.com/cheeeasy2501/book-library/internal/config"
)

type Authorization struct {
	Config *config.AuthConfig
}

func NewAuthorization(cnf *config.AuthConfig) *Authorization {
	return &Authorization{
		Config: cnf,
	}
}
