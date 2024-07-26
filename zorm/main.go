package main

import (
	"log"

	"gitee.com/chunanyong/zorm"
	_ "github.com/go-sql-driver/mysql"
)

// https://zorm.cn/post/doc/index.html

// https://gitee.com/chunanyong/zorm/tree/v1.7.6

func main() {
	dbConfig := &zorm.DataSourceConfig{
		DSN:        "root:123456@tcp(127.0.0.1:3306)/basic?charset=utf8&parseTime=true&loc=Local",
		DriverName: "mysql",
		Dialect:    "mysql",
	}

	dao, err := zorm.NewDBDao(dbConfig)
	if err != nil {
		log.Fatalln(err)
		return
	}

	defer func() {
		if err := dao.CloseDB(); err != nil {
			log.Fatalln(err)
		}
	}()

}
