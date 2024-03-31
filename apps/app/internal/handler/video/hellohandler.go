package video

import (
	"net/http"

	"mini_tiktok/apps/app/internal/logic/video"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/pkg/httpresult"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		httpresult.HttpResult(r, w, resp, err)
	}
}
