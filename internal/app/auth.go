package app

import (
	e "github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/model"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (a *App) SignInHandler(ctx *gin.Context) {
	var (
		usr *model.User
		err error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&usr)
	if err != nil {
		return
	}
	usr, token, err := a.service.Authorization.SignIn(ctx, usr)
	if err != nil {
		return
	}

	a.SendResponse(ctx, gin.H{
		"token": token,
		"user":  usr,
	})
}

func (a *App) SignUpHandler(ctx *gin.Context) {
	var (
		usr *model.User
		err error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&usr)
	if err != nil {
		return
	}

	token, err := a.service.Authorization.SignUp(ctx, usr)
	if err != nil {
		return
	}

	a.SendResponse(ctx, gin.H{
		"token": token,
		"user":  usr,
	})
}

func (a *App) ValidateTokenMiddleware(ctx *gin.Context) {
	var (
		err error
	)
	defer func() {
		a.SendError(ctx, err)
		ctx.Abort()
	}()

	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		err = e.Unauthorized("Authorization header is empty")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err = e.Unauthorized("Authorization header is invalid")
		return
	}
	//parse token
	userId, err := a.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		return
	}

	ctx.Set("userId", userId)
}
