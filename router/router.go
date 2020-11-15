package router

import (
	"os"

	"github.com/gin-gonic/gin"

	requestLogger "gin-cli/middleware/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	if os.Getenv("GIN_REQUEST_LOG") == "true" {
		Router.Use(requestLogger.LoggerHandler())
	}
}
