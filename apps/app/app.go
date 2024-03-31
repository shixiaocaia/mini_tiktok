package main

import (
	"flag"
	"fmt"
	"mini_tiktok/apps/app/internal/config"
	"mini_tiktok/apps/app/internal/handler"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/pkg/jwt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(jwt.JwtUnauthorizedResult))

	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
