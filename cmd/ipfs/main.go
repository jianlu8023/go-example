package main

import (
	"fmt"

	sh "github.com/ipfs/go-ipfs-api"
)

func main() {
	shell := sh.NewShell("127.0.0.1:5001")
	version, sha, err := shell.Version()
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf("Version: %s\nCommit: %s\n", version, sha)
}
