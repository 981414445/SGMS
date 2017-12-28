package admin

import (
	"SGMS/domain/face"
)

type IUserManager interface {
	Query(param UserManagerQueryParam) ([]UserInfo, int64)
	Get(Id int) *UserDetail
	ChangePassword(Id int, passsword string)
	ChangeGroup(Id, group int)
	Update(user UserUpdateParam)
	Delete(Id int)
	Select2(param UserSelect2Param) []face.Select2Result
}

// UserManagerQueryParam 后台用户查询参数
type UserManagerQueryParam struct {
	Uid int
	// 继承PageParam  Si,startIndex Ps，PageSize
	face.PageParam
	//付费用户
	IsPay bool
	//评测用户
	IsEvaluate bool
	//禁用用户
	IsDisable bool
	//用户名/手机号/邮箱
	Key string
	//注册日期查询 开始时间，结束时间
	RegisterStartDate, RegisterEndDate int
	//登录日期查询 开始时间，结束时间
	SigninStartDate, SigninEndDate int
	//用户类型
	Group int
}

type UserInfo struct {
	Id, Group, Disabled int
	RealName, Icon1     string
	Email, Phone        *string
	Ct                  int
	Duan                float32
	Password            string
}

type UserDetail struct {
	UserInfo
	Weixin, Qq       string
	Gender, Birthday int
}

type UserUpdateParam struct {
	Id, Disabled                  int
	Duan                          float32
	RealName, Email, Phone, Icon1 string
	weixin, qq, Password          string
	Gender, Birthday, Group       int
}

type UserSelect2Param struct {
	face.Select2Param
	Group int
}

type ExportUserInfos struct {
	Id       int
	RealName string
	Email    *string
	Phone    *string
	Password string
	Ct       int
	Duan     float32
	//1:已禁用
	Disabled      int
	Group         int
	Icon1         string
	Icon2         string
	Icon3         string
	Weixin        *string
	Qq            *string
	Gender        *int
	Score         *int
	Birthday      *string
	LastSiginTime *int
}

type UserInfoUpdateParam struct {
	Id       int
	Disabled int
	Duan     float32
	RealName string
	Email    *string
	Phone    string
	Password string
}
