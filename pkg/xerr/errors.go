package xerr

import (
	"fmt"
)

type CodeError struct {
	errCode    uint32
	errMsg     string
	errPayload map[string]string
}

// errorCode
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// errMsg
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) GetErrPayload() map[string]string {
	return e.errPayload
}

func (e *CodeError) SetErrMsg(errMsg string) *CodeError {
	e.errMsg = e.errMsg + " " + errMsg
	return e
}

func (e *CodeError) SetNewErrMsg(errMsg string) *CodeError {
	v := *e
	v.errMsg = errMsg
	return &v
}

func (e *CodeError) SetErrPayload(key, value string) {
	if e.errPayload == nil {
		e.errPayload = make(map[string]string)
	}
	e.errPayload[key] = value
}

func (e *CodeError) SetNewErrPayload(errPayload map[string]string) {
	e.errPayload = errPayload

}

func (e *CodeError) GetErrMsgValue() *CodeError {
	v := *e
	return &v
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}

func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: SERVER_COMMON_ERROR, errMsg: errMsg}
}
