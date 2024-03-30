package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	Id              int64  `gorm:"column:id; primary_key; auto_increment"` // 用户id
	Name            string `gorm:"column:user_name"`                       // 用户名
	Password        string `gorm:"column:password"`                        // 密码
	Follow          int64  `gorm:"column:follow_count"`                    // 关注数
	Follower        int64  `gorm:"column:follower_count"`                  // 粉丝数
	Avatar          string `gorm:"column:avatar"`                          // 头像
	BackgroundImage string `gorm:"column:background_image"`                // 背景图
	Signature       string `gorm:"column:signature"`                       // 签名
	TotalFavorited  int64  `gorm:"column:total_favorited"`                 // 获赞的总数
	FavoriteCount   int64  `gorm:"column:favorite_count"`                  // 点赞的视频总数
	WorkCount       int64  `gorm:"column:work_count"`                      // 发布的视频总数
}

func (r *User) TableName() string {
	return "user"
}

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{db: db}
}

func (m *UserModel) Insert(ctx context.Context, data *User) (int64, error) {
	result := m.db.WithContext(ctx).Create(&data)
	return data.Id, result.Error
}

func (m *UserModel) FindOneByUserName(ctx context.Context, userName string) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).
		Where("user_name = ?", userName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

func (m *UserModel) FindOneByUserId(ctx context.Context, userId int64) (*User, error) {
	var user User
	err := m.db.WithContext(ctx).
		Where("id = ?", userId).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}
