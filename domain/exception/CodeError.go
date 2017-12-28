package exception

import (
	"fmt"
)

//前端非域错误
type CodeError struct {
	Status  int
	Message string
}

func NewCodeError(status int, message string) *CodeError {
	return &CodeError{Status: status, Message: message}
}
func Panic(status int, message string) {
	panic(&CodeError{Status: status, Message: message})
}

func (this *CodeError) Error() string {
	return fmt.Sprintf("Status:%d,Message:%s", this.Status, this.Message)
}

//前端请求域错误，如校验错误
type RequestError struct {
	Status  int
	Message map[string]string
	//	RequestUrl string
}

func (this *RequestError) Error() string {
	return fmt.Sprintf("Status:%d,Message:%s", this.Status, this.Message)
}

// 参数错误
func NewParamError(errors map[string]string) *RequestError {
	return &RequestError{INVALID_PARAM, errors}
}

type HttpStatusError struct {
	Status  int
	Message map[string]interface{}
}

//Check mysql error
func CheckMysqlError(err error) {
	if nil != err {
		if _, ok := err.(*CodeError); ok {
			panic(err)
		}
		panic(&CodeError{MYSQL_ERROR, err.Error()})
	}
}
