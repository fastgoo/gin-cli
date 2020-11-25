package api

import (
	"gin-cli/modules/api_v1/models"
	myValidator "gin-cli/modules/api_v1/validator"
	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"gin-cli/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//注册用户
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

func Login(c *gin.Context) {
	//userInfoModel := &models.WkUserInfo{}
	response.Success(c, e.SUCCESS)
}

func Test(c *gin.Context) {
	response.Success(c, e.SUCCESS)
}
