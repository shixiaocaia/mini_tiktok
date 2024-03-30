package logic

import (
	"context"
	model2 "mini_tiktok/pkg/xmodel"

	"mini_tiktok/apps/user/internal/code"
	"mini_tiktok/apps/user/internal/svc"
	"mini_tiktok/apps/user/user"

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
	if in.GetUsername() == "" {
		return nil, code.RegisterNameEmpty
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

	if in.Username == "" || in.Password == "" {
		err = errors.Wrapf(xerr.UserRegisterError, "User Register error, username: %s, password: %s, ", in.Username, in.Password)
		l.Logger.Errorf("User Register error %+v", err)
		return nil, err
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	dbUser := &model2.User{
		UserName: in.Username,
		Password: string(hashPassword),
	}

	res, err := l.svcCtx.MysqlDB.UserModel.Insert(l.ctx, dbUser)
	if err != nil {
		l.Logger.Errorf("UserModel Insert %+v-%+v", xerr.UserInsertError, err)
		return nil, xerr.UserInsertError
	}
	lastInsertId, _ := res.LastInsertId()

	generateToken, err := l.svcCtx.Service.GenerateToken(&user.GenerateTokenReq{
		UserId: lastInsertId,
	})
	if err != nil {
		return nil, xerr.UserGenerateTokenError
	}

	return &user.RegisterResponse{
		UserId: lastInsertId,
		Token:  generateToken.AccessToken,
	}, nil
}
