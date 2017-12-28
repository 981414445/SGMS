package face

import (
	"SGMS/domain/table"
	"database/sql"

	"github.com/guregu/null"
	gorp "gopkg.in/gorp.v1"
)

type IUser interface {
	Get(id int) *User
	Fetch(mysql gorp.SqlExecutor, id int) *User
	Signup(p *UserSignupParam) *User
	Signin(p *UserSigninParam) *User
	SigninToken(p UserSigninToken) *User
	// 用户是否存在
	Exist(key string) bool
	// 绑定手机号
	BindPhone(id int, phone string)
	// 绑定邮箱
	BindEmail(id int, email string)
	// 解除绑定手机号
	UnbindPhone(id int)
	// 解除绑定邮箱
	UnbindEmail(id int)
	// 将三种尺寸的图片保存
	UpdateIcon(id int, icon1, icon2, icon3 string)
	// 获取用户信息
	GetProfile(id int) *UserProfile
	// 修改用户信息
	UpdateProfile(profile *UserProfile)
	// 修改密码
	UpdatePassword(id int, oldPassword, newPassword string)
	// 通过邮箱修改密码
	ResetPasswordEmail(email string, password string)
	// 通过手机修改密码
	ResetPasswordPhone(phone string, password string)
}
type IIUser interface {
	UpdateDuan(mysql *gorp.DbMap, id int, duan float32)
	UpdateScore(mysql *gorp.DbMap, id int, score int)
}

type IWeixinUser interface {
	Save(weixinUser WeixinUser)
	Get(openId string) *WeixinUser
}
type UserSigninToken struct{ WsToken, WebLoginToken, WxOpenId string }
type UserSignupParam struct {
	Email, Phone, RealName, Password string
	Group                            int
}
type UserSigninParam struct{ Key, Password, WebLoginToken, WxOpenId string }
type UserPublicInfo struct {
	Id, Group           int
	Duan                float32
	RealName, Pinyin    string
	Icon1, Icon2, Icon3 string
	Gender              *int
}

func (this UserPublicInfo) Sql() string {
	return "u.Id,u.`group`,u.Duan,u.realName,u.pinyin,u.icon1,u.icon2,u.icon3,u.gender"
}

type User struct {
	Id                  int
	Group               int
	Duan                float32
	RealName, Pinyin    string
	Phone, Email        *string
	Icon1, Icon2, Icon3 string
	Gender              *int
}

func (this User) Sql() string {
	return " u.Id,u.`group`,u.Duan,u.realName,u.pinyin,u.icon1,u.icon2,u.icon3,u.gender,u.phone,u.email "
}

type UserProfile struct {
	Id       int
	Gender   null.Int
	Birthday null.String
	Qq       null.String
	Weixin   null.String
	RealName string
}

const (
	USER_GENDER_MALE   = 1
	USER_GENDER_FEMALE = 2
)
const (
	USER_GROUP_STUDENT = 0
	USER_GROUP_TEACHER = 1
	//客服
	USER_GROUP_SERVICE = 100
	//运营
	USER_GROUP_OP = 101
	//财务
	USER_GROUP_FINANCE = 102
	//客服经理
	USER_GROUP_SERVICE_PM = 200
	//运营经理
	USER_GROUP_OP_PM = 201
	//财务经理
	USER_GROUP_FINANCE_PM = 202

	//超级管理员
	USER_GROUP_ROOT = 10000
)

var USER_GENDER = map[int]string{USER_GENDER_MALE: "男", USER_GENDER_FEMALE: "女"}

var USER_GROUPS = map[int]string{USER_GROUP_STUDENT: "学生", USER_GROUP_TEACHER: "老师", USER_GROUP_SERVICE: "客服", USER_GROUP_OP: "运营", USER_GROUP_FINANCE: "财务",
	USER_GROUP_SERVICE_PM: "客服主管", USER_GROUP_OP_PM: "运营主管", USER_GROUP_FINANCE_PM: "财务主管",
}

const (
	WEIXIN_USER_SRC_SERVICE = 1 //微信服务号登录
	WEIXIN_USER_SRC_QR      = 2 //微信二维码登录
)

type WeixinUser struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	UnionId      string `json:"unionid"`
	ExpiresIn    int    `json:"expires_in"`
	Src          int
	Id, Uid      int
}

func (u *User) ToPublic() *UserPublicInfo {
	if nil == u {
		return nil
	}
	return &UserPublicInfo{
		Id:       u.Id,
		Group:    u.Group,
		Duan:     u.Duan,
		RealName: u.RealName,
		Icon1:    u.Icon1,
		Icon2:    u.Icon2,
		Icon3:    u.Icon3,
		Gender:   u.Gender,
		Pinyin:   u.Pinyin,
	}
}

