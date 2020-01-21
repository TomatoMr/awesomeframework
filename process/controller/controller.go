package controller

import (
	"net/http"
	"time"
)

type ServerTimeReq struct {
}

type ServerTimeResp struct {
	ServerTime int64 `json:"server_time"`
}

func GetServerTime() (ServerTimeResp, int, string) {
	resp := ServerTimeResp{}
	resp.ServerTime = time.Now().Unix()
	return resp, http.StatusOK, ""
}
