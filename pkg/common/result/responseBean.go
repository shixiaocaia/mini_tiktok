package result

import (
	"github.com/zeromicro/go-zero/core/logx"
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

//func (hr *ResponseSuccessBean) MarshalJSON() ([]byte, error) {
//	// 创建一个map来存储所有字段
//	result := map[string]interface{}{
//		"code": hr.Code,
//		"msg":  hr.Msg,
//	}
//
//	// 反射获取Payload的字段并添加到map中
//	val := reflect.ValueOf(hr.Data).Elem()
//	typ := reflect.TypeOf(hr.Data).Elem()
//	for i := 0; i < val.NumField(); i++ {
//		result[typ.Field(i).Name] = val.Field(i).Interface()
//	}
//
//	return json.Marshal(result)
//}
