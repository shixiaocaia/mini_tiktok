package logic

import (
	"context"
	"github.com/pkg/errors"
	"mini_tiktok/pkg/xerr"
	model2 "mini_tiktok/pkg/xmodel"

	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// Check if the username exists
	userInfo, err := l.svcCtx.MysqlDB.UserModel.FindOneByUserName(l.ctx, in.GetUsername())
	if err != nil && err != model2.ErrNotFound {
		userDbFindUserNameError := xerr.UserDbFindUserNameError.SetNewErrMsg("UserModel")
		l.Logger.Errorf("User Register error %+v", userDbFindUserNameError)
		return nil, errors.Wrapf(userDbFindUserNameError, "Database error, username: %s, err: %v", in.GetUsername(), err)
	}

	if userInfo != nil {
		err = errors.Wrapf(xerr.UserNameExistsError, "User already registered, email: %s, err: %v", in.GetUsername(), err)
		l.Logger.Errorf("error %+v", err)
		return nil, err
	}

	// todo 密码位数校验

	dbUser := &model2.User{
		UserName: in.GetUsername(),
		Password: in.GetPassword(),
	}

	res, err := l.svcCtx.MysqlDB.UserModel.Insert(l.ctx, dbUser)
	if err != nil {
		l.Logger.Errorf("UserModel Insert %+v-%+v", xerr.UserInsertError, err)
		return nil, xerr.UserInsertError
	}

	lastInsertId, _ := res.LastInsertId()

	// TODO: 生成token

	return &user.RegisterResponse{
		UserId: lastInsertId,
		Token:  "2333",
	}, nil
}
