package app

import (
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/forms"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (a *App) SignInHandler(ctx *gin.Context) {
	var (
		err         error
		credentials *forms.Credentials
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.ShouldBindJSON(&credentials)
	if err != nil {
		return
	}
	user, token, err := a.service.Authorization.SignIn(ctx, credentials)
	if err != nil {
		return
	}

	a.SendResponse(ctx, gin.H{
		"token": token,
		"user":  user,
	})
}

func (a *App) SignUpHandler(ctx *gin.Context) {
	var (
		form *forms.CreateUser
		err  error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.ShouldBindJSON(&form)
	if err != nil {
		return
	}

	user, token, err := a.service.Authorization.SignUp(ctx, form)
	if err != nil {
		return
	}

	a.SendResponse(ctx, gin.H{
		"token": token,
		"user":  user,
	})
}

func (a *App) ValidateTokenMiddleware(ctx *gin.Context) {
	var (
		err error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		err = apperrors.EmptyAuthorizationHeader
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err = apperrors.InvalidAuthorizationHeader
		return
	}
	//parse token
	userId, err := a.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		return
	}

	ctx.Set("userId", userId)
}
