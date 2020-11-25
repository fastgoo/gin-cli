package validator

import (
	"fmt"
	"gin-cli/modules/api_v1/models"
	"gin-cli/pkg/e"
	"gin-cli/pkg/util"
	"gin-cli/plugins/redis"
	"github.com/go-playground/validator/v10"
)

type RegisterParams struct {
	Username string `form:"username" json:"username" binding:"required,UsernameValid,UsernameNotExistValid"`
	Password string `form:"password"  json:"password"  binding:"required,PassSecurityValid"`
}

// 绑定模型获取验证错误的方法
func (r *RegisterParams) GetError(errs validator.ValidationErrors) string {
	for _, err := range errs {
		if err.Field() == "Username" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_MOBILE)
			case "UsernameValid":
				return e.GetErrMsg(e.ERR_PARAMS_MOBILE_INVALID)
			case "UsernameExistValid":
				return e.GetErrMsg(e.ERR_PARAMS_USERNAME_NOTEXIST)
			case "UsernameNotExistValid":
				return e.GetErrMsg(e.ERR_PARAMS_USERNAME_EXIST)
			}
		}
		if err.Field() == "Password" {
			switch err.Tag() {
			case "required":
				return e.GetErrMsg(e.ERR_PARAMS_EMPTY_PASSWORD)
			case "PassSecurityValid":
				return e.GetErrMsg(e.ERR_PARAMS_PASSWORD_INVALID)
			}
		}
	}
	return e.GetErrMsg(e.INVALID_PARAMS)
}

//校验用户名的是否为手机号码
func usernameValid(fl validator.FieldLevel) bool {
	return util.VerifyMobileFormat(fl.Field().String())
}

//校验用户账号存在，如果存在就通过，如果不存在就不通过
func usernameExistValid(fl validator.FieldLevel) bool {
	userModel := &models.WkUserInfo{}
	_, has := userModel.Get("id", "username = ?", fl.Field().String())
	return has
}

//校验用户账号不存在，如果不存在就通过，如果存在就不通过
func usernameNotExistValid(fl validator.FieldLevel) bool {
	key := "username:has:" + fl.Field().String()
	n := redis.Get(key)
	fmt.Print(n)
	userModel := &models.WkUserInfo{}
	_, has := userModel.Get("id", "username = ?", fl.Field().String())
	if has {
		err := redis.Set(key,1,15)
		fmt.Println(err)
	}
	return !has
}

//校验密码安全级别
func passSecurityValid(fl validator.FieldLevel) bool {
	return util.VerifyPasswordSecurity(fl.Field().String()) >= 2
}
