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
	glog "github.com/jianlu8023/go-logger"
	fpath "github.com/jianlu8023/go-tools/pkg/path"
)

func main() {

	logger := glog.NewSugaredLogger(
		&glog.Config{
			LogLevel:    "DEBUG",
			DevelopMode: true,
		},
		glog.WithConsoleFormat(),
		glog.WithLumberjack(
			&glog.LumberjackConfig{
				FileName:   "./logs/lumberjack-ipfs.log",
				MaxBackups: 365,
				MaxSize:    5,
				MaxAge:     30,
				Localtime:  true,
				Compress:   true,
			},
		),
		glog.WithRotateLog(
			&glog.RotateLogConfig{
				FileName:     "./logs/rotatelogs-ipfs.log",
				LocalTime:    true,
				MaxAge:       "365",
				RotationTime: "1h",
			},
		),
	)

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	backend, err := gateway.NewRemoteCarBackend([]string{"http://172.25.138.45:8080"}, &client)
	if err != nil {
		logger.Errorf("NewRemoteCarBackend Error: %v", err)
		return
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()
	c, err := cid.Parse("bafybeigggyffcf6yfhx5irtwzx3cgnk6n3dwylkvcpckzhqqrigsxowjwe")
	if err != nil {
		logger.Errorf("parse cid Error %v", err)
		return
	}

	immutablePath, err := path.NewImmutablePath(path.FromCid(c))
	if err != nil {
		logger.Errorf("p.NewImmutablePath Error: %v", err)
		return
	}

	car, closer, err := backend.GetCAR(ctx, immutablePath, gateway.CarParams{
		Scope: gateway.DagScopeAll,
	})

	if err != nil {
		logger.Errorf("GetCAR Error: %v", err)
		return
	}
	logger.Infof("car: %v, closer: %v", car, closer)
	defer func(closer io.ReadCloser) {
		err := closer.Close()
		if err != nil {
			logger.Errorf("closer.Close Error: %v", err)

		}
	}(closer)
	p := "./testdata/test.car"

	err = fpath.Ensure(p[:len(p)-8])

	if err != nil {

		logger.Errorf("path.Ensure Error: %v", err)
		return

	}

	create, err := os.Create(p)
	if err != nil {
		logger.Errorf("os.Create Error: %v", err)
		return
	}

	_, err = io.Copy(create, closer)
	if err != nil {
		logger.Errorf("io.Copy Error: %v", err)
		return
	}
}
