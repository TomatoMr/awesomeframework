package http

import (
	"github.com/TomatoMr/awesomeframework/process/controller"
	"github.com/gin-gonic/gin"
)

type Request struct {
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GetServerTime(ctx *gin.Context) {
	resp := Response{}
	resp.Data, resp.Code, resp.Msg = controller.GetServerTime()
	ctx.JSON(resp.Code, resp)
}
