package logic

import (
	"context"
	model2 "mini_tiktok/pkg/xmodel"

	"mini_tiktok/apps/user/internal/code"
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
	// Check if the username is empty
	if in.GetUsername() == "" {
		return nil, code.RegisterNameEmpty
	}

	// Check if the username exists
	userInfo, err := l.svcCtx.MysqlDB.UserModel.FindOneByUserName(l.ctx, in.GetUsername())
	if err != nil && err != model2.ErrNotFound {
		logx.Errorf("Register req: %v FindOneByUserName error: %v", in, err)
		return nil, err
	}

	if userInfo != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, code.UserNameExists
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
