package rpc

import (
	"context"
	"github.com/TomatoMr/awesomeframework/process/controller"
	"github.com/TomatoMr/awesomeframework/process/rpc/server"
	"net"
)
import "google.golang.org/grpc"

type Server struct {
}

func StartRpcServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {

	}
	s := grpc.NewServer() //起一个服务
	server.RegisterServerServer(s, &Server{})
	if err := s.Serve(lis); err != nil {

	}
}

func (rp *Server) GetServerTime(ctx context.Context, request *server.ServerTimeRequest) (*server.ServerTimeResponse, error) {
	data, code, msg := controller.GetServerTime()
	resp := &server.ServerTimeResponse{}
	respData := &server.ServerTimeResponseData{}
	resp.Msg = msg
	resp.Code = uint32(code)
	respData.ServerTime = uint64(data.ServerTime)
	resp.Data = respData
	return resp, nil
}
