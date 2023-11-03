// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "mini_tiktok/apps/app/internal/handler/user"
	"mini_tiktok/apps/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/hello",
				Handler: user.HelloHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)
}
