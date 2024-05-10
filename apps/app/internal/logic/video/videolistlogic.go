package video

import (
	"context"
	"encoding/json"

	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoListLogic) VideoList(req *types.VideoListReq) (resp *types.VideoListResp, err error) {
	uid, _ := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	if req.UserId != 0 {
		uid = req.UserId
	}
	res, err := l.svcCtx.VideoRPC.VideoList(l.ctx, &video.VideoListRequest{
		UserId:   uid,
		Cursor:   req.Cursor,
		SortType: req.SortType,
		PageSize: req.PageSize,
	})
	if err != nil {
		logx.Errorf("GetPublishVideoList failed, err: %v", err)
		return nil, nil
	}

	resp = &types.VideoListResp{
		VideoList: make([]types.VideoInfo, 0),
		Cursor:    res.Cursor,
		IsEnd:     res.IsEnd,
		VideoId:   res.VideoId,
	}
	for _, v := range res.VideoList {
		resp.VideoList = append(resp.VideoList, types.VideoInfo{
			Id:            v.Id,
			AuthorId:      v.AuthorId,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			PublishTime:   v.PublishTime,
		})
	}
	return
}
