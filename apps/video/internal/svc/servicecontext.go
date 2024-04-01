package svc

import (
	"mini_tiktok/apps/video/internal/config"
	"mini_tiktok/apps/video/internal/model"
	"mini_tiktok/pkg/orm"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config

	MysqlDB    *orm.DB
	VideoModel *model.VideoModel
	BizRedis   *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.MustNewMysql(&orm.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})

	rds := redis.MustNewRedis(redis.RedisConf{
		Host: c.BizRedis.Host,
		Pass: c.BizRedis.Pass,
		Type: c.BizRedis.Type,
	})

	return &ServiceContext{
		Config:     c,
		MysqlDB:    db,
		BizRedis:   rds,
		VideoModel: model.NewVideoModel(db.DB),
	}
}
