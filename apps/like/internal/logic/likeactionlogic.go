package logic

import (
	"context"

	"mini_tiktok/apps/like/internal/svc"
	"mini_tiktok/apps/like/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeActionLogic {
	return &LikeActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeActionLogic) LikeAction(in *like.LikeActionRequest) (*like.LikeActionResponse, error) {
	// todo: add your logic here and delete this line

	return &like.LikeActionResponse{}, nil
}
