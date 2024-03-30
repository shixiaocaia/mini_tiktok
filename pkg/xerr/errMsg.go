package xerr

var message map[uint32]string

var (
	CommonServerError *CodeError

	UserRegisterError       *CodeError
	UserDbFindUserNameError *CodeError
	UserNameExistsError     *CodeError
	UserInsertError         *CodeError
	UserNotExistedError     *CodeError
	UserPasswordError       *CodeError
	UserGenerateTokenError  *CodeError
)

func init() {
	message = make(map[uint32]string)

	// Global error code
	message[OK] = "Success"
	message[SERVER_COMMON_ERROR] = "Server error"

	// User xservice
	message[USER_REGISTER_ERROR] = "Register error"
	message[USER_DB_FIND_USERNAME_ERROR] = "Database error, username: %s, err: %v"
	message[USER_USERNAME_EXISTS] = "UserName already exists"
	message[USER_INSERT_ERROR] = "Insert error"
	message[USER_NOT_EXISTED_ERROR] = "User not existed"
	message[USER_PASSWORD_ERROR] = "Password error"
	message[USER_GENERATE_TOKEN_ERROR] = "Generate token error"

	CommonServerError = NewErrCode(SERVER_COMMON_ERROR)
	UserRegisterError = NewErrCode(USER_REGISTER_ERROR)
	UserDbFindUserNameError = NewErrCode(USER_DB_FIND_USERNAME_ERROR)
	UserNameExistsError = NewErrCode(USER_USERNAME_EXISTS)
	UserInsertError = NewErrCode(USER_INSERT_ERROR)
	UserNotExistedError = NewErrCode(USER_NOT_EXISTED_ERROR)
	UserPasswordError = NewErrCode(USER_PASSWORD_ERROR)
	UserGenerateTokenError = NewErrCode(USER_GENERATE_TOKEN_ERROR)
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
