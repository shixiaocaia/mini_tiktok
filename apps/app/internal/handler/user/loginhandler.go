package user

import (
	"net/http"

	"mini_tiktok/apps/app/internal/logic/user"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/pkg/httpresult"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		httpresult.HttpResult(r, w, resp, err)
	}
}
