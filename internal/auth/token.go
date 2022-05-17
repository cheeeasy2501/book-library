package auth

import (
	e "github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"userId"`
}

func (auth *Authorization) GenerateToken(usr *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: usr.Id,
	})
	return token.SignedString([]byte(auth.AuthConfig.GetJWTKey()))
}

func (auth *Authorization) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.Unauthorized("Invalid signing method")
		}
		return []byte(auth.AuthConfig.GetJWTKey()), nil
	})
	if err != nil {
		return 0, e.Unauthorized(err.Error())
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, err
	}

	return claims.UserId, nil
}
