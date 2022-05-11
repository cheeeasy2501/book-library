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
	config := cnf.NewConfig()
	engine := gin.Default()
	logger := logrus.New()

	application := app.NewApp(appContext, config, engine, logger)
	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		err := application.StartHTTP()
		if err != nil {
			cancel()
		}
	}()

	logger.Infof("Ending app...\n")
	<-wait
	logger.Infof("Ended...\n")
}
