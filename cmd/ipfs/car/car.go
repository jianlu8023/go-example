package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ipfs/boxo/gateway"
	"github.com/ipfs/boxo/path"
	"github.com/ipfs/go-cid"
	"github.com/jianlu8023/go-example/pkg/logger"
	fpath "github.com/jianlu8023/go-tools/pkg/path"
)

var appLogger = logger.GetAppLogger()

func main() {

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	backend, err := gateway.NewRemoteCarBackend([]string{"http://172.25.138.45:8080"}, &client)
	if err != nil {
		appLogger.Errorf("NewRemoteCarBackend Error: %v", err)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	c, err := cid.Parse("bafybeigggyffcf6yfhx5irtwzx3cgnk6n3dwylkvcpckzhqqrigsxowjwe")
	if err != nil {
		appLogger.Errorf("parse cid Error %v", err)
		return
	}

	immutablePath, err := path.NewImmutablePath(path.FromCid(c))
	if err != nil {
		appLogger.Errorf("p.NewImmutablePath Error: %v", err)
		return
	}

	car, closer, err := backend.GetCAR(ctx, immutablePath, gateway.CarParams{
		Scope: gateway.DagScopeAll,
	})

	if err != nil {
		appLogger.Errorf("GetCAR Error: %v", err)
		return
	}
	appLogger.Infof("car: %v, closer: %v", car, closer)
	defer func(closer io.ReadCloser) {
		err := closer.Close()
		if err != nil {
			appLogger.Errorf("closer.Close Error: %v", err)

		}
	}(closer)
	p := "./testdata/test.car"

	err = fpath.Ensure(p[:len(p)-8])

	if err != nil {

		appLogger.Errorf("path.Ensure Error: %v", err)
		return

	}

	create, err := os.Create(p)
	if err != nil {
		appLogger.Errorf("os.Create Error: %v", err)
		return
	}

	_, err = io.Copy(create, closer)
	if err != nil {
		appLogger.Errorf("io.Copy Error: %v", err)
		return
	}
}
