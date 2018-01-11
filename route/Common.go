package route

import (
	"SGMS/domain/face/admin"
	"path"

	"github.com/kataras/iris"
)

const ADMIN_ROOT = "/admin001"

type Bread struct {
	Url, Title string
}

type AdminData struct {
	User  *admin.Admin
	Error map[string]string
	Query map[string]string
	// PageUrl string
	Ctx        *iris.Context
	Breadcrumb []Bread
}

//压入面包屑
func (this *AdminData) PushBreadAdmin(url, title string) *AdminData {
	this.Breadcrumb = append(this.Breadcrumb, Bread{Url: ADMIN_ROOT + url, Title: title})
	return this
}

// type PageData struct {
// 	route.PaginationParam
// 	AdminData
// }

func GetRequestPath(ctx *iris.Context) string {
	return ctx.RequestPath(true)[len(ADMIN_ROOT):]
}

func Redirect(ctx *iris.Context, p string) {
	ctx.Redirect(path.Join(ADMIN_ROOT, p), iris.StatusOK)
}

func Ok(ctx *iris.Context, data ...interface{}) {
	Ok(ctx, data...)
}
