package http

import (
	"github.com/TomatoMr/awesomeframework/admin"
	"github.com/TomatoMr/awesomeframework/process/controller"
	"github.com/gin-gonic/gin"
	"net/http"
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

func ServePprof(ctx *gin.Context) {
	resp := Response{}
	pCh := admin.Get()
	pCh.OnOff <- struct{}{}
	resp.Code = http.StatusOK
	resp.Msg = "open Pprof success. see: http://xx/debug/pprof"
	ctx.JSON(resp.Code, resp)
}
