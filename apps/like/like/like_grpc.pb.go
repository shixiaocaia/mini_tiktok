// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.1
// source: like.proto

package like

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

const (
	VideoRPC_LikeAction_FullMethodName = "/like.VideoRPC/LikeAction"
)

// VideoRPCClient is the client API for VideoRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoRPCClient interface {
	LikeAction(ctx context.Context, in *LikeActionRequest, opts ...grpc.CallOption) (*LikeActionResponse, error)
}

type videoRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoRPCClient(cc grpc.ClientConnInterface) VideoRPCClient {
	return &videoRPCClient{cc}
}

func (c *videoRPCClient) LikeAction(ctx context.Context, in *LikeActionRequest, opts ...grpc.CallOption) (*LikeActionResponse, error) {
	out := new(LikeActionResponse)
	err := c.cc.Invoke(ctx, VideoRPC_LikeAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoRPCServer is the server API for VideoRPC service.
// All implementations must embed UnimplementedVideoRPCServer
// for forward compatibility
type VideoRPCServer interface {
	LikeAction(context.Context, *LikeActionRequest) (*LikeActionResponse, error)
	mustEmbedUnimplementedVideoRPCServer()
}

// UnimplementedVideoRPCServer must be embedded to have forward compatible implementations.
type UnimplementedVideoRPCServer struct {
}

func (UnimplementedVideoRPCServer) LikeAction(context.Context, *LikeActionRequest) (*LikeActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LikeAction not implemented")
}
func (UnimplementedVideoRPCServer) mustEmbedUnimplementedVideoRPCServer() {}

// UnsafeVideoRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoRPCServer will
// result in compilation errors.
type UnsafeVideoRPCServer interface {
	mustEmbedUnimplementedVideoRPCServer()
}

func RegisterVideoRPCServer(s grpc.ServiceRegistrar, srv VideoRPCServer) {
	s.RegisterService(&VideoRPC_ServiceDesc, srv)
}

func _VideoRPC_LikeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LikeActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoRPCServer).LikeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoRPC_LikeAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoRPCServer).LikeAction(ctx, req.(*LikeActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoRPC_ServiceDesc is the grpc.ServiceDesc for VideoRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "like.VideoRPC",
	HandlerType: (*VideoRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LikeAction",
			Handler:    _VideoRPC_LikeAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "like.proto",
}