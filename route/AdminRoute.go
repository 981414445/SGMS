package route

import (
	"SGMS/domain/face"
	"SGMS/domain/manager"
	"SGMS/domain/table"

	"github.com/kataras/iris"
)

func RouteAdmin(app *iris.Framework) {
	// 课程管理
	app.Get("/admin/course", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseQueryParam{}
		param.Name = v.CheckQuery("name").Empty().ToString("")
		param.TeacherId = v.CheckQuery("teacherid").Empty().ToInt(0)
		param.StartTime = 0
		param.EndTime = 0
		param.Status = -1
		param.Si = v.CheckQuery("si").Empty().ToInt(0, "开始索引不是数字")
		param.Ps = v.CheckQuery("ps").Empty().ToInt(20, "每页显示不是数字")
		v.Check()
		data := struct {
			PageData
			List  []table.Course
			Total int64
		}{}
		ctx.Set("ps", param.Ps)
		data.User = SessionGetUser(ctx.Session())
		data.Ctx = ctx
		data.List, data.Total = new(manager.Course).Query(param)
		partial := v.CheckQuery("partial").ToInt(0, "")
		if partial > 0 {
			ctx.MustRender("entry/admin/courses_list.html", data)
		} else {
			ctx.MustRender("entry/admin/courses.html", data)
		}
	})
	// 专业管理
	app.Get("/admin/major", func(ctx *iris.Context) {
		// v := NewValidatorContext(ctx)
	})
	// 用户管理
	app.Get("/admin/users", func(ctx *iris.Context) {
		// v := NewValidatorContext(ctx)
	})
}
