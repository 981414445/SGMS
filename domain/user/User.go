package user

import (
	"SGMS/domain/config"
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
	"database/sql"
	"time"

	gorp "gopkg.in/gorp.v1"
)

/**

type IUser interface {
	Get(id int) *UserInfo
	Signup(p *SignupParam) *UserInfo
	SignIn(p *SigninParam) *UserInfo
	Exist(key string) bool
	BindPhone(id int, phone string)
	BindEmail(id int, email string)
	UnbindPhone(id int)
	UnbindEmail(id int)
	UpdateIcon(id int, icon1, icon2, icon3 string)
	GetProfile(id int) *UserProfile
	UpdateProfile(profile *UserProfile)
}
*/
type User struct {
}

func (this *User) Get(id int) *face.User {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	return this.Fetch(mysql, id)
}
func (this *User) Fetch(mysql gorp.SqlExecutor, id int) *face.User {
	var r []face.User
	_, err := mysql.Select(&r, "select "+face.User{}.Sql()+" from User u where u.id=?", id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}

func (this *User) Signin(p *face.UserSigninParam) *face.User {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	id, err := mysql.SelectNullInt("select id from User where (phone=? or email=?) and password=?", p.Key, p.Key, util.Md5(p.Password))
	exception.CheckMysqlError(err)
	if !id.Valid {
		panic(&exception.CodeError{exception.USER_PASSWORD_ERROR, "密码错误"})
	}
	return this.signin(mysql, int(id.Int64))
}

func (this *User) SigninToken(p face.UserSigninToken) *face.User {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	// ssql := "select uid from UserToken where "
	// now := int(time.Now().Unix())
	// nowStr := strconv.Itoa(now)
	// if p.WebLoginToken != "" {
	// 	ssql += " webLoginToken=:WebLoginToken and webLoginTokenExpire<=" + nowStr
	// }
	// if p.WsToken != "" {
	// 	ssql += " WsToken=:WsToken and WsTokenExpire<=" + nowStr
	// }
	// if p.WxOpenId != "" {
	// 	ssql = "select uid from WeixinUser where openid={WxOpenId} "
	// }
	// ssql += " limit 1"
	// id, err := mysql.SelectNullInt(ssql, p)
	pt := UidTokenUidQueryParam{}
	pt.WebLoginToken = p.WebLoginToken
	pt.WsToken = p.WsToken
	id := new(UserToken).GetUid(mysql, pt)
	if id <= 0 {
		panic(&exception.CodeError{exception.USER_TOKEN_INVALID, "Signin Token Error"})
	}
	return this.signin(mysql, id)
}
func (this *User) signin(mysql *gorp.DbMap, id int) *face.User {
	this.CheckUserDisabled(mysql, id)
	return this.Fetch(mysql, id)
}

func (this *User) SignupTable(mysql *gorp.DbMap, user *table.User) {
	//	mysql := db.InitMysql()
	//	defer mysql.Db.Close()
	mysql.AddTable(table.User{}).SetKeys(true, "Id")
	user.Password = util.Md5(user.Password)
	user.Ct = int(time.Now().Unix())
	user.Pinyin = util.Pinyin(user.RealName)
	if "" == user.Icon1 {
		user.Icon1 = config.UserIcon1
		user.Icon2 = config.UserIcon2
		user.Icon3 = config.UserIcon3
	}
	err := mysql.Insert(user)
	exception.CheckMysqlError(err)
}
func (this *User) Signup(info *face.UserSignupParam) *face.User {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	return this.signup(mysql, info)
}
func (this *User) signup(mysql *gorp.DbMap, info *face.UserSignupParam) *face.User {
	user := &table.User{}
	user.Email = sql.NullString{String: info.Email, Valid: info.Email != ""}
	user.Phone = sql.NullString{String: info.Phone, Valid: info.Phone != ""}
	user.Icon1 = config.UserIcon1
	user.Icon2 = config.UserIcon2
	user.Icon3 = config.UserIcon3
	user.RealName = info.RealName
	user.Password = info.Password
	user.Group = info.Group
	this.SignupTable(mysql, user)
	return this.Fetch(mysql, user.Id)
}

func (this *User) del(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	user := &table.User{}
	user.Id = id
	_, err := mysql.Delete(user)
	exception.CheckMysqlError(err)
}
func (this *User) CheckUserDisabled(mysql *gorp.DbMap, id int) {
	count, err := mysql.SelectInt("select count(*) from User where disabled=1 and id=?", id)
	exception.CheckMysqlError(err)
	if count >= 1 {
		panic(&exception.CodeError{exception.USER_DISABLED, "用户已禁用"})
	}
}
func (this *User) exist(mysql *gorp.DbMap, key string) bool {
	count, err := mysql.SelectInt("select count(*) from User where (phone=? or email=?)", key, key)
	exception.CheckMysqlError(err)
	return count >= 1
}
func (this *User) Exist(key string) bool {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	return this.exist(mysql, key)
}

func (this *User) BindPhone(id int, phone string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	if this.exist(mysql, phone) {
		panic(&exception.CodeError{exception.USER_EXISTS, "phone号码已经被使用了"})
	}
	this.CheckUserDisabled(mysql, id)
	_, err := mysql.Exec("update User set phone=? where id=?", phone, id)
	exception.CheckMysqlError(err)
}
func (this *User) UnbindPhone(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, id)
	_, err := mysql.Exec("update User set phone=null where id=?", id)
	exception.CheckMysqlError(err)
}
func (this *User) BindEmail(id int, email string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, id)
	if this.exist(mysql, email) {
		panic(&exception.CodeError{exception.USER_EXISTS, "email已经被使用了"})
	}
	_, err := mysql.Exec("update User set email=? where id=?", email, id)
	exception.CheckMysqlError(err)
}
func (this *User) UnbindEmail(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, id)
	_, err := mysql.Exec("update User set email=null where id=?", id)
	exception.CheckMysqlError(err)
}

