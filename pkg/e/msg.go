package e

import "os"

var errMsg = map[ErrorCode]map[string]string{
	SUCCESS:                      {"en": "ok", "ch": "成功"},
	ERROR:                        {"en": "error", "ch": "错误"},
	INVALID_PARAMS:               {"en": "params invalid", "ch": "请求参数异常或无效"},
	ERR_AUTH_CHECK_TOKEN_FAIL:    {"en": "auth fail", "ch": "Token鉴权失败"},
	ERR_AUTH_CHECK_TOKEN_TIMEOUT: {"en": "auth is timeout", "ch": "Token已过期"},
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

	ERR_UPLOAD_FILE_NOTFOUND:   {"en": "upload file not find", "ch": "上传文件未获取到"},
	ERR_UPLOAD_FILE_ERROR:      {"en": "upload file fail", "ch": "上传文件失败"},
	ERR_UPLOAD_FILE_QINIU_FAIL: {"en": "upload file fail by qiniuyun", "ch": "七牛云上传文件失败"},

	ERR_PARAMS_COMPANY_EMPTY_NAME:              {"en": "company name is empty", "ch": "公司名称不能为空"},
	ERR_PARAMS_COMPANY_EMPTY_CONTACT_NAME:      {"en": "contact name is  is empty", "ch": "公司联系人不能为空"},
	ERR_PARAMS_COMPANY_EMPTY_CONTACT_MOBILE:    {"en": "contact tel is empty", "ch": "公司联系电话不能为空"},
	ERR_PARAMS_COMPANY_EMPTY_ADDRESS:           {"en": "address is empty", "ch": "公司地址不能为空"},
	ERR_PARAMS_COMPANY_EMPTY_ADDRESS_POST:      {"en": "address pos is empty", "ch": "公司地址坐标不能为空"},
	ERR_PARAMS_COMPANY_EMPTY_LICENSE_PIC:       {"en": "license pic is empty", "ch": "营业执照不能为空"},
	ERR_COMPANY_NAME_EXIST:                     {"en": "company name exist", "ch": "公司名称已存在"},
	ERR_COMPANY_CONTACT_MOBILE_ERROR:           {"en": "contact tal invalid", "ch": "公司联系电话输入异常"},
	ERR_COMPANY_ADDRESS_POST_ERROR:             {"en": "address pos invalid", "ch": "公司地址坐标输入异常"},
	ERR_COMPANY_LICENSE_PIC_ERROR:              {"en": "license pic invalid", "ch": "营业执照输入异常"},
	ERR_COMPANY_CREATE_FAIL:                    {"en": "company create fail", "ch": "公司主体信息写入失败"},
	ERR_COMPANY_GET_FAIL:                       {"en": "company info get fail", "ch": "获取公司信息失败"},
	ERR_COMPANY_VERIFY_FAIL:                    {"en": "company verify fail", "ch": "公司主体审核失败"},
	ERR_PARAMS_COMPANY_VERIFY_EMPTY_COMPANY_ID: {"en": "company id is empty", "ch": "公司id不能为空"},
	ERR_PARAMS_COMPANY_VERIFY_STATUS_ERROR:     {"en": "company verify status invalid", "ch": "公司审核状态异常"},
}

var lang = os.Getenv("LANGUAGE")

func GetErrMsg(code ErrorCode) string {
	msg, ok := errMsg[code][lang]
	if ok {
		return msg
	}
	return errMsg[ERROR][lang]
}
