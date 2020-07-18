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

func Route() {
	engine.GET("/server_time", GetServerTime)
	engine.GET("/serve_pprof", ServePprof)
	engine.GET("/ws", Ws)
}
