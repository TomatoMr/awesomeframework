package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"os"
)

var engine *gin.Engine
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartHttpServer(addr string) {
	engine = gin.Default()
	Route()

	if err := engine.Run(addr); err != nil {
		zap.Error(err)
		os.Exit(1)
	}
}

//webSocket请求ping 返回pong
func Ws(ctx *gin.Context) {
	ws, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			message = []byte("pong")
		}
		//写入数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func Route() {
	engine.GET("/server_time", GetServerTime)
	engine.GET("/ws", Ws)
}
