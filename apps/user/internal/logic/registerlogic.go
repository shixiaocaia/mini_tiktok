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
		userDbFindUserNameError := xerr.UserDbFindUserNameError.SetErrMsg("UserModel")
		l.Logger.Errorf("error %+v", userDbFindUserNameError)
		return nil, errors.Wrapf(userDbFindUserNameError, "Database error, username: %s, err: %v", in.GetUsername(), err)
	}

	if userInfo != nil {
		err = errors.New("Username already exists")
		l.Logger.Errorf("error %+v", err)
		return nil, err
	}

	dbUser := &model2.User{
		UserName: in.GetUsername(),
		Password: in.GetPassword(),
	}

	res, err := l.svcCtx.MysqlDB.UserModel.Insert(l.ctx, dbUser)
	if err != nil {
		l.Logger.Errorf("error %+v", err)
		return nil, err
	}
	l.Logger.Infof("res: %+v", res)
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		l.Logger.Errorf("error %+v", err)
	}

	// TODO: 生成token

	return &user.RegisterResponse{
		StatusCode: 200,
		StatusMsg:  "success",
		UserId:     lastInsertId,
		Token:      "2333",
	}, nil
}
