package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/jianlu8023/go-example/internal/logger"
)

var (
	wsUpgrader = websocket.Upgrader{
		HandshakeTimeout: 10 * time.Second,
		ReadBufferSize:   1024,
		WriteBufferSize:  1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {

	router := gin.Default()

	router.GET("/ws", func(ctx *gin.Context) {

		ws, err := wsUpgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			logger.GetAppLogger().Errorf("websocket upgrade error %v", err)
			return
		}
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				logger.GetAppLogger().Errorf("websocket close error %v", err)
			}
		}(ws)

		for {
			messageType, p, err := ws.ReadMessage()
			if err != nil {
				if errors.Is(err, websocket.ErrCloseSent) {
					logger.GetAppLogger().Info("websocket连接已关闭")
					return
				}

				logger.GetAppLogger().Errorf("websocket read message error %v", err)
				break
			}
			switch messageType {
			case websocket.TextMessage:
				fmt.Printf("处理文本消息, %s\n", string(p))
				ws.WriteMessage(websocket.TextMessage, p)
				// c.Writer.Write(p)
			case websocket.BinaryMessage:
				fmt.Println("处理二进制消息")
			case websocket.CloseMessage:
				fmt.Println("关闭websocket连接")
				ws.Close()
				return
			case websocket.PingMessage:
				fmt.Println("处理ping消息")
				ws.WriteMessage(websocket.PongMessage, []byte("ping"))
			case websocket.PongMessage:
				fmt.Println("处理pong消息")
				ws.WriteMessage(websocket.PongMessage, []byte("pong"))
			default:
				fmt.Printf("未知消息类型: %d\n", messageType)
				return
			}
		}
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.GetAppLogger().Errorf("starter gin web error %v", err)
			quit <- syscall.SIGINT
		}
	}()

	<-quit
	if err := srv.Shutdown(context.Background()); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.GetAppLogger().Info("gin web服务 正常关闭")
		} else {
			logger.GetAppLogger().Errorf("gin web服务 关闭失败 %v", err)
		}
	}
}
