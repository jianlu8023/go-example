package main

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/jianlu8023/go-tools/pkg/format/json"

	"github.com/jianlu8023/go-example/pkg/logger"
)

var log = logger.GetSDKLogger()

func main() {

	gofakeit.New(time.Now().UnixMilli())
	car := gofakeit.Car()
	prettyJSON, _ := json.PrettyJSON(car)
	log.Infof("fake car :\n %v", prettyJSON)
	log.Infof("fake name : %v", gofakeit.Name())
	second := gofakeit.Second()
	log.Infof("second: %d", second)
}
