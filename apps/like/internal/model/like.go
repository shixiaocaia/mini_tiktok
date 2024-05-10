package model

import "gorm.io/gorm"

type Like struct {
	Id      int64 `gorm:"column:id; primary_key;"` // favorite_id
	UserId  int64 `gorm:"column:user_id"`          // user_id 谁点的赞
	VideoId int64 `gorm:"column:video_id"`         // video_id 被点赞的视频
}

func (r *Like) TableName() string {
	return "like"
}

type LikeModel struct {
	db *gorm.DB
}

func NewLikeModel(db *gorm.DB) *LikeModel {
	return &LikeModel{db: db}
}

func (m *LikeModel) Insert(data *Like) error {
	return m.db.Create(data).Error
}

func (m *LikeModel) DeleteByUserIdAndVideoId(userId, videoId int64) error {
	return m.db.Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&Like{}).Error
}
