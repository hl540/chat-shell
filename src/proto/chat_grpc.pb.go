// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServerClient is the client API for ChatServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServerClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (ChatServer_ChatClient, error)
	Users(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*UsersReply, error)
	Channels(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*ChannelsReply, error)
}

type chatServerClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServerClient(cc grpc.ClientConnInterface) ChatServerClient {
	return &chatServerClient{cc}
}

func (c *chatServerClient) Chat(ctx context.Context, opts ...grpc.CallOption) (ChatServer_ChatClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatServer_ServiceDesc.Streams[0], "/proto.ChatServer/chat", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServerChatClient{stream}
	return x, nil
}

type ChatServer_ChatClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatServerChatClient struct {
	grpc.ClientStream
}

func (x *chatServerChatClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServerChatClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServerClient) Users(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*UsersReply, error) {
	out := new(UsersReply)
	err := c.cc.Invoke(ctx, "/proto.ChatServer/users", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServerClient) Channels(ctx context.Context, in *BaseRequest, opts ...grpc.CallOption) (*ChannelsReply, error) {
	out := new(ChannelsReply)
	err := c.cc.Invoke(ctx, "/proto.ChatServer/channels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServerServer is the server API for ChatServer service.
// All implementations must embed UnimplementedChatServerServer
// for forward compatibility
type ChatServerServer interface {
	Chat(ChatServer_ChatServer) error
	Users(context.Context, *BaseRequest) (*UsersReply, error)
	Channels(context.Context, *BaseRequest) (*ChannelsReply, error)
	mustEmbedUnimplementedChatServerServer()
}

// UnimplementedChatServerServer must be embedded to have forward compatible implementations.
type UnimplementedChatServerServer struct {
}

func (UnimplementedChatServerServer) Chat(ChatServer_ChatServer) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatServerServer) Users(context.Context, *BaseRequest) (*UsersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Users not implemented")
}
func (UnimplementedChatServerServer) Channels(context.Context, *BaseRequest) (*ChannelsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Channels not implemented")
}
func (UnimplementedChatServerServer) mustEmbedUnimplementedChatServerServer() {}

// UnsafeChatServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServerServer will
// result in compilation errors.
type UnsafeChatServerServer interface {
	mustEmbedUnimplementedChatServerServer()
}

func RegisterChatServerServer(s grpc.ServiceRegistrar, srv ChatServerServer) {
	s.RegisterService(&ChatServer_ServiceDesc, srv)
}

func _ChatServer_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServerServer).Chat(&chatServerChatServer{stream})
}

type ChatServer_ChatServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatServerChatServer struct {
	grpc.ServerStream
}

func (x *chatServerChatServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServerChatServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChatServer_Users_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServerServer).Users(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatServer/users",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServerServer).Users(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatServer_Channels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServerServer).Channels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatServer/channels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServerServer).Channels(ctx, req.(*BaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatServer_ServiceDesc is the grpc.ServiceDesc for ChatServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatServer",
	HandlerType: (*ChatServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "users",
			Handler:    _ChatServer_Users_Handler,
		},
		{
			MethodName: "channels",
			Handler:    _ChatServer_Channels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "chat",
			Handler:       _ChatServer_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/proto/chat.proto",
}