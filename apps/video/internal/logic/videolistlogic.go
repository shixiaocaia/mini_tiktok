package logic

import (
	"cmp"
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"mini_tiktok/apps/video/internal/code"
	"mini_tiktok/apps/video/internal/model"
	"mini_tiktok/apps/video/internal/svc"
	"mini_tiktok/apps/video/internal/types"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	prefixVideos         = "videos:%d:%d"
	prefixVideosByUserId = "videosByUserId:%d%d"
	videosExpire         = 60 * 60 * 24
)

type VideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoListLogic {
	return &VideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoListLogic) VideoList(in *video.VideoListRequest) (*video.VideoListResponse, error) {
	if in.SortType != types.SortPublishTime && in.SortType != types.SortFavoriteCount {
		return nil, code.SortTypeInvalid
	}
	if in.UserId <= 0 {
		return nil, code.UserIdInvalid
	}
	if in.PageSize == 0 {
		in.PageSize = types.DefaultPageSize
	}
	if in.Cursor == 0 {
		if in.SortType == types.SortPublishTime {
			in.Cursor = time.Now().Unix()
		} else {
			in.Cursor = types.DefaultSortLikeCursor
		}
	}

	var (
		err            error
		isCache, isEnd bool
		lastId, cursor int64
		videos         []*model.Video
		sortField      string
		curPage        []*video.VideoItem
	)
	if in.SortType == types.SortFavoriteCount {
		sortField = "favorite_count"
	} else {
		sortField = "publish_time"
	}

	// 查询缓存
	videoIds, _ := l.cacheVideoIds(l.ctx, in.UserId, in.Cursor, in.PageSize, in.SortType)
	if len(videoIds) > 0 {
		isCache = true
		if videoIds[len(videoIds)-1] == -1 {
			isEnd = true
		}
		videos, err = l.videosByIds(l.ctx, videoIds)
		if err != nil {
			return nil, err
		}

		// 通过sortType对videos进行排序
		var cmpFunc func(a, b *model.Video) int
		if sortField == "favorite_count" {
			cmpFunc = func(a, b *model.Video) int {
				return cmp.Compare(a.FavoriteCount, b.FavoriteCount)
			}
		} else {
			cmpFunc = func(a, b *model.Video) int {
				return cmp.Compare(a.PublishTime, b.PublishTime)
			}
		}
		slices.SortFunc(videos, cmpFunc)

		for _, v := range videos {
			curPage = append(curPage, &video.VideoItem{
				Id:            v.Id,
				AuthorId:      v.AuthorId,
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				Title:         v.Title,
				PublishTime:   v.PublishTime,
			})
		}
	} else {
		// 缓存不存在，查询数据库
		// 多个请求到来，阻塞等待返回结果
		v, err, _ := l.svcCtx.SingleFlightGroup.Do(videosByUserId(in.UserId, in.SortType), func() (interface{}, error) {
			return l.svcCtx.VideoModel.FindByUserId(l.ctx, in.UserId, types.DefaultLimit, sortField, in.Cursor)
		})
		if err != nil {
			logx.Errorf("VideoModel.FindByUserId userId:%v, pageSize:%v,sortField:%v, err:%v", in.UserId, in.PageSize, sortField, err)
			return nil, err
		}
		if v == nil {
			return &video.VideoListResponse{}, nil
		}
		videos = v.([]*model.Video)
		var firstPageVideos []*model.Video
		if len(videos) > int(in.PageSize) {
			firstPageVideos = videos[:in.PageSize]
		} else {
			firstPageVideos = videos
		}
		for _, v := range firstPageVideos {
			curPage = append(curPage, &video.VideoItem{
				Id:            v.Id,
				AuthorId:      v.AuthorId,
				PlayUrl:       v.PlayUrl,
				CoverUrl:      v.CoverUrl,
				FavoriteCount: v.FavoriteCount,
				CommentCount:  v.CommentCount,
				Title:         v.Title,
				PublishTime:   v.PublishTime,
			})
		}
	}

	if len(curPage) > 0 {
		pageLast := curPage[len(curPage)-1]
		lastId = pageLast.Id
		if in.SortType == types.SortPublishTime {
			cursor = pageLast.PublishTime
		} else {
			cursor = pageLast.FavoriteCount
		}

		// 去重，从大到小排序，如果游标相等，且视频ID相等，去掉该视频
		for k, v := range curPage {
			if in.SortType == types.SortPublishTime {
				if v.PublishTime == in.Cursor && v.Id == in.VideoId {
					curPage = curPage[k+1:]
					break
				}
			} else {
				if v.FavoriteCount == in.Cursor && v.Id == in.VideoId {
					curPage = curPage[k+1:]
					break
				}
			}
		}
	}

	ret := &video.VideoListResponse{
		IsEnd:     isEnd,
		Cursor:    cursor,
		VideoId:   lastId,
		VideoList: curPage,
	}

	// 异步写入缓存
	if !isCache {
		threading.GoSafe(func() {
			if len(videos) < types.DefaultLimit && len(videos) > 0 {
				// 视频到底
				videos = append(videos, &model.Video{Id: -1})
			}
			err = l.addCacheVideos(context.Background(), videos, in.UserId, in.SortType)
			if err != nil {
				logx.Errorf("addCacheVideos err: %v", err)
			}
		})
	}

	return ret, nil
}

