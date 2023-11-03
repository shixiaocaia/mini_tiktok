package xmodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var ErrNotFound = sqlx.ErrNotFound

type MysqlDB struct {
	UserModel UserModel
}
