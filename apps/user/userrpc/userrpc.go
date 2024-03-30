// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userrpc

import (
	"context"

	"mini_tiktok/apps/user/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq  = user.GenerateTokenReq
	GenerateTokenResp = user.GenerateTokenResp
	GetUserInfoReq    = user.GetUserInfoReq
	GetUserInfoResp   = user.GetUserInfoResp
	LoginReq          = user.LoginReq
	LoginResp         = user.LoginResp
	RegisterRequest   = user.RegisterRequest
	RegisterResponse  = user.RegisterResponse
	Request           = user.Request
	Response          = user.Response
	User              = user.User

	UserRPC interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUserRPC struct {
		cli zrpc.Client
	}
)

func NewUserRPC(cli zrpc.Client) UserRPC {
	return &defaultUserRPC{
		cli: cli,
	}
}

func (m *defaultUserRPC) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := user.NewUserRPCClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

func (m *defaultUserRPC) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserRPCClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserRPC) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserRPCClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserRPC) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserRPCClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
