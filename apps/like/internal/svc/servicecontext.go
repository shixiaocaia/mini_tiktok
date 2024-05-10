package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"mini_tiktok/apps/like/internal/config"
	"mini_tiktok/apps/like/internal/model"
	"mini_tiktok/pkg/orm"
)

type ServiceContext struct {
	Config config.Config

	MysqlDB        *orm.DB
	LikeModel      *model.LikeModel
	BizRedis       *redis.Redis
	KqPusherClient *kq.Pusher
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
		Config:         c,
		MysqlDB:        db,
		LikeModel:      model.NewLikeModel(db.DB),
		BizRedis:       rds,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
	}
}
