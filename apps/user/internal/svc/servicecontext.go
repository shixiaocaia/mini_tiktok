package svc

import (
	"mini_tiktok/apps/user/internal/config"
	"mini_tiktok/apps/user/internal/model"
	"mini_tiktok/pkg/orm"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config

	MysqlDB   *orm.DB
	UserModel *model.UserModel
	BizRedis  *redis.Redis
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
		Config:    c,
		MysqlDB:   db,
		UserModel: model.NewUserModel(db.DB),
		BizRedis:  rds,
	}
}
