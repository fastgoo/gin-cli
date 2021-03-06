package api

import (
	"fmt"
	"gin-cli/modules/admin/models"
	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"gin-cli/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	jwtMiddle "gin-cli/middleware/jwt"
	myValidator "gin-cli/modules/admin/validator"
)

// 注册账号
func Register(c *gin.Context) {
	var params myValidator.RegisterParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS, params.GetError(err.(validator.ValidationErrors)))
		return
	}
	passwordHash := util.HashAndSalt(params.Password)
	if passwordHash == "" {
		response.Fail(c, 200, e.ERR_PASSWORD_CRYPT_FAIL, "")
		return
	}
	userInfoModel := models.WkUserInfo{
		Username: params.Username,
		Password: passwordHash,
	}
	userId := userInfoModel.Insert(userInfoModel)
	if userId == 0 {
		response.Fail(c, 200, e.ERR_AUTH_TOKEN_NOTFOUND, "")
		return
	}
	response.Success(c, e.SUCCESS)
}

// 账号登陆操作
func Login(c *gin.Context) {
	var params myValidator.LoginParams
	if err := c.ShouldBind(&params); err != nil {
		response.Fail(c, 200, e.INVALID_PARAMS, params.GetError(err.(validator.ValidationErrors)))
		return
	}
	userInfoModel := &models.WkUserInfo{}
	userInfo, has := userInfoModel.Get("id,password,status", "username = ?", params.Username)
	if !has {
		response.Fail(c, 200, e.ERR_PARAMS_USERNAME_NOTEXIST)
		return
	}
	if !util.ComparePassword(userInfo.Password, params.Password) {
		response.Fail(c, 200, e.ERR_PASSWORD_ERROR)
		return
	}
	switch userInfo.Status {
	case -1:
		response.Fail(c, 200, e.ERR_USERNAME_NOT_ACTIVATED)
		return
	case 1:
		response.Fail(c, 200, e.ERR_USERNAME_LOCKED)
		return
	}
	token, err := jwtMiddle.CreateToken(jwtMiddle.Claims{UserId: userInfo.ID, Username: params.Username, Status: userInfo.Status, IP: c.ClientIP()})
	if err != nil {
		response.Fail(c, 200, e.ERR_AUTH_TOKEN_NOTCREATE)
		return
	}
	//写入登录的相关记录
	go func() {
		loginRecordModel := models.WkUserLoginRecord{
			UserID: userInfo.ID,
			IP:     c.ClientIP(),
		}
		ipInfo, err := util.GetIPInfo(loginRecordModel.IP)
		if err == nil {
			loginRecordModel.Area = ipInfo["country"].(string) + " " + ipInfo["province"].(string) + " " + ipInfo["city"].(string) +
				" " + ipInfo["district"].(string) + " [" + ipInfo["isp"].(string) + "]"
		}
		loginRecordModel.Insert(loginRecordModel)
	}()
	response.Success(c, e.SUCCESS, map[string]string{"token": token})
}

func Auth(c *gin.Context) {
	auth, _ := c.Get("authInfo")
	fmt.Println(auth.(*jwtMiddle.Claims))
}

func Test(c *gin.Context) {
	response.Success(c, e.SUCCESS)
}
