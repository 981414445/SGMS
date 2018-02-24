package route

import (
	"SGMS/domain/face"
	"SGMS/domain/manager"
	"SGMS/domain/table"
	"SGMS/domain/user"
	"fmt"

	"github.com/kataras/iris"
)

func RouteUser(app *iris.Framework) {
	app.Get("/home", func(ctx *iris.Context) {
		data := struct {
			PageData
		}{}
		data.User = SessionGetUser(ctx.Session())
		ctx.MustRender("entry/home.html", data)
	})
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
			if user.Group == 0 {
				// 学生
				ctx.MustRender("entry/home_student.html", data)
			} else if user.Group == 1 {
				// 老师
				ctx.MustRender("entry/home_teacher.html", data)
			} else {
				// 管理员
				ctx.MustRender("entry/home_admin.html", data)
			}
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

	// 学生首页
	app.Get("/student/home", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		v.Check()
		data := struct {
			PageData
			Info face.UserBasic
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Info = new(manager.User).Get(SessionGetUserId(ctx.Session()))
		ctx.MustRender("entry/student/home.html", data)
	})
	// 查看可选课程页面
	app.Get("/student/course/user/choose", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseUserQueryParam{}
		param.ProfessionId = v.CheckQuery("professionid").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			UnList []table.Course
			List   []table.Course
		}{}
		data.User = SessionGetUser(ctx.Session())
		param.Uid = SessionGetUserId(ctx.Session())
		// choose=0查未选课程，choose>0查已选课程
		param.Choose = 0
		data.UnList = new(manager.CourseUser).Query(param)
		param.Choose = 1
		data.List = new(manager.CourseUser).Query(param)
		ctx.MustRender("entry/student/course_choose.html", data)
	})
	// 学生选课
	app.Post("/student/course/user/add", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseUserAddParam{}
		param.CourseId = v.CheckBody("courseid").NotEmpty().ToInt(0)
		v.Check()
		param.Uid = SessionGetUserId(ctx.Session())
		new(manager.CourseUser).Add(param)
		Redirect(ctx, "/student/course/user/choose")
	})
	// 删除选课
	app.Get("/student/course/user/del", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		new(manager.CourseUser).Del(id)
		Redirect(ctx, "/student/course/user/choose")
	})
}
