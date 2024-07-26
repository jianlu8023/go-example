package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	glog "github.com/jianlu8023/go-logger"
)

func main() {

	logger := glog.NewSugaredLogger(
		&glog.Config{
			LogLevel:    "DEBUG",
			DevelopMode: true,
			Caller:      true,
			ModuleName:  "[LOGGER]",
			StackLevel:  "ERROR",
		},
		glog.WithConsoleFormat(),
	)

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {

			logger.Debug("debug log")
			logger.Info("info log")
			logger.Warn("warn log")
			logger.Error("error log")
			// logger.Fatal("fatal log")
		}
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-exit:
		ticker.Stop()
		logger.Infof("server stopped ...")
	}
}
