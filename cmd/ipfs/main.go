package main

import (
	sh "github.com/ipfs/go-ipfs-api"

	"github.com/jianlu8023/go-example/pkg/logger"
)

var appLogger = logger.GetAPPLogger()

func main() {
	shell := sh.NewShell("127.0.0.1:5001")
	version, sha, err := shell.Version()
	if err != nil {
		appLogger.Errorf("获取ipfs的版本信息失败: %s", err)
		return
	}
	appLogger.Infof("获取ipfs的版本信息成功: version:%s, commit:%s", version, sha)
}
