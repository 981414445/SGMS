package user

import (
	"SGMS/domain/face/admin"

	gorp "gopkg.in/gorp.v1"
)

type Admin struct {
	User
}

func (this *Admin) GetAdminMenu(mysql *gorp.DbMap, uid int) ([]*admin.AdminMenuSet, []*admin.AdminMenu) {
	// mysql := db.InitMysql()
	// defer mysql.Db.Close()
	var gm []struct {
		MenuId, MenuMode, MenuNo, MenuSetId, MenuSetNo int
		MenuName, MenuUrl, MenuSetName                 string
		MenuShow                                       bool
		MenuUIClass, MenuSetUIClass                    *string
	}
	_, err := mysql.Select(&gm, "select distinct ms.UIClass as MenuSetUIClass, gm.menuId,gm.mode as MenuMode,m.Show as MenuShow,m.UIClass as MenuUIClass, m.name as menuName,m.url as menuUrl,m.no as menuNo,ms.Id as MenuSetId,ms.Name as MenuSetName,ms.No as MenuSetNo from AdminGroupMenu gm left join AdminMenu m on m.id=gm.menuId left join AdminMenuSet ms on ms.id=m.setId where gm.group=(select `group` from User where id=?) or gm.id in (select `group` from UserGroup where uid=?) order by ms.no desc,m.no desc", uid, uid)
	if nil != err {
		panic(err)
	}
	var r []*admin.AdminMenuSet
	var r1 []*admin.AdminMenu
	msKeyMap := make(map[int]*admin.AdminMenuSet)
	for _, i := range gm {
		ms, ok := msKeyMap[i.MenuSetId]
		uiclass := ""
		if !ok {
			ms = &admin.AdminMenuSet{}
			ms.Id = i.MenuSetId
			ms.Name = i.MenuSetName
			if nil != i.MenuSetUIClass {
				uiclass = *i.MenuSetUIClass
			}
			ms.UIClass = uiclass
			msKeyMap[i.MenuSetId] = ms
			r = append(r, ms)
		}
		if nil != i.MenuUIClass {
			uiclass = *i.MenuUIClass
		}
		if i.MenuShow {
			ms.Menus = append(ms.Menus, &admin.AdminMenu{i.MenuId, i.MenuName, i.MenuUrl, i.MenuMode, i.MenuShow, uiclass})
		}
		r1 = append(r1, &admin.AdminMenu{i.MenuId, i.MenuName, i.MenuUrl, i.MenuMode, i.MenuShow, uiclass})
	}
	return r, r1
}
