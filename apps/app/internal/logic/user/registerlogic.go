package user

import (
	"context"

	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	pb "mini_tiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	registerRsp, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		l.Logger.Errorf("error %+v", err)
		return nil, err
	}
	return &types.RegisterResp{
		UserID: registerRsp.UserId,
		Token:  registerRsp.Token,
	}, nil
}
