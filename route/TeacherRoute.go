package route

import (
	"SGMS/domain/face"
	"SGMS/domain/manager"
	"SGMS/domain/table"
	"database/sql"

	"github.com/guregu/null"
	"github.com/kataras/iris"
)

func RouteTeacher(app *iris.Framework) {
	// 教师首页
	app.Get("/teacher", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		v.Check()
		data := struct {
			PageData
			Info  face.UserBasic
			Title string
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Title = "登陆" + HTML_TITLE_SUFFIX
		data.Ctx = ctx
		data.Info = new(manager.User).Get(SessionGetUserId(ctx.Session()))
		ctx.MustRender("entry/teacher/home.html", data)
	})
	// 查看教师所负责专业的所有学生信息
	app.Get("/teacher/major/users", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Info face.ProfessionDetail
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Info = new(manager.Profession).Get(id)
		ctx.MustRender("entry/teacher/major.html", data)
	})
	// 查看教师所负责的所有课程
	app.Get("/teacher/courses", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseQueryParam{}
		param.TeacherId = v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			List  []table.Course
			Total int64
		}{}
		data.List, data.Total = new(manager.Course).Query(param)
		data.Ctx = ctx
		data.User = SessionGetUser(ctx.Session())
		ctx.MustRender("entry/teacher/courses.html", data)
	})
	// 查看教师所负责的某一课程及学生信息
	app.Get("/teacher/course/detail", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		id := v.CheckQuery("id").NotEmpty().ToInt(0)
		v.Check()
		data := struct {
			PageData
			Info face.CourseDetail
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Info = new(manager.Course).Get(id)
		ctx.MustRender("entry/teacher/course_detail.html", data)
	})
	// 教师给学生打分
	app.Get("/json/teacher/score/update", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseUserUpdateParam{}
		param.Id = v.CheckQuery("id").NotEmpty().ToInt(0)
		param.Score = null.Int{sql.NullInt64{int64(v.CheckQuery("score").NotEmpty().ToInt(0)), true}}
		v.Check()
		new(manager.CourseUser).Update(param)
		Ok(ctx)
	})
	// 教师课程 查询
	app.Get("/teacher/course", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		param := face.CourseQueryParam{}
		param.Name = v.CheckQuery("name").Empty().ToString()
		param.TeacherId = v.CheckQuery("teacherId").Empty().ToInt(0)
		param.Status = v.CheckQuery("status").Empty().ToInt(0)
		param.StartTime = v.CheckQuery("startTime").Empty().ToInt(0)
		param.EndTime = v.CheckQuery("endTime").Empty().ToInt(0)
		data := struct {
			face.PageParam
			Total   int64
			Title   string
			Courses []table.Course
		}{}
		data.Courses, data.Total = new(manager.Course).Query(param)
		v.Check()
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/courses.html", data)
	})
	// 教师专业
	app.Get("/teacher/major", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)

		v.Check()
		data := struct {
			PageData
			Title string
		}{}
		data.User = SessionGetUser(ctx.Session())
		data.Ctx = ctx
		data.Title = "个人信息页" + HTML_TITLE_SUFFIX
		ctx.MustRender("entry/teacher/majors.html", data)
	})
}
