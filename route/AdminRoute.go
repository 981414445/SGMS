package route

import (
	"SGMS/domain/face"
	"SGMS/domain/manager"
	"SGMS/domain/table"
	"strconv"

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
		partial := v.CheckQuery("partial").ToInt(0, "")
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
		if partial > 0 {
			ctx.MustRender("entry/admin/courses_list.html", data)
		} else {
			ctx.MustRender("entry/admin/courses.html", data)
		}
	})
	// 课程添加
	app.Post("admin/course/add", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseInsertParam{}
		param.Address = v.CheckBody("address").Empty().ToString()
		param.EndTime = v.CheckBody("endtime").NotEmpty().ToInt(0)
		param.Limit = v.CheckBody("limit").NotEmpty().ToInt(0)
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.StartTime = v.CheckBody("starttime").NotEmpty().ToInt(0)
		param.Status = 0
		param.TeacherId = v.CheckBody("teacherid").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Course).Add(param)
		Redirect(ctx, "/admin/course")
	})
	// 课程详情
	app.Get("admin/course/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("courseid").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			CourseDetail face.CourseDetail
		}{}
		data.CourseDetail = new(manager.Course).Get(id)
		data.User = SessionGetUser(ctx.Session())
		ctx.MustRender("entry/admin/course_detail.html", data)
	})
	// 课程内容修改
	app.Post("admin/course/update", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseUpdateParam{}
		param.Id = v.CheckBody("courseid").NotEmpty().ToInt(0)
		param.Address = v.CheckBody("address").Empty().ToString()
		param.EndTime = v.CheckBody("endtime").NotEmpty().ToInt(0)
		param.Limit = v.CheckBody("limit").NotEmpty().ToInt(0)
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.StartTime = v.CheckBody("starttime").NotEmpty().ToInt(0)
		param.TeacherId = v.CheckBody("teacherid").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Course).Update(param)
		Redirect(ctx, "/admin/course/detail?id="+strconv.Itoa(param.Id))
	})
	// 课程删除
	app.Get("admin/course/del", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("courseid").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Course).Del(id)
		Redirect(ctx, "/admin/course")
	})
	// 专业管理
	app.Get("/admin/major", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.ProfessionQueryParam{}
		param.Name = v.CheckQuery("name").Empty().ToString()
		param.No = v.CheckQuery("no").Empty().ToInt(0)
		param.TeacherId = v.CheckQuery("teacherid").Empty().ToInt(0)
		param.Si = v.CheckQuery("si").Empty().ToInt(0, "开始索引不是数字")
		param.Ps = v.CheckQuery("ps").Empty().ToInt(20, "每页显示不是数字")
		partial := v.CheckQuery("partial").ToInt(0, "")
		v.Check()
		data := struct {
			PageData
			List  []table.Profession
			Total int64
		}{}
		data.List, data.Total = new(manager.Profession).Query(param)
		if partial > 0 {
			ctx.MustRender("entry/admin/courses_list.html", data)
		} else {
			ctx.MustRender("/admin/majors.html", ctx)
		}
	})
	// 用户管理
	app.Get("/admin/users", func(ctx *iris.Context) {
		// v := NewValidatorContext(ctx)
	})
}
