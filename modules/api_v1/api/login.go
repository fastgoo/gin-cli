package api

import (
	"gin-cli/pkg/e"
	"gin-cli/pkg/response"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	response.Success(c, e.SUCCESS)
}
