package user

import (
	"mini_tiktok/pkg/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini_tiktok/apps/app/internal/logic/user"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		result.HttpResult(r, w, resp, err)
	}
}
