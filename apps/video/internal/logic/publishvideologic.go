package logic

import (
	"context"

	"mini_tiktok/apps/video/internal/model"
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

func (l *PublishVideoLogic) PublishVideo(in *video.PublishVideoRequest) (*video.PublishVideoResponse, error) {
	_, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Video{
		AuthorId: in.AuthorId,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
		Title:    in.Title,
	})
	if err != nil {
		logx.Errorf("VideoModel.Insert failed, err: %v", err)
		return nil, err
	}
	return &video.PublishVideoResponse{}, nil
}
