package api_v1

import (
	"gin-cli/middleware/jwt"
	"gin-cli/modules/api_v1/api"
	"gin-cli/modules/api_v1/models"
	"gin-cli/plugins/logger"

	cRouter "gin-cli/router"
	//注册自定义验证器
	_ "gin-cli/modules/api_v1/validator"
)

func init() {
	models.Initialize()
	router()
}

func router() {
	authorized := cRouter.Router.Group("/")
	authorized.Use(jwt.Auth())
	{
		authorized.POST("/login", api.Login)
	}
	cRouter.Router.POST("/login2", api.Login)
	cRouter.Router.POST("/register", api.Register)
	cRouter.Router.POST("/test", api.Test)
	//调试写入日志
	logger.Info("test", "666", "hahah")
}
