package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/hl540/chat-shell/src/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name1 = flag.String("name1", "", "name1")
var name2 = flag.String("name2", "", "name2")

func main() {
	flag.Parse()
	conn, err := grpc.Dial(
		"0.0.0.0:9999",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err.Error())
	}
	client := proto.NewChatServerClient(conn)
	stream, err := client.Chat(context.Background())

	go func() {
		for {
			time.Sleep(time.Second)
			stream.Send(&proto.Message{
				Type:       proto.MessageType_MessageType_Text,
				Context:    "你好",
				From:       *name1,
				Target:     *name2,
				TargetType: proto.TargetType_TargetType_User,
			})
		}
	}()

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(msg)
	}
}
