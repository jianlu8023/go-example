package main

import (
	"github.com/jianlu8023/go-example/pkg/crypto/sm4"
	"github.com/jianlu8023/go-example/pkg/logger"
)

var log = logger.GetAppLogger()

func main() {

	key := "0123456789abcdef"
	original := "testdata/sm4/sm4.txt"
	encrypt := "testdata/sm4/sm4-enc.txt"
	decrypt := "testdata/sm4/sm4-dec.txt"
	err := sm4.EncryptFile(key, original, encrypt)
	if err != nil {
		log.Errorf("sm4 encrypt file failed: %v", err)
		return
	}

	err = sm4.DecryptFile(key, encrypt, decrypt)
	if err != nil {
		log.Errorf("sm4 decrypt file failed: %v", err)
		return
	}
}
