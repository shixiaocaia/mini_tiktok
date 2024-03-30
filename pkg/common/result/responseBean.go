package result

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"reflect"
)

type ResponseSuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const SuccessMsg = "success"
const StatusCode = 0

func Success(data interface{}) interface{} {
	// 对data 断言，修改其中的值
	getValue := reflect.ValueOf(data)
	field := getValue.Elem().FieldByName("StatusMsg")
	if field.CanSet() {
		field.SetString(SuccessMsg)
	} else {
		logx.Debug("cant set StatusMsg")
	}
	fieldCode := getValue.Elem().FieldByName("StatusCode")
	if !fieldCode.IsValid() {
		logx.Debug("Field StatusCode not found in the structure")
	}
	if fieldCode.CanSet() {
		fieldCode.SetInt(StatusCode)
	} else {
		logx.Debug("cant set StatusCode")
	}

	return data
}

type ResponseErrorBean struct {
	StatusCode uint32 `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}

func JwtUnauthorizedResult(w http.ResponseWriter, r *http.Request, err error) {
	httpx.WriteJson(w, http.StatusBadRequest, Error(401, err.Error()))
}
