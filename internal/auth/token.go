package auth

import (
	"github.com/cheeeasy2501/book-library/internal/app/errors"
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

type Claims struct {
	jwt.RegisteredClaims
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (auth *Authorization) GenerateToken(usr *user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(1))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserName: usr.UserName,
	})

	return token.SignedString([]byte("key"))
}

func (s *Authorization) ValidateToken(token string) (*user.User, error) {
	// TODO: verify token implementations
	//claims, err := validateIDToken(token, s.PubKey) // uses public RSA key
	//
	//// We'll just return unauthorized error in all instances of failing to verify user
	//if err != nil {
	//	log.Printf("Unable to validate or parse idToken - Error: %v\n", err)
	//	return nil, apperrors.NewAuthorization("Unable to verify user from idToken")
	//}
	//
	//return claims.User, nil
}

type authHeader struct {
	Token string `header:"Authorization"`
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := authHeader{}
		if err := c.ShouldBindHeader(&header); err != nil {
			if _, ok := err.(validator.ValidationErrors); ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err,
				})
				c.Abort()
				return
			}

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": err,
				})
			}
		}

		idToken := strings.Split(header.Token, "Bearer ")

		if len(idToken) < 2 {
			err := errors.ValidateError("Must provide Authorization header with format `Bearer {token}`")

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			c.Abort()
			return
		}

		//user, err := s.ValidateIDToken(idTokenHeader[1])
	}
}
