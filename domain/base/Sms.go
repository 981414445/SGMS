package base

import (
	"SGMS/domain/config"
	"SGMS/domain/exception"
	"fmt"
	"net/http"

	"github.com/mozillazg/request"
)

type ISmsSender interface {
	Send()
}

func NewSmsSender(mobile, message string) ISmsSender {
	return &LuosimaoSmsSender{mobile, message}
}

type LuosimaoSmsSender struct {
	Mobile, Message string
}

// 发送短信验证码
func (this *LuosimaoSmsSender) Send() {
	c := new(http.Client)
	req := request.NewRequest(c)
	fmt.Println("key-" + config.LuosimaoKey)
	req.BasicAuth = request.BasicAuth{"api", "key-" + config.LuosimaoKey}
	req.Data = map[string]string{
		"message": this.Message + "【爱棋道】",
		"mobile":  this.Mobile,
	}
	resp, err := req.Post("http://sms-api.luosimao.com/v1/send.json")
	if nil != err {
		panic(err)
	}
	r, err := resp.Json()
	if status, err := r.Get("error").Int(); status != 0 || nil != err {
		errorBody, _ := resp.Text()
		panic(&exception.CodeError{exception.INTERNAL_ERROR, errorBody})
	}
}
