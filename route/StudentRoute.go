package route

import "github.com/kataras/iris"

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

	// 教师课程
	app.Get("/student/course", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/courses.html", data)
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
}
