package types

type LikeMsg struct {
	BizId    string ` json:"bizId,omitempty"`    // 业务id
	VideoId  int64  ` json:"videoId,omitempty"`  // 点赞对象id
	UserId   int64  ` json:"userId,omitempty"`   // 用户id
	LikeType int32  ` json:"likeType,omitempty"` // 类型
}

const (
	LikeTypeLike    int32 = 1
	LikeTypeDislike int32 = 0
)
