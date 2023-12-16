package user

import (
	"context"
	"mini_tiktok/apps/user/user"

	"github.com/jinzhu/copier"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userInfo, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		UserId: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	var user types.User
	_ = copier.Copy(&user, userInfo.User)
	resp = &types.UserInfoResp{
		User: user,
	}
	return
}
