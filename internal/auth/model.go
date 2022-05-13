package auth

import (
	"cheeeasy2501/book-library/internal/user"
	"context"
	"crypto/sha1"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type Authorization struct {
}

func (auth *Authorization) SingIn(context context.Context, usr *user.User) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(usr.Password))
	usr.Password = string(pwd.Sum(nil))
	usr, err := repo.Get(context, usr.Username, usr.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Username: usr.Username,
	})

	return token.SignedString("key")
}
