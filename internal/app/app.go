package app

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/cheeeasy2501/book-library/internal/database"
	e "github.com/cheeeasy2501/book-library/internal/errors"
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
	//authController *auth.Authorization
	//bookController *book.BookController
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
	//TODO refactoring
	repos := repository.NewRepository(database.Instance.Conn)
	services := service.NewService(repos)
	//authorizationController := auth.NewAuthorization(cnf.Auth)
	//bookController := book.NewBookController()

	application := &App{
		ctx:    ctx,
		cnf:    cnf,
		engine: engine,
		logger: logger,
		//authController: authorizationController,
		//bookController: bookController,
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
	case e.ValidateError:
		code, message = http.StatusBadRequest, value.Error()
	case e.NotFoundError:
		code, message = http.StatusNotFound, value.Error()
	case e.Unauthorized:
		code, message = http.StatusUnauthorized, value.Error()
	default:
		code, message = http.StatusInternalServerError, value.Error()
	}

	ctx.JSON(code, HTTPError{Message: message})
}
