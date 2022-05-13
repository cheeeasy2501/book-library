package main

import (
	"cheeeasy2501/book-library/internal/app"
	cnf "cheeeasy2501/book-library/internal/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	wait := make(chan struct{}, 0)
	backgroundContext := context.Background()
	appContext, cancel := signal.NotifyContext(backgroundContext, syscall.SIGTERM, syscall.SIGINT)
	logger := logrus.New()
	config := cnf.NewConfig()
	err := config.LoadEnv()
	if err != nil {
		logger.Errorf("Env is not loaded in config - %s", err)
		wait <- struct{}{}
		return
	}
	engine := gin.Default()

	application := app.NewApp(appContext, config, engine, logger)
	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		err = application.StartHTTP()
		if err != nil {
			cancel()
		}
	}()

	<-wait
	logger.Infof("Ending app...\n")
	logger.Infof("Ended...\n")
}
