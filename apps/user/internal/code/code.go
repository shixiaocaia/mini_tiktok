package code

import "mini_tiktok/pkg/xcode"

var (
	RegisterInfoEmpty = xcode.New(20001, "用户名或密码为空")
	UserNameExists    = xcode.New(20002, "用户名已存在")
	UserNotExist      = xcode.New(20003, "用户不存在")
	UserPasswordError = xcode.New(20004, "密码错误")
)
