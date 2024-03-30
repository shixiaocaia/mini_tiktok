package logic

import (
	"context"
	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"
	"mini_tiktok/pkg/xerr"
	"mini_tiktok/pkg/xmodel"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logger logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userInfo, err := l.svcCtx.MysqlDB.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil && err != xmodel.ErrNotFound {
		l.logger.Errorf("error %v-%+v-%+v", xmodel.ErrNotFound, in.UserId, err)
		return nil, err
	}

	if userInfo == nil {
		l.logger.Info("user not found")
		return nil, xerr.UserNotExistedError
	}

	return &user.GetUserInfoResp{User: ToUserInfo(userInfo)}, nil

}

func ToUserInfo(userInfo *xmodel.User) *user.User {
	return &user.User{
		Id:              userInfo.Id,
		Name:            userInfo.UserName,
		FollowCount:     userInfo.FollowCount,
		FollowerCount:   userInfo.FollowerCount,
		Avatar:          userInfo.Avatar.String,
		BackgroundImage: userInfo.BackgroundImage.String,
		Signature:       userInfo.Signature.String,
		TotalFavorited:  userInfo.TotalFavorited,
		FavoriteCount:   userInfo.FavoriteCount,
		WorkCount:       userInfo.WorkCount,
	}
}
