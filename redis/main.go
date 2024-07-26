package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	rds := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "123456",
	})
	ping := rds.Ping(context.TODO())
	fmt.Println(ping)
}
