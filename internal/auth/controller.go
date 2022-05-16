package auth

import (
	"context"
	"database/sql"
	"github.com/cheeeasy2501/book-library/internal/user"
)

func (auth *Authorization) SingIn(ctx context.Context, usr *user.User) (*user.User, string, error) {
	usr, err := auth.UserRepo.CheckSignIn(ctx, usr)
	if err != nil {
		return nil, "", err
	}
	token, err := auth.GenerateToken(usr)
	if err != nil {
		return nil, "", err
	}

	return usr, token, nil
}

func (auth *Authorization) SignUp(ctx context.Context, usr *user.User) (string, error) {
	encryptedPass, err := HashPassword(usr.Password)
	if err != nil {
		return "", err
	}
	usr.Password = encryptedPass
	_, err = auth.UserRepo.FindByUsername(ctx, usr.UserName)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	err = auth.UserRepo.Create(ctx, usr)
	if err != nil {
		return "", err
	}

	return auth.GenerateToken(usr)
}
