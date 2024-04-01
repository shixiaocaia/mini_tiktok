package video

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"mini_tiktok/apps/app/internal/code"
	"mini_tiktok/apps/app/internal/svc"
	"mini_tiktok/apps/app/internal/types"
	"mini_tiktok/apps/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10MB

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *http.Request) (resp *types.PublishVideoResp, err error) {
	_ = req.ParseMultipartForm(maxFileSize)
	file, handler, err := req.FormFile("save_file")
	if err != nil {
		return nil, code.UploadVideoFailed
	}
	defer file.Close()

	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.VideoBucket)
	if err != nil {
		logx.Errorf("get bucket failed, err:%v", err)
		return nil, code.GetOssBucketErr
	}

	userId, err := l.ctx.Value(types.UserIdKey).(json.Number).Int64()
	objectKey := genFilename(userId, handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, code.PutBucketErr
	}
	logx.Info("upload success")

	// 插入记录
	_, err = l.svcCtx.VideoRPC.PublishVideo(l.ctx, &video.PublishVideoRequest{
		AuthorId: userId,
		PlayUrl:  genFileURL(objectKey),
		CoverUrl: "not support now",
		Title:    req.PostForm.Get("title"),
	})
	if err != nil {
		logx.Errorf("Insert video table failed, err: %v", err)
		return nil, err
	}

	return
}

func genFilename(userId int64, fileName string) string {
	extension := filepath.Ext(fileName)
	return fmt.Sprintf("%d_%d%s", time.Now().UnixMilli(), userId, extension)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://minitiktokvideo.oss-cn-shanghai.aliyuncs.com/%s", objectKey)
}
