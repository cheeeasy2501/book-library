package service

import (
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserId int64 `json:"user_id"`
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
			return nil, apperrors.InvalidSigningMethod
		}
		return []byte(auth.secretKey), nil
	})
	if err != nil {
		return 0, apperrors.Unauthorized(err.Error())
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return 0, err
	}

	return claims.UserId, nil
}

func (auth *AuthorizationService) HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return bytes, err
}

func (auth *AuthorizationService) VerifyPassword(userPass, credentialsPass string) error {
	comparePass, err := base64.StdEncoding.DecodeString(userPass)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(comparePass, []byte(credentialsPass))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return apperrors.InvalidCredentionals
		}

		return err
	}

	return nil
}

func (auth *AuthorizationService) SignIn(ctx context.Context, credentials *forms.Credentials) (*model.User, string, error) {
	user, err := auth.repo.FindByUserName(ctx, credentials.UserName)
	if err != nil {
		return nil, "", err
	}
	err = auth.VerifyPassword(user.Password(), credentials.Password)
	if err != nil {
		return nil, "", err
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (auth *AuthorizationService) SignUp(ctx context.Context, userForm *forms.CreateUser) (*model.User, string, error) {
	var user = &model.User{}
	userForm.ToUserModel(user)
	hashPassword, err := auth.HashPassword(user.Password())
	if err != nil {
		return nil, "", err
	}
	encryptedPass := base64.StdEncoding.EncodeToString(hashPassword)
	user.SetPassword(encryptedPass)
	_, err = auth.repo.FindByUserName(ctx, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return nil, "", err
	}

	err = auth.repo.Create(ctx, user)
	if err != nil {
		return nil, "", err
	}

	token, err := auth.GenerateToken(user)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}
