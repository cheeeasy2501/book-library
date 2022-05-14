package auth

import (
	"context"
	"crypto/sha1"
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func (auth *Authorization) SingIn(context context.Context, usr *user.User) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(usr.Password))
	usr.Password = string(pwd.Sum(nil))
	usr, err := auth.UserRepo.GetAuth(context, usr.UserName, usr.Password)
	if err != nil {
		return "", err
	}

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
