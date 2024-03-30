package user

import (
	"net/http"
	"strconv"

	"mini_tiktok/apps/app/internal/logic/user"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserInfoReq
		//if err := httpx.Parse(r, &req); err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//	return
		//}
		userId := r.URL.Query().Get("user_id")

		req.UserID, _ = strconv.ParseInt(userId, 10, 64)
		l := user.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
