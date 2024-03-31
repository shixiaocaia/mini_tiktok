package video

import (
	"net/http"

	"mini_tiktok/apps/app/internal/logic/video"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/pkg/httpresult"
)

func PublishHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := video.NewPublishLogic(r.Context(), svcCtx)
		resp, err := l.Publish(r)
		httpresult.HttpResult(r, w, resp, err)
	}
}
