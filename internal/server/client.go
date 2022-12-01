package server

import (
	"context"
	"fmt"
	"log"

	"github.com/hl540/chat-shell/src/proto"
)

var defaultPool map[string]*Client
var channelsPool map[string]map[string]*Client

func init() {
	defaultPool = make(map[string]*Client)
	channelsPool = make(map[string]map[string]*Client)
	channelsPool["channel1"] = make(map[string]*Client)
	channelsPool["channel2"] = make(map[string]*Client)
	channelsPool["channel3"] = make(map[string]*Client)
	channelsPool["channel4"] = make(map[string]*Client)
	channelsPool["channel5"] = make(map[string]*Client)
}

func NewClient(ctx context.Context, stream proto.ChatServer_ChatServer) *Client {
	ctx, cancel := context.WithCancel(ctx)
	return &Client{
		id:       "",
		ctx:      ctx,
		cancel:   cancel,
		conn:     stream,
		SendChan: make(chan *proto.Message),
	}
}

type Client struct {
	id       string
	ctx      context.Context
	conn     proto.ChatServer_ChatServer
	SendChan chan *proto.Message
	cancel   context.CancelFunc
}

func (c *Client) Send() {
	defer close(c.SendChan)
	for message := range c.SendChan {
		select {
		case <-c.ctx.Done():
			log.Printf("调整sendchan")
			return
		default:
			c.conn.Send(message)
		}
	}
}

func (c *Client) Recv() {
	for {
		message, err := c.conn.Recv()
		if err != nil {
			log.Printf("接收消息发生错误：%s\n", err.Error())
			// 关闭客户端
			c.Close()
			return
		}
		// 更新客户端id并且添加到映射
		c.id = message.From
		defaultPool[c.id] = c

		// 处理消息
		switch message.TargetType {
		case proto.TargetType_TargetType_User:
			c.userMessageHandler(message)
		case proto.TargetType_TargetType_Channel:
			c.channelMessageHandler(message)
		default:
			log.Printf("消息目标类型错误：%s\n", message.TargetType)
		}
	}
}

func (c *Client) Close() {
	c.cancel()
	delete(defaultPool, c.id)
	for _, channel := range channelsPool {
		delete(channel, c.id)
	}
}

func (c *Client) userMessageHandler(message *proto.Message) {
	target, ok := defaultPool[message.Target]
	if !ok {
		msg := fmt.Sprintf("[%s]已下线\n", message.Target)
		c.SendChan <- &proto.Message{
			Type:       proto.MessageType_MessageType_Text,
			Context:    msg,
			From:       "system",
			Target:     message.From,
			TargetType: proto.TargetType_TargetType_User,
		}
		return
	}
	target.SendChan <- message
}

func (c *Client) channelMessageHandler(message *proto.Message) {
	channel, ok := channelsPool[message.Target]
	if !ok {
		msg := fmt.Sprintf("频道[%s]不存在\n", message.Target)
		c.SendChan <- &proto.Message{
			Type:       proto.MessageType_MessageType_Text,
			Context:    msg,
			From:       "system",
			Target:     message.From,
			TargetType: proto.TargetType_TargetType_User,
		}
		return
	}
	for _, target := range channel {
		target.SendChan <- message
	}
}
