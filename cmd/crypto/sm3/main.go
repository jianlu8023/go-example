package main

import (
	"fmt"

	"github.com/jianlu8023/go-example/pkg/crypto/sm3"
)

func main() {
	hash := sm3.Hash("hello world")
	fmt.Println(string(hash))
}
