package main

import (
	"context"
	"github.com/TomatoMr/awesomeframework/process/rpc/server"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	Addr = "127.0.0.1:8081"
)

func main() {
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := server.NewServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.GetServerTime(ctx, &server.ServerTimeRequest{})

	log.Printf("Data: %v", r.Data.ServerTime)
}
