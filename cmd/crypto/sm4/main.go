package main

import "github.com/jianlu8023/go-example/pkg/crypto/sm4"

func main() {

	key := "0123456789abcdef"
	sm4.EncryptFile(key, "testdata/sm4.txt",
		"testdata/sm4-enc.txt")

	sm4.DecryptFile(key,
		"testdata/sm4-enc.txt",
		"testdata/sm4-dec.txt")
}
