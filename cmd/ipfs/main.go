package main

import (
	"fmt"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {

	newShell := shell.NewShell("http://127.0.0.1:5001")
	up := newShell.IsUp()
	fmt.Println(up)

}
