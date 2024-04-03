package video

import (
	"context"

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
	res, err := l.svcCtx.VideoRPC.VideoList(l.ctx, &video.VideoListRequest{
		UserId:   req.UserId,
		Cursor:   req.Cursor,
		SortType: req.SortType,
	})
	if err != nil {
		logx.Errorf("GetPublishVideoList failed, err: %v", err)
		return nil, nil
	}

	resp = &types.VideoListResp{
		VideoList: make([]types.VideoInfo, 0),
	}
	for _, v := range res.VideoList {
		resp.VideoList = append(resp.VideoList, types.VideoInfo{
			Id:            v.Id,
			AuthorId:      v.AuthorId,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		})
	}
	return
}
