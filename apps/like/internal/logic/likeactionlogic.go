package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/threading"
	"mini_tiktok/apps/like/internal/types"

	"mini_tiktok/apps/like/internal/svc"
	"mini_tiktok/apps/like/like"

	"encoding/json"
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

func (l *LikeActionLogic) LikeAction(in *like.LikeActionReq) (*like.LikeActionResp, error) {
	msg := &types.LikeMsg{
		BizId:    "like",
		VideoId:  in.VideoId,
		UserId:   in.UserId,
		LikeType: in.ActionType,
	}

	// 发送kafka消息，异步
	threading.GoSafe(func() {
		data, err := json.Marshal(msg)
		if err != nil {
			l.Logger.Errorf("likeMsg marshal msg: %v error: %v", msg, err)
			return
		}
		err = l.svcCtx.KqPusherClient.Push(string(data))
		if err != nil {
			l.Logger.Errorf("likeMsg kq push data: %s error: %v", data, err)
		}
	})

	return &like.LikeActionResp{}, nil
}
