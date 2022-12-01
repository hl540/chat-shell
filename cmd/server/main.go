package main

import (
	"log"
	"net"

	"github.com/hl540/chat-shell/internal/server"
	"github.com/hl540/chat-shell/src/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		panic(err.Error())
	}
	grpcServer := grpc.NewServer()
	proto.RegisterChatServerServer(grpcServer, new(server.Server))

	log.Printf("服务启动，%s", lis.Addr())
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
