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
		auth.POST("/auth", api.Auth)
		admin.POST("/upload", api.Upload)

		admin.POST("/company/save", api.CompanySave)
		admin.POST("/company/change", api.CompanyChange)
		admin.POST("/company/verify", api.CompanyVerify)
		admin.GET("/company/list", api.CompanyList)
		admin.GET("/company/detail/:company_id", api.CompanyInfo)
	}
	admin.POST("/login", api.Login)
	admin.POST("/register", api.Register)
	admin.POST("/test", api.Test)

	//调试写入日志
	//logger.Info("test", "666", "hahah")
}
