package main

import (
	"SGMS/route/admin"
	"log"
	"mime"
	"runtime"
	"time"
	"gopkg.in/alecthomas/kingpin.v2"
    "SGMS/route"
	"strconv"

	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

const ADMIN_ROOT = "/admin001"
var isDevMode bool

func main() {
	handleArgs()
	runtime.GOMAXPROCS(runtime.NumCPU())
	time.LoadLocation("Local")
	app := initApp()
	app.Static("/static", "./static", 1)
	app.Static("/static_resources", "./static_resources", 1)
	app.Static("/templates", "./templates", 1)
	mainsite(app)
	adminsite(app)

	app.Listen(":" + strconv.Itoa(9000))
}

func mainsite(app *iris.Framework) {
	app.Use(route.CatchException)
	app.Use(route.UrlAccessPermission)
	// route.Init(app)
	// teacher.Init(app)
}

func adminsite(app *iris.Framework) {
	adminApp := app.Party(ADMIN_ROOT)
	adminApp.Use(admin.ACLFilter)
	admin.Init(adminApp)
}


func handleArgs() {
	dev := kingpin.Flag("dev", "app work development mode").Short('d').Bool()
	// ssl := kingpin.Flag("ssl", "app work https mode").Short('s').Bool()
	kingpin.Parse()
	isDevMode = *dev
	// if *ssl {
	// 	config.IsHttps = true
	// }
}

func initApp() *iris.Framework {
	log.Println("initApp")
	mime.AddExtensionType(".css", "text/css")
	app := iris.New()
	if isDevMode {
		app.Config.IsDevelopment = true
		log.Println("app working development mode.")
	} else {
		log.Println("app working release mode.")
		app.Config.IsDevelopment = false
	}
	
	route.InitErrorPage(app)
	tplConfig := html.DefaultConfig()

	tplConfig.Funcs["adminroot"] = func() (string, error) { return ADMIN_ROOT, nil }
	tplConfig.Funcs["adminurl"] = func(url string) string { return ADMIN_ROOT + url }
	tplConfig.Funcs["isDevMode"] = func() bool { return isDevMode }
	route.EhanceTemplate(&tplConfig)
	app.UseTemplate(html.New(tplConfig)).Directory("./templates", ".html")

	return app
}