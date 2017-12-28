package route

import "github.com/kataras/iris"

func InitErrorPage(app *iris.Framework) {
	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.SetStatusCode(404)
		ctx.Write("404 NOT FOUND ERROR PAGE")
		ctx.Log(string(ctx.URI().FullURI()), "http status: 404 happened!")
		// data := PageData{}
		// data.Title = "404 未找到"
		// ctx.MustRender("static/404.html", data)
		// ctx.Redirect("/404")
	})
	iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
		ctx.Write(string(ctx.URI().FullURI()), " 500 INTERNAL SERVER ERROR PAGE")
		// or ctx.Render, ctx.HTML any render method you want
		ctx.Log("http status: 500 happened!")
	})
}
