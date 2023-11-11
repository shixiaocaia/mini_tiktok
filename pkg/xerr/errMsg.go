package xerr

var message map[uint32]string

var (
	CommonServerError *CodeError

	RegisterError           *CodeError
	UserDbFindUserNameError *CodeError
	UserNameExistsError     *CodeError
)

func init() {
	message = make(map[uint32]string)

	// Global error code
	message[OK] = "Success"
	message[SERVER_COMMON_ERROR] = "Server error"

	// User xservice
	message[USER_REGISTER_ERROR] = "Register error"
	message[USER_DB_FIND_USERNAME_ERROR] = "Database error, username: %s, err: %v"
	message[USER_USERNAME_EXISTS] = "Username already exists"

	CommonServerError = NewErrCode(SERVER_COMMON_ERROR)
	RegisterError = NewErrCode(USER_REGISTER_ERROR)
	UserDbFindUserNameError = NewErrCode(USER_DB_FIND_USERNAME_ERROR)
	UserNameExistsError = NewErrCode(USER_USERNAME_EXISTS)
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "Oops, we are experiencing xservice shortage, please try again later."
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
