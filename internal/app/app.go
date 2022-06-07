package app

import (
	"context"
	"fmt"
	_ "github.com/cheeeasy2501/book-library/docs/app"
	"github.com/cheeeasy2501/book-library/internal/app/apperrors"
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/cheeeasy2501/book-library/internal/database"
	"github.com/cheeeasy2501/book-library/internal/repository"
	"github.com/cheeeasy2501/book-library/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

// @title           Swagger Book Library API
// @version         1.0
// @description     Book Library Microservice
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
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
		routes.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		routes.GET("swagger", func(context *gin.Context) {
			context.Redirect(http.StatusPermanentRedirect, "api/v1/swagger/index.html")
		})
		fmt.Println(routes.BasePath())
		routes.POST("signIn", application.SignInHandler)
		routes.POST("signUp", application.SignUpHandler)
		books := routes.Group("books", application.ValidateTokenMiddleware)
		{
			books.GET("/", application.GetBooks)
			books.GET("/:id", application.GetBook)
			books.POST("/", application.CreateBook)
			books.PATCH("/:id", application.UpdateBook)
			books.DELETE("/:id", application.DeleteBook)
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
		code, message = http.StatusBadRequest, value.Error()
	}

	ctx.JSON(code, HTTPError{Message: message})
	ctx.Abort()
}
