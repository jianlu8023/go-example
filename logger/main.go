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
		},
		glog.WithConsoleFormat(),
		glog.WithLumberjack(
			&glog.LumberjackConfig{
				FileName:   "./logs/lumberjack-logger.log",
				MaxBackups: 365,
				MaxSize:    5,
				MaxAge:     30,
				Localtime:  true,
				Compress:   true,
			},
		),
		glog.WithRotateLog(
			&glog.RotateLogConfig{
				FileName:     "./logs/rotatelogs-logger.log",
				LocalTime:    true,
				MaxAge:       "365",
				RotationTime: "1h",
			},
		),
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
