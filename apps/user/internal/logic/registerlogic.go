package logic

import (
	"context"

	"mini_tiktok/apps/user/internal/code"
	"mini_tiktok/apps/user/internal/model"
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
	userInfo, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, in.Username)
	if err != nil {
		logx.Errorf("Register req: %v FindOneByUserName error: %v", in, err)
		return nil, err
	}

	if userInfo != nil {
		logx.Errorf("Register req: %v error: %v", in, err)
		return nil, code.UserNameExists
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)

	lastInsertId, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Name:            in.Username,
		Password:        string(hashPassword),
		Avatar:          "https://tse1-mm.cn.bing.net/th/id/R-C.d83ded12079fa9e407e9928b8f300802?rik=Gzu6EnSylX9f1Q&riu=http%3a%2f%2fwww.webcarpenter.com%2fpictures%2fGo-gopher-programming-language.jpg&ehk=giVQvdvQiENrabreHFM8x%2fyOU70l%2fy6FOa6RS3viJ24%3d&risl=&pid=ImgRaw&r=0",
		BackgroundImage: "https://tse2-mm.cn.bing.net/th/id/OIP-C.sDoybxmH4DIpvO33-wQEPgHaEq?pid=ImgDet&rs=1",
		Signature:       "test mode",
	})

	if err != nil {
		logx.Errorf("UserModel Insert %+v", err)
		return nil, err
	}

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
