package svc

import (
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/apps/like/likerpc"
	"mini_tiktok/apps/user/userrpc"
	"mini_tiktok/apps/video/videorpc"
	"mini_tiktok/pkg/interceptors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	defaultOssConnectTimeout   = 1
	defaultOssReadWriteTimeout = 3
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   userrpc.UserRPC
	VideoRPC  videorpc.VideoRPC
	LikeRPC   likerpc.LikeRPC
	OssClient *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	if c.Oss.ConnectTimeout == 0 {
		c.Oss.ConnectTimeout = defaultOssConnectTimeout
	}
	if c.Oss.ReadWriteTimeout == 0 {
		c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	}
	oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret,
		oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}

	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	videoRPC := zrpc.MustNewClient(c.VideoRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	likeRPC := zrpc.MustNewClient(c.LikeRpcConf, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:    c,
		UserRpc:   userrpc.NewUserRPC(userRPC),
		VideoRPC:  videorpc.NewVideoRPC(videoRPC),
		LikeRPC:   likerpc.NewLikeRPC(likeRPC),
		OssClient: oc,
	}
}
