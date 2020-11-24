package e

var errMsg = map[ErrorCode]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "请求参数错误",
	ERR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERR_AUTH_TOKEN_NOTCREATE:     "Token生成失败",
	ERR_AUTH_TOKEN_NOTFOUND:      "无法获取用户认证信息",
}

func GetErrMsg(code ErrorCode) string {
	msg, ok := errMsg[code]
	if ok {
		return msg
	}
	return errMsg[ERROR]
}
