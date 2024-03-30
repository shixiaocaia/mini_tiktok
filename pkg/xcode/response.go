package xcode

import (
	"net/http"

	"mini_tiktok/pkg/xcode/types"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)

	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
