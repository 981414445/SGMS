package user

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"

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
