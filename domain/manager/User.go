package manager

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
	"database/sql"
	"strconv"

	"gopkg.in/gorp.v1"
)

type User struct {
}

func (this *User) Query(param face.UserQueryParam) ([]face.UserBasic, int64) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "select * from User "
	csql := "select count(*) from User"
	wsql := " where 1=1 "
	if param.Name != "" {
		param.Name = "'%" + param.Name + "%'"
		wsql += " and name like :Name "
	}
	if param.Id > 0 {
		wsql += " and id = :Id "
	}
	if param.ProfessionId > 0 {
		wsql += " and professionId = :ProfessionId "
	}
	sql += wsql + " order by ct desc " + param.Limit()
	csql += wsql
	r := []face.UserBasic{}
	_, err := mysql.Select(&r, sql, param)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, param)
	exception.CheckMysqlError(err)
	return r, total
}

func (this *User) Get(id int) face.UserBasic {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "select * from User where id = ?"
	r := []face.UserBasic{}
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return r[0]
	}
	return face.UserBasic{}
}

func (this *User) Update(param face.UserUpdateParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := this.fetch(mysql, param.Id)
	r.Name = param.Name
	r.Birthday = sql.NullInt64{param.Birthday.Int64, true}
	r.Password = param.Password
	r.Phone = sql.NullString{param.Phone, true}
	r.ProfessionId = sql.NullInt64{int64(param.ProfessionId), true}
	r.Sex = param.Sex
	mysql.AddTable(*r).SetKeys(true, "Id")
	_, err := mysql.Update(r)
	exception.CheckMysqlError(err)
}

func (this *User) Add(param face.UserAddParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := table.User{}
	r.Name = param.Name
	r.Birthday = sql.NullInt64{param.Birthday.Int64, true}
	r.Group = param.Group
	r.Password = util.Md5(strconv.Itoa(111111))
	r.Phone = sql.NullString{param.Phone, true}
	r.ProfessionId = sql.NullInt64{int64(param.ProfessionId), true}
	r.Sex = param.Sex
	mysql.AddTable(r).SetKeys(true, "Id")
	err := mysql.Insert(r)
	exception.CheckMysqlError(err)
}

func (this *User) Del(id int) {
	sql := "delete from User where id = ?"
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	_, err := mysql.Exec(sql, id)
	exception.CheckMysqlError(err)
}

func (this *User) fetch(mysql *gorp.DbMap, id int) *table.User {
	sql := "select * from User where id = ?"
	r := []table.User{}
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}
