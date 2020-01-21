package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

var engine *gin.Engine

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
}
