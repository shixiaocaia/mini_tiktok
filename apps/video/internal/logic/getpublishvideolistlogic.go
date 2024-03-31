package logic

import (
	"context"

	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublishVideoListLogic {
	return &GetPublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublishVideoListLogic) GetPublishVideoList(in *video.GetPublishVideoListRequest) (*video.GetPublishVideoListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetPublishVideoListResponse{}, nil
}