func (l *VideoListLogic) cacheVideoIds(ctx context.Context, uid, cursor, pageSize int64, sortType int32) ([]int64, error) {
	key := videosKey(uid, sortType)
	// 查询缓存是否存在，存在更新键过期时间
	exists, err := l.svcCtx.BizRedis.ExistsCtx(ctx, key)
	if err != nil {
		logx.Errorf("ExistsCtx key: %s error: %v", key, err)
		return nil, err
	}
	if exists {
		err = l.svcCtx.BizRedis.ExpireCtx(ctx, key, videosExpire)
		if err != nil {
			logx.Errorf("ExpireCtx key: %s error: %v", key, err)
			return nil, err
		}
	}

	// 从缓存中按照指定分数范围和数量获取视频 ID 列表，并按照分数从大到小排序
	// ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
	// 查询当前游标之前的数据
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx, key, 0, cursor, 0, int(pageSize))
	if err != nil {
		logx.Errorf("ZrevrangebyscoreWithScoresAndLimit key: %s, err: %v", key, err)
		return nil, err
	}

	var ids []int64
	for _, pair := range pairs {
		id, err := strconv.ParseInt(pair.Key, 10, 64)
		if err != nil {
			logx.Errorf("strconv.ParseInt key: %s, err: %v", pair.Key, err)
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil

}

func videosKey(uid int64, sortType int32) string {
	return fmt.Sprintf(prefixVideos, uid, sortType)
}

func videosByUserId(uid int64, sortType int32) string {
	return fmt.Sprintf(prefixVideosByUserId, uid, sortType)
}

func (l *VideoListLogic) videosByIds(ctx context.Context, videoIds []int64) ([]*model.Video, error) {
	vid := make([]int64, 0)
	// 过滤掉-1, -1表示视频列表底部
	for _, id := range videoIds {
		if id != -1 {
			vid = append(vid, id)
		}
	}

	videos, err := l.svcCtx.VideoModel.FindByIds(ctx, vid)
	if err != nil {
		logx.Errorf("svcCtx.VideoModel.FindByIds err: %v", err)
		return nil, err
	}
	return videos, nil
}

func (l *VideoListLogic) addCacheVideos(ctx context.Context, videos []*model.Video, uid int64, sortType int32) error {
	if len(videos) == 0 {
		return nil
	}

	key := videosKey(uid, sortType)
	for _, v := range videos {
		var score int64
		if sortType == types.SortPublishTime {
			score = v.PublishTime
		} else {
			score = v.FavoriteCount
		}

		if score < 0 {
			score = 0
		}

		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, key, score, strconv.Itoa(int(v.Id)))
		if err != nil {
			return err
		}
	}
	return l.svcCtx.BizRedis.ExpireCtx(ctx, key, videosExpire)
}
