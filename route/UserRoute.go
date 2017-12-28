package route

import (
	"fmt"

	"github.com/kataras/iris"
)

func RouteUser(app *iris.Framework) {
	app.Get("/users", func(ctx *iris.Context) {
		v := NewValidatorContext(ctx)
		name := v.CheckQuery("name").Empty().ToString()
		fmt.Println(name)
	})
}
