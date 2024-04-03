package logic

import (
	"context"

	"mini_tiktok/apps/user/internal/code"
	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"
	"mini_tiktok/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	userInfo, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.GetUsername())
	if err != nil {
		logx.Errorf("FindOneByUserName userName: %v error %+v", in.GetUsername(), err)
		return nil, err
	}
	if userInfo == nil {
		return nil, code.UserNotExist
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password))
	if err != nil {
		return nil, code.UserPasswordError
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.JwtAuth.AccessSecret,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": userInfo.Id,
		},
	})
	if err != nil {
		logx.Errorf("BuildTokens %+v", err)
		return nil, err
	}

	return &user.LoginResp{
		AccessToken: token.AccessToken,
	}, nil
}
