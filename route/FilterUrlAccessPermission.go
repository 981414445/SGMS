package route

import (
	"SGMS/domain/face"
	"strings"

	"SGMS/domain/exception"

	"github.com/kataras/iris"
)

var loginEnclude map[string]int = map[string]int{}
var loginInclude map[string]int = map[string]int{
	"/home":      1,
	"/course":    1,
	"/json/user": 2,
}
var UrlAccessPermission = iris.HandlerFunc(func(ctx *iris.Context) {
	user := SessionGetUser(ctx.Session())
	uri := string(ctx.URI().RequestURI())
	//老师页面和接口
	if (0 == strings.Index(uri, "/t/") || 0 == strings.Index(uri, "/json/t/")) && (nil == user || user.Group < face.USER_GROUP_TEACHER) {
		panic(exception.NewCodeError(exception.NO_AUTHORITY, "无权访问"))
	}
	ctx.Next()
})
