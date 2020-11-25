package e

type ErrorCode uint32

const (
	SUCCESS        ErrorCode = 200
	ERROR          ErrorCode = 500
	INVALID_PARAMS ErrorCode = 10001
)

const (
	ERR_AUTH_CHECK_TOKEN_FAIL ErrorCode = iota + 10100
	ERR_AUTH_CHECK_TOKEN_TIMEOUT
	ERR_AUTH_TOKEN_NOTFOUND
	ERR_AUTH_TOKEN_NOTCREATE
	ERR_AUTH_OTHER

	ERR_PASSWORD_CRYPT_FAIL
	ERR_PASSWORD_NOTFOUND

	ERR_PARAMS_EMPTY_MOBILE
	ERR_PARAMS_EMPTY_PASSWORD
	ERR_PARAMS_MOBILE_INVALID
	ERR_PARAMS_USERNAME_EXIST
	ERR_PARAMS_USERNAME_NOTEXIST
	ERR_PARAMS_PASSWORD_INVALID
)