const (
	// 未认证
	USER_VALIDATION_NOTSTART = 0
	// 认证中
	USER_VALIDATION_REVIEWING = 1
	// 认证通过
	USER_VALIDATION_SUCCESS = 2
	// 认证失败
	USER_VALIDATION_FAIL = 3
)

// 用户认证段位
const (
	USER_VALIDATION_RANK_FIFTEEN   = -15
	USER_VALIDATION_RANK_FOURTEEN  = -14
	USER_VALIDATION_RANK_THIRDTEEN = -13
	USER_VALIDATION_RANK_TWELVE    = -12
	USER_VALIDATION_RANK_ELEVEN    = -11
	USER_VALIDATION_RANK_TEN       = -10
	USER_VALIDATION_RANK_NINE      = -9
	USER_VALIDATION_RANK_EIGHT     = -8
	USER_VALIDATION_RANK_SEVEN     = -7
	USER_VALIDATION_RANK_SIX       = -6
	USER_VALIDATION_RANK_FIVE      = -5
	USER_VALIDATION_RANK_FOUR      = -4
	USER_VALIDATION_RANK_THREE     = -3
	USER_VALIDATION_RANK_TWO       = -2
	USER_VALIDATION_RANK_ONE       = -1
	USER_VALIDATION_DUAN_ONE       = 1
	USER_VALIDATION_DUAN_TWO       = 2
	USER_VALIDATION_DUAN_THREE     = 3
	USER_VALIDATION_DUAN_FOUR      = 4
	USER_VALIDATION_DUAN_FIVE      = 5
	USER_VALIDATION_DUAN_SIX       = 6
)

var USER_VALIDATION_STATUS = map[int]string{USER_VALIDATION_NOTSTART: "未认证", USER_VALIDATION_REVIEWING: "认证中", USER_VALIDATION_SUCCESS: "已认证", USER_VALIDATION_FAIL: "认证失败"}
var USER_VALIDATION_GRADE = map[int]string{USER_VALIDATION_RANK_FIFTEEN: "15级", USER_VALIDATION_RANK_FOURTEEN: "14级", USER_VALIDATION_RANK_THIRDTEEN: "13级", USER_VALIDATION_RANK_TWELVE: "12级", USER_VALIDATION_RANK_ELEVEN: "11级", USER_VALIDATION_RANK_TEN: "10级", USER_VALIDATION_RANK_NINE: "9级", USER_VALIDATION_RANK_EIGHT: "8级", USER_VALIDATION_RANK_SEVEN: "7级", USER_VALIDATION_RANK_SIX: "6级", USER_VALIDATION_RANK_FIVE: "5级", USER_VALIDATION_RANK_FOUR: "4级", USER_VALIDATION_RANK_THREE: "3级", USER_VALIDATION_RANK_TWO: "2级", USER_VALIDATION_RANK_ONE: "1级", USER_VALIDATION_DUAN_ONE: "1段", USER_VALIDATION_DUAN_TWO: "2段", USER_VALIDATION_DUAN_THREE: "3段", USER_VALIDATION_DUAN_FOUR: "4段", USER_VALIDATION_DUAN_FIVE: "5段", USER_VALIDATION_DUAN_SIX: "6段"}

type IUserValidation interface {
	//用户后台--上传认证信息
	Submit(param UserValidationParam)
	//管理后台--展示列表
	Query(param UserValidationQueryParam) ([]UserValidation, int64)
	//展示详细申请信息
	View(uid int) *table.UserValidation
	//管理后台--审核结果
	Review(uid, status int)
	//活动报名校验
	IsValid(uid int) int
	//用户认证查询
	Fetch(mysql gorp.SqlExecutor, uid int) *table.UserValidation
}

type UserValidationParam struct {
	Uid int
	//学生真实姓名
	Name string
	Duan int
	//段位证书图片
	DuanImg string
	//身份证
	IdImg    string
	Birthday int64
	Ct       int
	//1:男，2:女
	Gender int
	Status int
}
type UserValidation struct {
	Id  int
	Uid int
	//学生真实姓名
	Name string
	Duan sql.NullInt64
	//段位证书图片
	DuanImg sql.NullString
	//身份证
	IdImg    sql.NullString
	Birthday sql.NullInt64
	Ct       int
	//1:男，2:女
	Gender   sql.NullInt64
	Status   int
	RealName string
	Phone    string
}
type UserValidationQueryParam struct {
	PageParam
	Name, RealName, Phone string
	Status, Grade         int
}

type HistoryAdminUserValidationBasic struct {
	PageParam
	Id, AdminId, Action, Uid, Ct int
	AdminName, UserName, Phone   string
}

type HistoryAdminUserValidationQueryParam struct {
	PageParam
	AdminId, Action, Uid, StartTime, EndTime int
}
