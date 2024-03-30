package main

import (
	"flag"
	"fmt"
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/apps/app/internal/handler"
	"mini_tiktok/apps/app/internal/svc"
<<<<<<< HEAD
	"mini_tiktok/pkg/xcode"
=======
	"mini_tiktok/pkg/common/result"
>>>>>>> 4222557ea248bd679896f11fdfec79a333ccad96

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(result.JwtUnauthorizedResult))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误方法
	httpx.SetErrorHandler(xcode.ErrHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
