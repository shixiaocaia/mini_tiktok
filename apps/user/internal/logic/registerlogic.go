package logic

import (
	"context"
	model2 "mini_tiktok/pkg/xmodel"

	"mini_tiktok/apps/user/internal/code"
	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"
	"mini_tiktok/pkg/jwt"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
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
	if in.GetUsername() == "" || in.GetPassword() == "" {
		return nil, code.RegisterInfoEmpty
	}

	// Check if the username exists
	userInfo, err := l.svcCtx.MysqlDB.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil && err != model2.ErrNotFound {
		logx.Errorf("Register req: %v FindOneByUserName error: %v", in, err)
		return nil, err
	}

	if userInfo != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, code.UserNameExists
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	dbUser := &model2.User{
		UserName: in.Username,
		Password: string(hashPassword),
	}

	res, err := l.svcCtx.MysqlDB.UserModel.Insert(l.ctx, dbUser)
	if err != nil {
		logx.Errorf("UserModel Insert %+v", err)
		return nil, err
	}
	lastInsertId, _ := res.LastInsertId()

	generateToken, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.JwtAuth.AccessSecret,
		AccessExpire: l.svcCtx.Config.JwtAuth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": lastInsertId,
		},
	})
	if err != nil {
		logx.Errorf("Register req: %v BuildTokens %+v", in, err)
		return nil, err
	}

	return &user.RegisterResponse{
		UserId: lastInsertId,
		Token:  generateToken.AccessToken,
	}, nil
}
