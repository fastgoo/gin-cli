package gin

import (
	"fmt"
	"gin-cli/plugins/logger"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志记录到文件
func LoggerHandler() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(map[string]interface{}{
			"code":      statusCode,
			"time":      fmt.Sprintln(latencyTime),
			"client_ip": clientIP,
			"method":    reqMethod,
			"uri":       reqUri,
		}).Info()
	}
}
