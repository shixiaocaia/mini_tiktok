package httpresult

import (
	"mini_tiktok/pkg/xcode"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type ResponseSuccessBean struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseFailBean struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

type Data struct {
	Code uint32 `json:"code"`
	Msg  string `json:"message"`
}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

func Fail(err error) *ResponseFailBean {
	//grpcStatus, _ := status.FromError(err)
	//xc := xcode.GrpcStatusToXCode(grpcStatus)
	xc := xcode.CodeFromError(err)
	return &ResponseFailBean{uint32(xc.Code()), xc.Message()}
}

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		r := Fail(err)
		httpx.WriteJson(w, http.StatusBadRequest, r)
	}
}
