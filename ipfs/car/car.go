package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ipfs/boxo/gateway"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/jianlu8023/go-tools/pkg/path"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	backend, err := gateway.NewRemoteCarBackend([]string{"http://172.25.138.45:8080"}, &client)
	if err != nil {
		fmt.Println("NewRemoteCarBackend Error:", err)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	c, err := cid.Parse("bafybeigggyffcf6yfhx5irtwzx3cgnk6n3dwylkvcpckzhqqrigsxowjwe")
	if err != nil {
		fmt.Println("c.Parse Error:", err)
		return
	}

	immutablePath, err := path.NewImmutablePath(path.FromCid(c))
	if err != nil {
		fmt.Println("p.NewImmutablePath Error:", err)
		return
	}

	car, closer, err := backend.GetCAR(ctx, immutablePath, gateway.CarParams{
		Scope: gateway.DagScopeAll,
	})

	if err != nil {
		fmt.Println("GetCAR Error:", err)
		return
	}
	fmt.Println("car:", car)
	fmt.Println("closer:", closer)
	defer func(closer io.ReadCloser) {
		err := closer.Close()
		if err != nil {
			fmt.Println("closer.Close Error:", err)
		}
	}(closer)
	p := "./testdata/test.car"

	err = path.Ensure(p)

	if err != nil {
		fmt.Println("path.Ensure Error:", err)
		return
	}

	create, err := os.Create(p)
	if err != nil {
		fmt.Println("os.Create Error:", err)
		return
	}

	_, err = io.Copy(create, closer)
	if err != nil {
		fmt.Println("io.Copy Error:", err)
		return
	}
}
