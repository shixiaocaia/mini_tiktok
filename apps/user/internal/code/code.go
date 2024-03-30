package code

import "mini_tiktok/pkg/xcode"

var (
	RegisterNameEmpty = xcode.New(20001, "注册名字不能为空")
	UserNameExists    = xcode.New(20002, "用户名已存在")
)
