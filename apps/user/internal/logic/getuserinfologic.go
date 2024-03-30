package logic

import (
	"context"
	"mini_tiktok/apps/user/internal/code"
	"mini_tiktok/apps/user/internal/model"
	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"

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
	userInfo, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("UserModel.FindOneByUserId error %+v-%+v", in.UserId, err)
		return nil, err
	}

	if userInfo == nil {
		return nil, code.UserNotExist
	}

	return &user.GetUserInfoResp{User: ToUserInfo(userInfo)}, nil

}

func ToUserInfo(userInfo *model.User) *user.User {
	return &user.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		FollowCount:     userInfo.Follow,
		FollowerCount:   userInfo.Follower,
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:       userInfo.Signature,
		TotalFavorited:  userInfo.TotalFavorited,
		FavoriteCount:   userInfo.FavoriteCount,
		WorkCount:       userInfo.WorkCount,
	}
}
