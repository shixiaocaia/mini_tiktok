// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videorpc

import (
	"context"

	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetFeedRequest              = video.GetFeedRequest
	GetFeedResponse             = video.GetFeedResponse
	GetPublishVideoListRequest  = video.GetPublishVideoListRequest
	GetPublishVideoListResponse = video.GetPublishVideoListResponse
	PublishVideoRequest         = video.PublishVideoRequest
	PublishVideoResponse        = video.PublishVideoResponse
	VideoInfo                   = video.VideoInfo

	VideoRPC interface {
		GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error)
		GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error)
		PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
	}

	defaultVideoRPC struct {
		cli zrpc.Client
	}
)

func NewVideoRPC(cli zrpc.Client) VideoRPC {
	return &defaultVideoRPC{
		cli: cli,
	}
}

func (m *defaultVideoRPC) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error) {
	client := video.NewVideoRPCClient(m.cli.Conn())
	return client.GetFeed(ctx, in, opts...)
}

func (m *defaultVideoRPC) GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error) {
	client := video.NewVideoRPCClient(m.cli.Conn())
	return client.GetPublishVideoList(ctx, in, opts...)
}

func (m *defaultVideoRPC) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	client := video.NewVideoRPCClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}