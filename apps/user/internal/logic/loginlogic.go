package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"mini_tiktok/pkg/xerr"
	model2 "mini_tiktok/pkg/xmodel"

	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
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
	userInfo, err := l.svcCtx.MysqlDB.UserModel.FindOneByUserName(l.ctx, in.GetUsername())
	if err != nil && err != model2.ErrNotFound {
		userLoginError := xerr.UserDbFindUserNameError.SetNewErrMsg("UserModel")
		l.Logger.Errorf("User Login error %+v", userLoginError)
		return nil, userLoginError
	}
	if userInfo == nil {
		return nil, xerr.UserNotExistedError
	}

	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(in.Password))
	if err != nil {
		l.Logger.Errorf("User Login error %+v", xerr.UserPasswordError)
		return nil, xerr.UserPasswordError
	}

	generateToken, err := l.svcCtx.Service.GenerateToken(&user.GenerateTokenReq{
		UserId: userInfo.Id,
	})
	if err != nil {
		return nil, xerr.UserGenerateTokenError
	}

	return &user.LoginResp{
		AccessToken: generateToken.AccessToken,
	}, nil
}