func (this *User) UpdateIcon(id int, icon1, icon2, icon3 string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, id)
	_, err := mysql.Exec("update User set icon1=?,icon2=?,icon3=? where Id=?", icon1, icon2, icon3, id)
	exception.CheckMysqlError(err)
}
func (this *User) GetProfile(id int) *face.UserProfile {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	profile := new(face.UserProfile)
	this.CheckUserDisabled(mysql, id)
	err := mysql.SelectOne(profile, "select Id,realName,Gender,Birthday,Qq,Weixin from User where id=?", id)
	exception.CheckMysqlError(err)
	return profile
}
func (this *User) UpdateProfile(profile *face.UserProfile) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, profile.Id)
	pinyin := util.Pinyin(profile.RealName)
	_, err := mysql.Exec("update User set realName=?,Gender=?,Birthday=?,Qq=?,Weixin=?,Pinyin=? where Id=?", profile.RealName, profile.Gender, profile.Birthday, profile.Qq, profile.Weixin, pinyin, profile.Id)
	exception.CheckMysqlError(err)
}
func (this *User) UpdatePassword(id int, oldPassword, newPassword string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	this.CheckUserDisabled(mysql, id)
	count, err := mysql.SelectInt("select count(*) from User where id=? and password=?", id, util.Md5(oldPassword))
	exception.CheckMysqlError(err)
	if count <= 0 {
		panic(&exception.CodeError{exception.USER_OLDPASSWORD_ERROR, "用户旧密码错误"})
	}
	_, err = mysql.Exec("update User set password=? where id=?", util.Md5(newPassword), id)
	exception.CheckMysqlError(err)
}
func (this *User) ResetPassword(mysql *gorp.DbMap, id int, password string) {
	this.CheckUserDisabled(mysql, id)
	_, err := mysql.Exec("update User set password=? where id=?", util.Md5(password), id)
	exception.CheckMysqlError(err)
}
func (this *User) ResetPasswordEmail(email string, password string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	id, err := mysql.SelectNullInt("select id from User where email=? limit 1", email)
	exception.CheckMysqlError(err)
	if !id.Valid {
		panic(exception.CodeError{exception.USER_NOT_EXISTS, "用户不存在"})
	}
	this.ResetPassword(mysql, int(id.Int64), password)
}

func (this *User) ResetPasswordPhone(phone string, password string) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	id, err := mysql.SelectNullInt("select id from User where phone=? limit 1", phone)
	exception.CheckMysqlError(err)
	if !id.Valid {
		panic(exception.CodeError{exception.USER_NOT_EXISTS, "用户不存在"})
	}
	this.ResetPassword(mysql, int(id.Int64), password)
}

func (this *User) GetUidByPhone(phone string) int {
	if "" == phone {
		return 0
	}
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	id, err := mysql.SelectNullInt("select id from User where phone=? limit 1", phone)
	exception.CheckMysqlError(err)
	return int(id.Int64)
}

//是管理员吗
func (this *User) IsManager(mysql gorp.SqlExecutor, uid int) bool {
	u := this.Fetch(mysql, uid)
	if nil == u {
		return false
	}
	return u.Group > face.USER_GROUP_STUDENT
}
