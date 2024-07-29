package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/jianlu8023/go-example/pkg/logger"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/basic"), &gorm.Config{
		PrepareStmt:     true,
		Logger:          logger.GetDBLogger(),
		CreateBatchSize: 1000,
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	var version string
	err = db.Raw("select version()").Row().Scan(&version)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(version)
}
