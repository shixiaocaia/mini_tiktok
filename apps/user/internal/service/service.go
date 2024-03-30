package service

import (
	"mini_tiktok/apps/user/internal/config"

	"github.com/zeromicro/go-zero/core/logx"
)

type Service struct {
	config config.Config
	logx.Logger
}

func NewService(c config.Config) *Service {
	return &Service{
		config: c,
	}
}

// type UserClaims struct {
// 	UserId int64 `json:"user_id"`
// 	jwt.RegisteredClaims
// }

// func (l *Service) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
// 	accessExpire := l.config.JwtAuth.AccessExpire
// 	accessToken, err := l.getJwtToken(l.config.JwtAuth.AccessSecret, accessExpire, in.UserId)
// 	if err != nil {
// 		return nil, errors.Wrapf(xerr.UserGenerateTokenError, "getJwtToken err userId:%d , err:%v", in.UserId, err)
// 	}

// 	return &pb.GenerateTokenResp{
// 		AccessToken: accessToken,
// 	}, nil
// }

// func (l *Service) getJwtToken(secretKey string, expire, userId int64) (string, error) {
// 	userClaims := UserClaims{
// 		UserId: userId,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 			NotBefore: jwt.NewNumericDate(time.Now()),
// 			Issuer:    "app-user",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
// 	signed, err := token.SignedString([]byte(secretKey))
// 	return signed, err
// }
