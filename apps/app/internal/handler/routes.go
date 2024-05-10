// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	like "mini_tiktok/apps/app/internal/handler/like"
	user "mini_tiktok/apps/app/internal/handler/user"
	video "mini_tiktok/apps/app/internal/handler/video"
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
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/userInfo",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: video.FeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/video"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: video.VideoListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/publish",
				Handler: video.PublishHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/douyin/video"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: like.LikeActionHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithSignature(serverCtx.Config.Signature),
		rest.WithPrefix("/douyin/like"),
	)
}
