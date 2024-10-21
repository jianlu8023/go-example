package logger

import (
	"sync"
	"time"

	glog "github.com/jianlu8023/go-logger"
	"github.com/jianlu8023/go-logger/dblogger"
	"go.uber.org/zap"
)

var (
	appLogger *zap.SugaredLogger
	sdkLogger *zap.SugaredLogger
	dbLogger  *dblogger.Logger
)

var once sync.Once

func init() {

	once.Do(func() {
		appLogger = glog.NewSugaredLogger(
			&glog.Config{
				LogLevel:    "DEBUG",
				DevelopMode: true,
				StackLevel:  "ERROR",
				ModuleName:  "APP",
				Caller:      true,
			},
			glog.WithConsoleFormat(),
			glog.WithRotateLog(glog.RotateLogDefaultConfig()),
			glog.WithRotateLog(
				&glog.RotateLogConfig{
					FileName:     "./logs/rotatelogs-app.log",
					LocalTime:    true,
					MaxAge:       "365",
					RotationTime: "1h",
				},
			),
			glog.WithLumberjack(glog.LumberjackDefaultConfig()),
			glog.WithLumberjack(
				&glog.LumberjackConfig{
					FileName:   "./logs/lumberjack-app.log",
					Localtime:  true,
					MaxAge:     30,
					MaxBackups: 10,
					MaxSize:    5,
					Compress:   false,
				},
			),
		)

		sdkLogger = glog.NewSugaredLogger(
			&glog.Config{
				LogLevel:    "DEBUG",
				DevelopMode: true,
				StackLevel:  "ERROR",
				ModuleName:  "SDK",
				Caller:      true,
			},
			glog.WithConsoleFormat(),
			glog.WithRotateLog(glog.RotateLogDefaultConfig()),
			glog.WithRotateLog(
				&glog.RotateLogConfig{
					FileName:     "./logs/rotatelogs-sdk.log",
					LocalTime:    true,
					MaxAge:       "365",
					RotationTime: "1h",
				},
			),
			glog.WithLumberjack(glog.LumberjackDefaultConfig()),
			glog.WithLumberjack(
				&glog.LumberjackConfig{
					FileName:   "./logs/lumberjack-sdk.log",
					Localtime:  true,
					MaxAge:     30,
					MaxBackups: 10,
					MaxSize:    5,
					Compress:   false,
				},
			),
		)

		dbLogger = dblogger.NewDBLogger(
			dblogger.Config{
				Logger: glog.NewLogger(
					&glog.Config{
						LogLevel:    "DEBUG",
						DevelopMode: true,
						StackLevel:  "ERROR",
						ModuleName:  "DB",
						Caller:      true,
					},
					glog.WithConsoleFormat(),
					glog.WithRotateLog(glog.RotateLogDefaultConfig()),
					glog.WithRotateLog(
						&glog.RotateLogConfig{
							FileName:     "./logs/rotatelogs-db.log",
							LocalTime:    true,
							MaxAge:       "365",
							RotationTime: "1h",
						},
					),
					glog.WithLumberjack(glog.LumberjackDefaultConfig()),
					glog.WithLumberjack(
						&glog.LumberjackConfig{
							FileName:   "./logs/lumberjack-db.log",
							Localtime:  true,
							MaxAge:     30,
							MaxBackups: 10,
							MaxSize:    5,
							Compress:   false,
						},
					),
				),
				LogLevel:                  dblogger.INFO,
				SlowThreshold:             2 * time.Second,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
				Colorful:                  true,
				ShowSql:                   true,
			},
		)
	})
}

func GetAPPLogger() *zap.SugaredLogger { return appLogger }

func GetDBLogger() *dblogger.Logger {
	return dbLogger
}

func GetSDKLogger() *zap.SugaredLogger { return sdkLogger }
