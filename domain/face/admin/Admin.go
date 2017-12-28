package admin

import (
	"SGMS/domain/face"
)

type IAdmin interface {
	//return user and groups
	Signin(phone, password string) *Admin
	// GetAdminMenu(uid int) ([]AdminMenuSet, []AdminMenu)
	// CheckPermission(groupMenu []AdminMenuSet, url, httpMethod string) bool
}

// type AdminGroupMenuSet struct {
// 	Id, Group int
// 	MenuSet   []*AdminMenuSet
// }

const (
	ADMIN_MENU_MODE_GET    = 1
	ADMIN_MENU_MODE_ADD    = 2
	ADMIN_MENU_MODE_UPDATE = 4
	ADMIN_MENU_MODE_DEL    = 8
)

type Admin struct {
	face.User
	MenuSets []*AdminMenuSet
	Menus    []*AdminMenu
}

type AdminMenuSet struct {
	Id      int
	Name    string
	UIClass string
	Menus   []*AdminMenu
}
type AdminMenu struct {
	Id   int
	Name string
	Url  string
	Mode int
	//是否显示在界面上
	Show    bool
	UIClass string
}
