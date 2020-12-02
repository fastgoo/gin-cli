package admin

import (
	"gin-cli/middleware/jwt"
	"gin-cli/modules/admin/api"
	"gin-cli/modules/admin/models"

	gRouter "gin-cli/router"
	//注册自定义验证器
	_ "gin-cli/modules/admin/validator"
)

func init() {
	models.Initialize()
	router()
}

func router() {
	admin := gRouter.Router.Group("/admin")
	auth := admin.Group("/")
	auth.Use(jwt.Auth())
	{
		auth.POST("/login", api.Login)
		auth.POST("/auth", api.Auth)
	}
	admin.POST("/login2", api.Login)
	admin.POST("/register", api.Register)
	admin.POST("/test", api.Test)
	admin.POST("/upload", api.Upload)
	//调试写入日志
	//logger.Info("test", "666", "hahah")
}
