package model

import (
	"context"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64  `gorm:"column:id; primary_key;"` // video_id
	AuthorId      int64  `gorm:"column:author_id;"`       // 谁发布的
	PlayUrl       string `gorm:"column:play_url;"`        // videoURL
	CoverUrl      string `gorm:"column:cover_url;"`       // picURL
	FavoriteCount int64  `gorm:"column:favorite_count;"`  // 点赞数
	CommentCount  int64  `gorm:"column:comment_count;"`   // 评论数
	PublishTime   int64  `gorm:"autoCreateTime"`          // 发布时间
	Title         string `gorm:"column:title;"`           // 标题
}

func (r *Video) TableName() string {
	return "video"
}

type VideoModel struct {
	db *gorm.DB
}

func NewVideoModel(db *gorm.DB) *VideoModel {
	return &VideoModel{db: db}
}

func (m *VideoModel) Insert(ctx context.Context, data *Video) (int64, error) {
	result := m.db.WithContext(ctx).Create((&data))
	return data.Id, result.Error
}
