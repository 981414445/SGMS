package admin

import (
	"log"

	"github.com/kataras/iris"
)

//ACLFilter 用户访问权限控制
var ACLFilter = iris.HandlerFunc(
	func(ctx *iris.Context) {
		path := GetRequestPath(ctx)
		user := SessionGetAdmin(ctx.Session())
		log.Println("admin ", ctx.MethodString(), ctx.RequestPath(true))
		if "/" == path || "/signin" == path || user != nil && user.Group >= 100 {
			ctx.Next()
			return
		}
		if nil == SessionGetAdmin(ctx.Session()) {
			Redirect(ctx, "/")
			return
		}
		ctx.EmitError(iris.StatusProxyAuthRequired)
	})
