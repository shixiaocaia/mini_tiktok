package svc

import (
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/apps/user/userrpc"
	"mini_tiktok/pkg/interceptors"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userrpc.UserRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRPC(userRPC),
	}
}
