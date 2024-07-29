package main

import (
	"log"

	// _ "gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"

	"github.com/jianlu8023/go-example/pkg/logger"
)

// https://gitea.com/xorm/xorm/src/tag/v1.3.6

func main() {

	// 定义logger等到core.db
	// xorm.NewEngineWithDB("", "", &core.DB{})

	engine, err := xorm.NewEngine("mysql",
		"root:123456@tcp(127.0.0.1:3306)/basic?charset=utf8mb4")
	if err != nil {
		log.Fatalln(err)
		return
	}
	engine.SetLogger(logger.GetDBLogger())
	defer func() {
		if err := engine.Close(); err != nil {
			log.Fatalln(err)
			return
		}
	}()
	err = engine.Ping()
	if err != nil {
		log.Fatalln(err)
		return
	}
	version, err := engine.DBVersion()
	if err != nil {

		log.Fatalln(err)
		return

	}
	log.Println("mysql version:", version)
}
