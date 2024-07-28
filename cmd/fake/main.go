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
		ModuleName:  "[FAKE]",
		StackLevel:  "ERROR",
		Caller:      true,
	},
	glog.WithConsoleFormat(),
)

func main() {

	gofakeit.New(time.Now().UnixMilli())
	car := gofakeit.Car()
	prettyJSON, _ := json.PrettyJSON(car)
	log.Infof("fake car :\n %v", prettyJSON)
	log.Infof("fake name : %v", gofakeit.Name())
	second := gofakeit.Second()
	log.Infof("second: %d", second)
}
