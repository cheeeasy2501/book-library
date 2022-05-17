package service

import (
	"context"
	"database/sql"
	e "github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"userId"`
}

type AuthorizationService struct {
	repo      repository.UserRepoInterface
	secretKey string
}

func NewAuthorizationService(repo repository.UserRepoInterface, secretKey string) *AuthorizationService {
	return &AuthorizationService{
		repo:      repo,
		secretKey: secretKey,
	}
}

func (auth *AuthorizationService) GenerateToken(usr *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: usr.Id,
	})
	return token.SignedString([]byte(auth.secretKey))
}

func (auth *AuthorizationService) ParseToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.Unauthorized("Invalid signing method")
		}
		return []byte(auth.secretKey), nil
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

func (auth *AuthorizationService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (auth *AuthorizationService) SignIn(ctx context.Context, usr *model.User) (*model.User, string, error) {
	usr, err := auth.repo.CheckSignIn(ctx, usr)
	if err != nil {
		return nil, "", err
	}
	token, err := auth.GenerateToken(usr)
	if err != nil {
		return nil, "", err
	}

	return usr, token, nil
}

func (auth *AuthorizationService) SignUp(ctx context.Context, usr *model.User) (string, error) {
	encryptedPass, err := auth.HashPassword(usr.Password)
	if err != nil {
		return "", err
	}
	usr.Password = encryptedPass
	_, err = auth.repo.FindByUsername(ctx, usr.UserName)

	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	err = auth.repo.Create(ctx, usr)
	if err != nil {
		return "", err
	}

	return auth.GenerateToken(usr)
}
