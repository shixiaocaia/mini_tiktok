package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"mini_tiktok/apps/likeMQ/internal/model"
	"mini_tiktok/apps/likeMQ/internal/svc"
	"mini_tiktok/apps/likeMQ/internal/types"
)

type LikeMQLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeMQLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeMQLogic {
	return &LikeMQLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeMQLogic) Consume(key, val string) error {
	var likeMsg types.LikeMsg
	err := json.Unmarshal([]byte(val), &likeMsg)
	if err != nil {
		l.Logger.Errorf("unmarshal likeMsg error: %v", err)
		return err
	}
	if likeMsg.LikeType == types.LikeTypeLike {
		// 点赞
		err = l.svcCtx.LikeModel.Insert(&model.Like{
			VideoId: likeMsg.VideoId,
			UserId:  likeMsg.UserId,
		})
		if err != nil {
			l.Logger.Errorf("Insert error: %v", err)
			return err
		}
	} else {
		// 取消点赞
		err := l.svcCtx.LikeModel.DeleteByUserIdAndVideoId(likeMsg.UserId, likeMsg.VideoId)
		if err != nil {
			l.Logger.Errorf("DeleteByUserIdAndVideoId error: %v", err)
			return err
		}
	}
	// todo 消费后Offset如何提交的
	l.Logger.Info("Consume likeMsg success")
	return nil
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewLikeMQLogic(ctx, svcCtx)),
	}
}
