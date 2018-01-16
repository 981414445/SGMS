package user

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
)

type User struct {
}

func (this *User) Signin(param face.UserSigninParam) table.User {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	// csql := "select count(*) from User where `no` = :Key"
	// c, err := mysql.SelectInt(csql, param)
	// exception.CheckMysqlError(err)
	// if c <= 0 {
	// 	return table.User{}
	// }
	psql := "select * from User where professionNo = :Key and (password = :Password or password = " + util.Md5(param.Password) + ")"
	u := []table.User{}
	_, err := mysql.Select(&u, psql, param)
	exception.CheckMysqlError(err)
	if len(u) > 0 {
		return u[0]
	}
	return table.User{}
}
