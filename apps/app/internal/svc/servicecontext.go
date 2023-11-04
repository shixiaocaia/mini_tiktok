package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/apps/user/userrpc"
)

type ServiceContext struct {
	Config config.Config

	UserRpc userrpc.UserRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRPC(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
