package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	glog "github.com/jianlu8023/go-logger"
	"github.com/jianlu8023/go-tools/pkg/format/json"
)

var log = glog.NewSugaredLogger(
	&glog.Config{
		LogLevel:    "DEBUG",
		DevelopMode: true,
	},
	glog.WithConsoleFormat(),
	glog.WithLumberjack(
		&glog.LumberjackConfig{
			FileName:   "./logs/lumberjack-fake.log",
			MaxBackups: 365,
			MaxSize:    5,
			MaxAge:     30,
			Localtime:  true,
			Compress:   true,
		},
	),
	glog.WithRotateLog(
		&glog.RotateLogConfig{
			FileName:     "./logs/rotatelogs-fake.log",
			LocalTime:    true,
			MaxAge:       "365",
			RotationTime: "1h",
		},
	),
)

func main() {

	gofakeit.New(time.Now().UnixMilli())
	car := gofakeit.Car()
	prettyJSON, _ := json.ToJSON(car)
	log.Infof("fake car : %v", prettyJSON)
	log.Infof("fake name : %v", gofakeit.Name())
	second := gofakeit.Second()
	log.Infof("second: %d", second)
}
