package route

import (
	"github.com/kataras/iris"
)

func Init(app *iris.Framework) {
	RouteIndex(app)
	RouteStudent(app)
	RouteTeacher(app)
}
