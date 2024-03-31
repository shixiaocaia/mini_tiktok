package code

import "mini_tiktok/pkg/xcode"

var (
	UploadVideoFailed = xcode.New(10001, "上传视频失败")
	GetOssBucketErr   = xcode.New(10002, "获取OSS-Bucket失败")
	GetObjDetailErr   = xcode.New(30003, "获取对象详细信息失败")
	PutBucketErr      = xcode.New(30002, "上传bucket失败")
)
