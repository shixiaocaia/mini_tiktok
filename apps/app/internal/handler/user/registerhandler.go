package user

import (
	"net/http"

	"mini_tiktok/apps/app/internal/logic/user"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/pkg/httpresult"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}
		req.UserName = r.URL.Query().Get("username")
		req.Password = r.URL.Query().Get("password")

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		httpresult.HttpResult(r, w, resp, err)
	}
}
