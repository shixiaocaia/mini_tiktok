package logic

import (
	"context"

	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布视频
func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoRequest) (*video.PublishVideoResponse, error) {
	// todo: add your logic here and delete this line

	return &video.PublishVideoResponse{}, nil
}
