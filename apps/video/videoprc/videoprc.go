// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package videoprc

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

	VideoPRC interface {
		// 视频流
		GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error)
		// 获取用户发布的视频
		GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error)
		// 发布视频
		PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error)
	}

	defaultVideoPRC struct {
		cli zrpc.Client
	}
)

func NewVideoPRC(cli zrpc.Client) VideoPRC {
	return &defaultVideoPRC{
		cli: cli,
	}
}

// 视频流
func (m *defaultVideoPRC) GetFeed(ctx context.Context, in *GetFeedRequest, opts ...grpc.CallOption) (*GetFeedResponse, error) {
	client := video.NewVideoPRCClient(m.cli.Conn())
	return client.GetFeed(ctx, in, opts...)
}

// 获取用户发布的视频
func (m *defaultVideoPRC) GetPublishVideoList(ctx context.Context, in *GetPublishVideoListRequest, opts ...grpc.CallOption) (*GetPublishVideoListResponse, error) {
	client := video.NewVideoPRCClient(m.cli.Conn())
	return client.GetPublishVideoList(ctx, in, opts...)
}

// 发布视频
func (m *defaultVideoPRC) PublishVideo(ctx context.Context, in *PublishVideoRequest, opts ...grpc.CallOption) (*PublishVideoResponse, error) {
	client := video.NewVideoPRCClient(m.cli.Conn())
	return client.PublishVideo(ctx, in, opts...)
}