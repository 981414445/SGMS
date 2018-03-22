package route

import (
	"SGMS/domain/face"
	"SGMS/domain/manager"
	"SGMS/domain/table"

	"github.com/kataras/iris"
)

func RouteStudent(app *iris.Framework) {
	// 教师首页
	app.Get("/student", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "登陆" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/student.html", data)
	})

	// 学生课程管理
	app.Get("/student/course", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseQueryParam{}
		param.Name = v.CheckQuery("name").Empty().ToString()
		param.TeacherId = v.CheckQuery("teacherid").Empty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Title string
			List  []table.Course
			Total int64
		}{}
		v.Check()
		data.Ctx = ctx
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		data.List, data.Total = new(manager.Course).Query(param)

		ctx.MustRender("entry/student/courses.html", data)
	})
	// 教师专业
	app.Get("/student/major", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/majors.html", data)
	})
	// 用户信息设置
	app.Get("/user/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Info  face.UserBasic
			Title string
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Ctx = ctx
		data.Info = new(manager.User).Get(id)
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/detail.html", data)
	})
}
