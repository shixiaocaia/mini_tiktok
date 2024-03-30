package user

import (
	"github.com/gorilla/schema"
	"mini_tiktok/pkg/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mini_tiktok/apps/app/internal/logic/user"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
)

func HelloHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HelloReq
		decoder := schema.NewDecoder()
		err := decoder.Decode(&req, r.URL.Query())
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewHelloLogic(r.Context(), svcCtx)
		resp, err := l.Hello(&req)
		result.HttpResult(r, w, resp, err)
	}
}
