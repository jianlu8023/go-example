package main

import (
	"context"

	"github.com/go-redis/redis/v8"

	"github.com/jianlu8023/go-example/pkg/logger"
)

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "123456",
	})
	ping := rds.Ping(context.TODO())
	logger.GetAPPLogger().Infof("ping result: %v", ping)
}
