package logic

import (
	"context"

	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FeedLogic) Feed(in *video.FeedRequest) (*video.FeedResponse, error) {
	// todo: add your logic here and delete this line

	return &video.FeedResponse{}, nil
}
