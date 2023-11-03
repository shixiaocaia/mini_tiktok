// Code generated by goctl. DO NOT EDIT.
package types

type HelloReq struct {
	Name string `json:"name"`
}

type HelloResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResp struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}
