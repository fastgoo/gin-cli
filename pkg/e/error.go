package e

import (
	"fmt"
	"gin-cli/plugins/logger"
)

type ServiceError struct {
	code ErrorCode
	msg  string
	Err  error
}

//自定义error，用于业务方面的处理
func (s ServiceError) Error() string {
	return fmt.Sprintf("%d - %s", s.code, s.msg)
}

func (s ServiceError) Code() ErrorCode {
	return s.code
}

func (s ServiceError) Msg() string {
	return s.msg
}

func New(code ErrorCode, err ...error) ServiceError {
	msg := GetErrMsg(code)
	var newErr error
	if len(err) > 0 {
		newErr = err[0]
	}
	return ServiceError{code, msg, newErr}
}

func CheckError(err ServiceError) {
	if err.Code() == 0 {
		logger.Errorf("msg - %s error - %v", err.Msg(), err.Err)
	}
}
