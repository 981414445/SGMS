package route

import "github.com/kataras/iris"

func RouteTeacher(app *iris.Framework) {
	// 教师首页
	app.Get("/teacher", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "登陆" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/teacher.html", data)
	})

	// 教师课程 查询
	app.Get("/teacher/course", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseQueryParam{}
		param.Name = v.checkQuery("name").Empty().toString()
		param.TeacherId = v.checkQuery("teacherId").Empty().toInt(0)
		param.Status = v.checkQuery("status").Empty().toInt(0)
		param.StartTime = v.checkQuery("startTime").Empty().toInt(0)
		param.EndTime = v.checkQuery("endTime").Empty().toInt(0)
		data := struct {
			PageParam
			Title string
			Courses []table.Course
		}{}
		data.Course, data.Total = new(manager.Course).Query(param)
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/courses.html", data)
	})
	// 教师专业
	app.Get("/teacher/major", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		data := struct {
			Title string
		}{}
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/majors.html", data)
	})
}
