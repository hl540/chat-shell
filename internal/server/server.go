package server

import (
	"context"
	"time"

	"github.com/hl540/chat-shell/src/proto"
)

type Server struct {
	proto.UnimplementedChatServerServer
}

var i = 0

func (s *Server) Chat(stream proto.ChatServer_ChatServer) error {
	if i == 0 {
		i++
		time.Sleep(1000 * time.Second)
	}
	client := NewClient(context.Background(), stream)
	go client.Recv()
	go client.Send()
	<-client.ctx.Done()
	return nil
}

func (s *Server) Users(ctx context.Context, request *proto.BaseRequest) (*proto.UsersReply, error) {
	users := make([]string, 0)
	for _, client := range defaultPool {
		if client.id != "" {
			users = append(users, client.id)
		}
	}
	return &proto.UsersReply{Data: users}, nil
}

func (s *Server) Channels(ctx context.Context, request *proto.BaseRequest) (*proto.ChannelsReply, error) {
	channels := make([]string, 0)
	for key := range channelsPool {
		channels = append(channels, key)
	}
	return &proto.ChannelsReply{Data: channels}, nil
}
