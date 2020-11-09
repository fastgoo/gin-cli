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
	ERR_AUTH_TOKEN
	ERR_AUTH_OTHER
)
