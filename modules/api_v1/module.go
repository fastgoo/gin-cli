package api_v1

import (
	"gin-cli/middleware/jwt"
	"gin-cli/modules/api_v1/api"
	"gin-cli/modules/api_v1/models"
	"gin-cli/plugins/logger"

	cRouter "gin-cli/router"
)

func init() {
	models.SetUp()
	router()
}

func router() {
	authorized := cRouter.Router.Group("/")
	authorized.Use(jwt.Auth())
	{
		authorized.POST("/login", api.Login)
	}
	cRouter.Router.POST("/login2", api.Login)

	logger.Info("test","666","hahah")
}
