package auth

import (
	"context"
	"database/sql"
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (auth *Authorization) GenerateToken(usr *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserName:  usr.UserName,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
	})

	return token.SignedString([]byte("key"))
}
func (auth *Authorization) SingIn(ctx context.Context, usr *user.User) (string, error) {
	encryptedPass, err := HashPassword(usr.Password)
	if err != nil {
		return "", err
	}
	usr.Password = encryptedPass
	usr, err = auth.UserRepo.CheckSignIn(ctx, usr)
	if err != nil {
		return "", err
	}

	return auth.GenerateToken(usr)
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
