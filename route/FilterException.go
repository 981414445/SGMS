package route

import (
	"SGMS/domain/exception"
	"log"
	"runtime/debug"

	"strings"

	"github.com/kataras/iris"
)

// 捕获异常
var CatchException = iris.HandlerFunc(func(ctx *iris.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ex, ok := err.(error); ok {
				log.Println(ex.Error())
			}
			if ex, ok := err.(*ValidatorError); ok {
				ctx.JSON(iris.StatusOK, exception.NewParamError(ex.Errors))
				return
			}
			if ex, ok := err.(*exception.CodeError); ok {
				uri := string(ctx.URI().RequestURI())
				if ex.Status == exception.USER_NO_SIGNIN && !strings.HasPrefix(uri, "/json") && !strings.HasPrefix(uri, "/signin") {
					ctx.Redirect("/signin?from=" + uri)
					return
				}
				if ex.Status > 0 && ex.Status < 10 {
					debug.PrintStack()
				}
				ctx.JSON(iris.StatusOK, ex)
				return
			}
			if ex, ok := err.(*exception.RequestError); ok {
				ctx.JSON(iris.StatusOK, ex)
				return
			}
			if ex, ok := err.(*exception.HttpStatusError); ok {
				ctx.EmitError(ex.Status)
				return
			}
			debug.PrintStack()
			//			debug.PrintStack()
			ctx.Log("Recovery from panic\n%s", err)
			ctx.SetStatusCode(500)
			//ctx.Panic just sends  http status 500 by default, but you can change it by: iris.OnPanic(func( c *iris.Context){})
			// ctx.Panic()
			return
		}
	}()
	log.Println(ctx.MethodString() + " " + string(ctx.URI().RequestURI()))
	ctx.Next()
})
