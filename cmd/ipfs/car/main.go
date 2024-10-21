package main

import (
	sh "github.com/ipfs/go-ipfs-api"

	"github.com/jianlu8023/go-example/internal/logger"
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

	err = shell.Get("bafybeif4zkmu7qdhkpf3pnhwxipylqleof7rl6ojbe7mq3fzogz6m4xk3i", "/root/go/src/github.com/jianlu8023/go-example/testdata/car")
	if err != nil {
		appLogger.Errorf("获取ipfs的cat信息失败: %s", err)
		return
	}
}
