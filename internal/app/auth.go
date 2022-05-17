package app

import (
	e "github.com/cheeeasy2501/book-library/internal/errors"
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (a *App) SignInHandler(ctx *gin.Context) {
	var (
		usr *user.User
		err error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&usr)
	if err != nil {
		return
	}
	usr, token, err := a.authController.SingIn(ctx, usr)
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
		usr *user.User
		err error
	)
	defer func() {
		a.SendError(ctx, err)
	}()

	err = ctx.BindJSON(&usr)
	if err != nil {
		return
	}

	token, err := a.authController.SignUp(ctx, usr)
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
	userId, err := a.authController.ParseToken(headerParts[1])
	if err != nil {
		return
	}

	ctx.Set("userId", userId)
}
