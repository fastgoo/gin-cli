package response

import (
	"gin-cli/pkg/e"
	"github.com/gin-gonic/gin"
)

type reponseH struct {
	Code e.ErrorCode `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// http异常响应
func Fail(c *gin.Context, statusCode int, code e.ErrorCode, msg string) {
	if msg == "" {
		msg = e.GetErrMsg(code)
	}
	c.JSON(statusCode, reponseH{code, msg, struct{}{}})
}

// http 成功响应
func Success(c *gin.Context, code e.ErrorCode, data ...interface{}) {
	var retData interface{}
	if len(data) == 0 || data[0] == nil {
		retData = struct{}{}
	} else {
		retData = data[0]
	}
	c.JSON(200, reponseH{code, e.GetErrMsg(code), retData})
}
