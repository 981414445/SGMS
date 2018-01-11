package route

import (
	"github.com/kataras/iris"
)

func RouteIndex(app *iris.Framework) {
	// 登陆
	app.Get("/", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "登陆" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/login.html", data)
	})

	// 个人信息页
	app.Get("/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/detail.html", data)
	})

}
