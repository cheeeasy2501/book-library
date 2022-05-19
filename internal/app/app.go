package app

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"github.com/cheeeasy2501/book-library/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type App struct {
	ctx     context.Context
	cnf     *config.Config
	engine  *gin.Engine
	logger  *logrus.Logger
	service *service.Service
}

type HTTPError struct {
	Message string `json:"message"`
}

func NewApp(ctx context.Context, cnf *config.Config, logger *logrus.Logger) (*App, error) {
	// create and open new connection
	connection, err := database.NewDatabaseConnection(cnf.Database)
	if err != nil {
		return nil, err
	}
	engine := gin.Default()
	repos := repository.NewRepository(connection)
	services := service.NewService(repos)

	application := &App{
		ctx:     ctx,
		cnf:     cnf,
		engine:  engine,
		logger:  logger,
		service: services,
	}

	routes := engine.Group("api/v1/")
	{
		routes.POST("signIn", application.SignInHandler)
		routes.POST("signUp", application.SignUpHandler)
		books := routes.Group("books", application.ValidateTokenMiddleware)
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
	case apperrors.ValidateError:
		code, message = http.StatusBadRequest, value.Error()
	case apperrors.NotFoundError:
		code, message = http.StatusNotFound, value.Error()
	case apperrors.Unauthorized:
		code, message = http.StatusUnauthorized, value.Error()
	default:
		code, message = http.StatusInternalServerError, value.Error()
	}

	ctx.JSON(code, HTTPError{Message: message})
}
