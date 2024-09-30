package main

import (
	"fmt"
	"log"

	"github.com/jianlu8023/go-example/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/basic"), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.GetDBLogger(),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	var version string
	db.Raw("SELECT VERSION()").Row().Scan(&version)
	fmt.Println("数据库版本:", version)
}
