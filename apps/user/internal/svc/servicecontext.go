package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mini_tiktok/apps/user/internal/config"
	model2 "mini_tiktok/pkg/xmodel"
)

type ServiceContext struct {
	Config config.Config

	MysqlDB *model2.MysqlDB
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:  c,
		MysqlDB: NewMysqlDB(sqlConn),
	}
}

func NewMysqlDB(conn sqlx.SqlConn) *model2.MysqlDB {
	return &model2.MysqlDB{
		UserModel: model2.NewUserModel(conn),
	}
}
