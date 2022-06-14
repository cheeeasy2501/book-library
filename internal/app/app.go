package app

import (
	"context"
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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host      localhost:8080
// @BasePath  /api/v1
func NewApp(ctx context.Context, cnf *config.Config, logger *logrus.Logger) (*App, error) {
	db := database.NewDatabase(cnf.Database)
	err := db.OpenConnection()
	if err != nil {
		return nil, err
	}
	engine := gin.Default()
	repos := repository.NewRepository(db)
	services := service.NewService(db, repos)

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
		routes.POST("signIn", application.SignInHandler)
		routes.POST("signUp", application.SignUpHandler)
		secureRoutes := routes.Group("", application.ValidateTokenMiddleware)
		{
			books := secureRoutes.Group("books")
			{
				books.GET("/", application.GetBooks)
				books.GET("/:id", application.GetBook)
				books.POST("/", application.CreateBook)
				books.PATCH("/:id", application.UpdateBook)
				books.DELETE("/:id", application.DeleteBook)
			}
			authors := secureRoutes.Group("authors")
			{
				authors.GET("/", application.GetAuthors)
				authors.GET("/:id", application.GetAuthor)
				authors.POST("/", application.CreateAuthor)
				authors.PATCH("/:id", application.UpdateAuthor)
				authors.DELETE("/:id", application.DeleteAuthor)
			}

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
