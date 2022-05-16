package app

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/app/errors"
	"github.com/cheeeasy2501/book-library/internal/auth"
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	ctx            context.Context
	cnf            *config.Config
	engine         *gin.Engine
	logger         *logrus.Logger
	authController *auth.Authorization
}

type HTTPError struct {
	Message string `json:"message"`
}

func NewApp(ctx context.Context, cnf *config.Config, logger *logrus.Logger) (*App, error) {
	database.SetNewDatabaseInstance()
	err := database.Instance.OpenConnection(cnf)
	if err != nil {
		return nil, err
	}

	engine := gin.Default()
	authorizationController := auth.NewAuthorization()

	application := &App{
		ctx:            ctx,
		cnf:            cnf,
		engine:         engine,
		logger:         logger,
		authController: authorizationController,
	}

	routes := engine.Group("api/v1/")
	{
		routes.POST("signIn", application.SignInHandler)
		routes.POST("signUp", application.SignUpHandler)
		books := routes.Group("books", auth.TokenMiddleware())
		{
			books.GET("/", application.GetBooks)
			books.GET("/:id", application.GetBook)
			books.POST("/", application.CreateBook)
			books.PATCH("/", application.UpdateBook)
			books.DELETE("/", application.DeleteBook)
		}
	}

	return application, nil
}

func (a App) StartHTTP() error {
	err := a.engine.Run(":" + a.cnf.Api.Port)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) SendResponse(ctx *gin.Context, body interface{}) {
	ctx.JSON(http.StatusOK, body)
}

func (a *App) SendError(ctx *gin.Context, err error) {
	if err == nil {
		return
	}

	var (
		code    int
		message string
	)

	switch value := err.(type) {
	case errors.ValidateError:
		code, message = http.StatusBadRequest, value.Error()
	case errors.NotFoundError:
		code, message = http.StatusNotFound, value.Error()
	case errors.Unauthorized:
		code, message = http.StatusUnauthorized, value.Error()
	default:
		code, message = http.StatusInternalServerError, value.Error()
	}

	ctx.JSON(code, HTTPError{Message: message})
}
