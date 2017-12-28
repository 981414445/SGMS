package route

import (
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/factory/basef"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"SGMS/domain/util"

	"log"

	"runtime/debug"

	"github.com/kataras/iris"
)

const (
	HTML_TITLE_SUFFIX = " - 爱棋道"
)

type PageData struct {
	Ctx     *iris.Context
	User    *face.User
	Title   string
	TabName string //导航tab高亮
	Now     int
	Query   map[string]string
}

type PaginationParam struct {
	Total int64
	// Si, Ps     int
	// RequestURI string
}
type Pagination struct {
	PageData
	PaginationParam
}

func SendFile(app *iris.Framework, path string, staticFile string) {
	app.Get(path, func(ctx *iris.Context) {
		ctx.ServeFile(staticFile, false)
	})
}

func RenderWithUser(app *iris.Framework, path string, template string) {
	app.Get(path, func(ctx *iris.Context) {
		ctx.MustRender(template, map[string]interface{}{"User": 1})
	})
}

type ResJson struct {
	Status int
	Data   interface{}
}
type OnePageData struct {
	Items interface{}
	Total int64
}

// 以JSON格式将错误信息返回前台
func Err(ctx *iris.Context, errorCode int, data ...interface{}) {
	r := new(ResJson)
	r.Status = errorCode
	if data != nil && len(data) > 0 {
		r.Data = data[0]
	}
	ctx.JSON(iris.StatusOK, r)
}

// 以JSON格式将数据返回前台
func Ok(ctx *iris.Context, data ...interface{}) {
	r := new(ResJson)
	r.Status = exception.OK
	if data != nil && len(data) > 0 {
		r.Data = data[0]
	}
	ctx.JSON(iris.StatusOK, r)
}

func OkPage(ctx *iris.Context, list interface{}, total int64) {
	Ok(ctx, OnePageData{list, total})
}

// 检查用户是否登录
func CheckSignin(ctx *iris.Context) {
	if nil == SessionGetUser(ctx.Session()) {
		panic(&exception.CodeError{exception.USER_NO_SIGNIN, "请登录"})
	}
}
func NotFound(ctx *iris.Context) {
	ctx.EmitError(http.StatusNotFound)
}

func ToJson(obj interface{}) string {
	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func GetPathParamInt(ctx *iris.Context, name string) int {
	r, err := strconv.Atoi(ctx.Param(name))
	if nil != err {
		panic(exception.HttpStatusError{iris.StatusNotFound, nil})
	}
	return r
}

// 静态页
func StaticPage(app *iris.Framework, route string, tpl string, title string) {
	app.Get(route, func(ctx *iris.Context) {
		type Data struct {
			PageData
			Now int64
		}
		data := new(Data)
		data.User = SessionGetUser(ctx.Session())
		data.Title = title + HTML_TITLE_SUFFIX
		data.Now = int64(time.Now().Unix()) * 1000
		// log.Println("StaticPage , tpl:" + tpl + ",route:" + route + ",title:" + title)
		defer func() {
			if err := recover(); err != nil {
				debug.PrintStack()
				log.Println("StaticPage , tpl:" + tpl + ",route:" + route + ",title:" + title)
			}
		}()
		err := ctx.Render(tpl, data)
		if nil != err {
			debug.PrintStack()
			log.Println("StaticPage , tpl:" + tpl + ",route:" + route + ",title:" + title)

		}
		// ctx.MustRender(tpl, data)
	})
}

func SaveImage(ctx *iris.Context, name string) string {
	field, err := ctx.FormFile(name)
	if nil == err && nil != field {
		file, err := field.Open()
		if nil != err {
			panic(exception.NewParamError(map[string]string{name: "请添加图片！"}))
		}
		defer file.Close()
		repo, err := basef.NewFileRepo().SaveImage(file, field.Filename, nil, nil)
		if nil != err {
			panic(exception.NewParamError(map[string]string{name: "请添加图片！"}))
		}
		return repo.RawImage
	}
	return ""
}

func FileToString(ctx *iris.Context, name string) (string, string, bool) {
	field, err := ctx.FormFile(name)
	if nil == err && nil != field {
		file, err := field.Open()
		if nil != err {
			panic(exception.NewParamError(map[string]string{name: "文件不正确"}))
		}
		defer file.Close()
		bs, err := ioutil.ReadAll(file)
		if nil != err {
			panic(exception.NewParamError(map[string]string{name: "文件不正确"}))
		}
		return field.Filename, string(bs), true
	}
	return "", "", false
}

type FileString struct {
	Name, FileName, Content string
}

func FileToStrings(ctx *iris.Context) ([]FileString, bool) {
	form, err := ctx.MultipartForm()
	var r []FileString
	if nil == err && nil != form {
		for k, fs := range form.File {
			for _, f := range fs {
				i := FileString{Name: k, FileName: f.Filename}
				file, err := f.Open()
				if nil != err {
					panic(exception.NewParamError(map[string]string{k: "文件不正确"}))
				}
				defer file.Close()
				bs, err := ioutil.ReadAll(file)
				if nil != err {
					panic(exception.NewParamError(map[string]string{k: "文件不正确"}))
				}
				i.Content = string(bs)
				r = append(r, i)
			}

		}
		return r, true
	}
	return nil, false
}

// 试题加密
func EncryptQuiz(sgf string, quizId int) string {
	s, err := util.AESEncryptBase64(util.Md5(strconv.Itoa(quizId)), sgf)
	if nil != err {
		panic(err)
	}
	return s
}

// 下载棋谱
func Download(ctx *iris.Context, name string, data []byte) {
	fileName := url.QueryEscape(name)
	ctx.SetHeader("Cache-Control", "max-age=0")
	ctx.SetHeader("X-Content-Type-Options", "nosniff")
	ctx.SetContentType("application/octet-stream")
	if ok, _ := regexp.Match("MSIE (?:6.0|7.0|8.0)", ctx.UserAgent()); ok {
		ctx.SetHeader("Content-Disposition", `attachment; filename="`+fileName+`"`)
	} else {
		ctx.SetHeader("Content-Disposition", `attachment; filename*=UTF-8''`+fileName)
	}
	ctx.SetBody(data)
}

// 获取IP地址
func GetRealIp(ctx *iris.Context) string {
	ip := ctx.RequestHeader("X-Forwarded-For")
	if "" != ip {
		return ip
	}
	return ctx.RemoteIP().String()
}
