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
		data.User = SessionGetUser(ctx.Session())
		data.List, data.Total = new(manager.Profession).Query(param)
		if partial > 0 {
			ctx.MustRender("entry/admin/majors_list.html", data)
		} else {
			ctx.MustRender("/admin/majors.html", ctx)
		}
	})
	// 专业添加
	app.Post("/admin/major/add", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.ProfessionInsertParam{}
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.No = v.CheckBody("no").NotEmpty().ToInt(0)
		param.TeacherId = v.CheckBody("teacherid").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Profession).Add(param)
		Redirect(ctx, "/admin/major")
	})
	// 专业详情
	app.Get("/admin/major/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Info face.ProfessionDetail
		}{}
		data.Info = new(manager.Profession).Get(id)
		data.User = SessionGetUser(ctx.Session())
		ctx.MustRender("/admin/major/detail.html", data)
	})
	// 专业修改
	app.Post("/admin/major/update", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.ProfessionUpdateParam{}
		param.Id = v.CheckBody("id").NotEmpty().ToInt(0)
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.No = v.CheckBody("no").NotEmpty().ToInt(0)
		param.TeacherId = v.CheckBody("teacherid").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Profession).Update(param)
		Redirect(ctx, "/admin/major/detail?id="+strconv.Itoa(param.Id))
	})
	// 专业删除
	app.Get("/admin/major/del", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		new(manager.Profession).Del(id)
		Redirect(ctx, "/admin/major")
	})
	// 用户管理
	app.Get("/admin/users", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.UserQueryParam{}
		param.Id = v.CheckQuery("id").Empty().ToInt(0)
		param.ProfessionId = v.CheckQuery("professionid").Empty().ToInt(0)
		param.Si = v.CheckQuery("si").Empty().ToInt(0, "开始索引不是数字")
		param.Ps = v.CheckQuery("ps").Empty().ToInt(20, "每页显示不是数字")
		partial := v.CheckQuery("partial").ToInt(0, "")
		v.Check()
		data := struct {
			PageData
			List  []face.UserBasic
			Total int64
		}{}
		data.List, data.Total = new(manager.User).Query(param)
		data.User = SessionGetUser(ctx.Session())
		if partial > 0 {
			ctx.MustRender("/admin/users_list.html", data)
		} else {
			ctx.MustRender("/admin/users.html", data)
		}
	})
	// 用户添加
	app.Post("/admin/user/add", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.UserAddParam{}
		birthday := v.CheckBody("birthday").Empty().ToString()
		param.Birthday.String = birthday
		param.Group = v.CheckBody("group").NotEmpty().ToInt(0)
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.Phone = v.CheckBody("phone").NotEmpty().ToString()
		param.ProfessionId = v.CheckBody("professionid").NotEmpty().ToInt(0)
		param.Sex = v.CheckBody("sex").Empty().ToInt(0)
		v.Check()
		new(manager.User).Add(param)
		Redirect(ctx, "admin/users")
	})
	// 用户详情
	app.Get("admin/user/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Info face.UserBasic
		}{}
		data.Info = new(manager.User).Get(id)
		data.User = SessionGetUser(ctx.Session())
		ctx.MustRender("admin/user/detail.html", ctx)
	})
	// 用户修改
	app.Post("/admin/user/update", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.UserUpdateParam{}
		param.Id = v.CheckBody("id").NotEmpty().ToInt(0)
		param.Birthday.String = v.CheckBody("birthday").Empty().ToString()
		param.Name = v.CheckBody("name").NotEmpty().ToString()
		param.Password = v.CheckBody("password").NotEmpty().ToString()
		param.Phone = v.CheckBody("phone").Empty().ToString()
		param.ProfessionId = v.CheckBody("professionid").NotEmpty().ToInt(0)
		param.Sex = v.CheckBody("sex").Empty().ToInt(0)
		v.Check()
		new(manager.User).Update(param)
		Redirect(ctx, "/admin/user/detail?id="+strconv.Itoa(param.Id))
	})
	// 用户删除
	app.Get("/admin/user/del", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckBody("id").NotEmpty().ToInt(0)
		v.Check()
		new(manager.User).Del(id)
		Redirect(ctx, "/admin/users")
	})
}
