package like

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini_tiktok/apps/app/internal/logic/like"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/pkg/httpresult"
)

func LikeActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewLikeActionLogic(r.Context(), svcCtx)
		resp, err := l.LikeAction(&req)
		httpresult.HttpResult(r, w, resp, err)
	}
}
