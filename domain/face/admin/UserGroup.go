package admin

import "SGMS/domain/face"

type IUesrGroup interface {
	//查询所有的管理员
	QueryAdmin(param UserGroupQueryAdminParam) ([]AdminInfo, int64)
	//获取管理员详情
	GetAdminDetail(uid string) *AdminDetail
	//更改用户主组
	ChangeMainGroup(uid, group int)
	//添加用户组
	AddGroup(uid, group int)
	//获取组的菜单
	GetGroupMenus(group int) []AdminMenuInfo
	//添加组菜单
	AddGroupMenu(group int) []AdminMenuInfo
	//删除组菜单
	DelGroupMenu(group int) []AdminMenuInfo
	//获取所有菜单
	GetMenus() []AdminMenuInfo
	//添加菜单
	AddMenu(menu *AdminMenuInfo)
	//删除菜单
	DelMenu(id int)
	//更新菜单
	UpdateMenu(menu AdminMenuInfo)
	//查询菜单分类
	QueryMenuSet()
	//添加菜单分类
	AddMenuSet()
	//更新菜单分类
	UpdateMenuSet()
	//删除菜单分类
	DelMenuSet()
}
type AdminInfo struct {
	face.User
	Groups []int
}

type AdminDetail struct {
	AdminInfo
}
type UserGroupQueryAdminParam struct {
	face.PageParam
	//mobile or phone or realName
	Key string
	//group
	Group int
}
type AdminMenuInfo struct {
	Id, SetId, No int
	Name, Url     string
	UIClass       *string
}
type GroupMenuInfo struct {
	Id, Group, MenuId, Mode int
}
type AdminMenuSetInfo struct {
}
