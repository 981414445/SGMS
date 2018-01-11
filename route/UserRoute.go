package route

import (
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/user"
	"fmt"

	"github.com/kataras/iris"
)

func RouteUser(app *iris.Framework) {
	app.Get("/users", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		name := v.CheckQuery("name").Empty().ToString()
		fmt.Println(name)
	})
	app.Get("/login", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		//Key, Password, WebLoginToken, WxOpenId
		param := &face.UserSigninParam{}
		param.Password = v.CheckQuery("password").Len(6, 30, "密码长度必须大于6位").ToString()
		param.Key = v.CheckQuery("number").NotEmpty().ToString()
		v.Check()
		user := new(user.User).Signin(*param)
		// 登录成功自动跳转到主页面
		// 登录失败返回登录页面
		if user.Id > 0 {
			data := struct {
				User table.User
			}{}
			data.User = user
			ssu := face.User{}
			ssu.Id = user.Id
			ssu.Group = user.Group
			SessionSetUser(ctx.Session(), &ssu)
			ctx.MustRender("entry/home.html", data)
		} else {
			data := struct {
				Info  string
				Title string
			}{}
			data.Title = "登陆" + HTML_TITLE_SUFFIX
			data.Info = "用户名或密码有误"
			ctx.MustRender("entry/login.html", data)
		}
	})

	app.Get("/signout", func(ctx *iris.Context) {
		ctx.SessionDestroy()
		Redirect(ctx, "/")
	})
}
