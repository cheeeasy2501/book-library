package repository

import (
	"github.com/tsenart/nap"
)

type AuthorizationRepository struct {
	db *nap.DB
}

func NewAuthorizationRepository(db *nap.DB) *AuthorizationRepository {
	return &AuthorizationRepository{
		db: db,
	}
}
