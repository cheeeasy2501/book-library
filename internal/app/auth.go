package app

import (
	"github.com/cheeeasy2501/book-library/internal/user"
	"github.com/gin-gonic/gin"
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
	token, err := a.authController.SingIn(ctx, usr)
	if err != nil {
		return
	}

	a.SendResponse(ctx, gin.H{
		"token": token,
	})
}
