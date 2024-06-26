package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf  zrpc.RpcClientConf
	VideoRpcConf zrpc.RpcClientConf
	LikeRpcConf  zrpc.RpcClientConf

	JwtAuth struct {
		AccessSecret  string
		AccessExpire  int64
		RefreshSecret string
		RefreshExpire int64
		RefreshAfter  int64
	}

	Oss struct {
		Endpoint         string
		AccessKeyId      string
		AccessKeySecret  string
		VideoBucket      string
		ConnectTimeout   int64 `json:",optional"`
		ReadWriteTimeout int64 `json:",optional"`
	}
}
