package logic

import (
	"context"

	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedLogic {
	return &GetFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 视频流
func (l *GetFeedLogic) GetFeed(in *video.GetFeedRequest) (*video.GetFeedResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetFeedResponse{}, nil
}
