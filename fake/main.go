package main

import (
	"github.com/brianvoe/gofakeit/v7"
	glog "github.com/jianlu8023/go-logger"
	"github.com/jianlu8023/go-tools/pkg/format/json"
)

var log = glog.NewSugaredLogger(
	&glog.Config{
		LogLevel:    "DEBUG",
		DevelopMode: true,
		Caller:      true,
		ModuleName:  "[FAKE]",
		StackLevel:  "",
	},
	glog.WithConsoleFormat(),
)

func main() {

	gofakeit.New(0)
	car := gofakeit.Car()
	prettyJSON, _ := json.PrettyJSON(car)
	log.Infof("fake car :\n %v", prettyJSON)
	log.Infof("fake name : %v", gofakeit.Name())
	second := gofakeit.Second()
	log.Infof("second: %d", second)
}
