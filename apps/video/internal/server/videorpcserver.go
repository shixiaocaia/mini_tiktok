// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"mini_tiktok/apps/video/internal/logic"
	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/video"
)

type VideoRPCServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedVideoRPCServer
}

func NewVideoRPCServer(svcCtx *svc.ServiceContext) *VideoRPCServer {
	return &VideoRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoRPCServer) VideoList(ctx context.Context, in *video.VideoListRequest) (*video.VideoListResponse, error) {
	l := logic.NewVideoListLogic(ctx, s.svcCtx)
	return l.VideoList(in)
}

func (s *VideoRPCServer) PublishVideo(ctx context.Context, in *video.PublishVideoRequest) (*video.PublishVideoResponse, error) {
	l := logic.NewPublishVideoLogic(ctx, s.svcCtx)
	return l.PublishVideo(in)
}
