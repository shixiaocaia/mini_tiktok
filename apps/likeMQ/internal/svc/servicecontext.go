package svc

import (
	"mini_tiktok/apps/likeMQ/internal/config"
	"mini_tiktok/apps/likeMQ/internal/model"
	"mini_tiktok/pkg/orm"
)

type ServiceContext struct {
	Config    config.Config
	MysqlDB   *orm.DB
	LikeModel *model.LikeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	return &ServiceContext{
		Config:    c,
		MysqlDB:   db,
		LikeModel: model.NewLikeModel(db.DB),
	}
}
