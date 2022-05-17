package main

import (
	"context"
	"github.com/cheeeasy2501/book-library/internal/app"
	"github.com/cheeeasy2501/book-library/internal/config"
	"github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

func main() {
	wait := make(chan struct{}, 0)
	backgroundContext := context.Background()
	appContext, cancel := signal.NotifyContext(backgroundContext, syscall.SIGTERM, syscall.SIGINT)
	logger := logrus.New()
	cnf := config.NewConfig()
	err := cnf.LoadEnv()
	if err != nil {
		logger.Errorf("Env is not loaded in config - %s", err)
		wait <- struct{}{}
		return
	}

	application, err := app.NewApp(appContext, cnf, logger)
	if err != nil {
		logger.Errorf("App not started - %s", err.Error())
		wait <- struct{}{}
		return
	}

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
	logger.Infof("Stop application!\n")
	logger.Infof("Application stoped!\n")
}
