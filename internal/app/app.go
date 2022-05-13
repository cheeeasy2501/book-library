package app

import (
	"cheeeasy2501/book-library/internal/config"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type App struct {
	context context.Context
	config  *config.Config
	engine  *gin.Engine
	logger  *logrus.Logger
}

func NewApp(context context.Context, config *config.Config, engine *gin.Engine, logger *logrus.Logger) App {
	return App{
		context,
		config,
		engine,
		logger,
	}
}

func (a App) StartHTTP() error {
	fmt.Println(a.config)
	err := a.engine.Run(":" + a.config.Api.Port)
	if err != nil {
		return err
	}

	return nil
}
