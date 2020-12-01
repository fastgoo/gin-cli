package e

import "os"

var errMsg = map[ErrorCode]map[string]string{
	SUCCESS:                      {"en": "ok", "ch": "成功"},
	ERROR:                        {"en": "error", "ch": "错误"},
	INVALID_PARAMS:               {"en": "params invalid", "ch": "请求参数异常或无效"},
	ERR_AUTH_CHECK_TOKEN_FAIL:    {"en": "auth fail", "ch": "Token鉴权失败"},
	ERR_AUTH_CHECK_TOKEN_TIMEOUT: {"en": "auth is timeout", "ch": "Token已超时"},
	ERR_AUTH_TOKEN_NOTCREATE:     {"en": "token create fail", "ch": "Token生成失败"},
	ERR_AUTH_TOKEN_NOTFOUND:      {"en": "cannt get userinfo", "ch": "无法获取用户认证信息"},

	ERR_PASSWORD_CRYPT_FAIL: {"en": "password crypt error", "ch": "密码加密出现异常"},
	ERR_PASSWORD_ERROR:      {"en": "password error", "ch": "密码错误"},

	ERR_PARAMS_EMPTY_MOBILE:     {"en": "mobile param is empty", "ch": "手机号码不能为空"},
	ERR_PARAMS_MOBILE_INVALID:   {"en": "mobile param  invalid", "ch": "手机号码输入有误"},
	ERR_PARAMS_EMPTY_PASSWORD:   {"en": "password param is empty", "ch": "密码不能为空"},
	ERR_PARAMS_PASSWORD_INVALID: {"en": "password security too low. len >= 6, must letters and numbers", "ch": "密码安全级别太低，密码需要字符串+数字的组合，同时不可低于6位"},

	ERR_PARAMS_USERNAME_EXIST:    {"en": "username exist", "ch": "该账号已存在"},
	ERR_PARAMS_USERNAME_NOTEXIST: {"en": "username not exist", "ch": "该账号不已存在"},

	ERR_USERNAME_NOT_ACTIVATED: {"en": "username not activated", "ch": "该账号未激活"},
	ERR_USERNAME_LOCKED:        {"en": "username is locked", "ch": "该账号已被锁定"},
}

var lang = os.Getenv("LANGUAGE")

func GetErrMsg(code ErrorCode) string {
	msg, ok := errMsg[code][lang]
	if ok {
		return msg
	}
	return errMsg[ERROR][lang]
}
