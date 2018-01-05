package manager

import (
	"database/sql"
	"gopkg.in/gorp.v1"
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
)

type CourseUser struct {
}

func (this *CourseUser) Add(param face.CourseUserAddParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	p := table.CourseUser{}
	p.CourseId = param.CourseId
	p.Uid = param.Uid
	p.Ct = util.Now()
	mysql.AddTable(p).SetKeys(true, "Id")
	err := mysql.Insert(&p)
	exception.CheckMysqlError(err)
}

func (this *CourseUser) Update(param face.CourseUserUpdateParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := this.fetch(mysql,param.Id)
	r.Grade = sql.NullInt64{param.Grade.Int64,true}
	_,err := mysql.Update(&r)
	exception.CheckMysqlError(err)
}

func (this *CourseUser) fetch(mysql *gorp.DbMap,id int) *table.CourseUser {
	sql := "select * from CourseUser where id = ?"
	r := []table.CourseUser{}
	_,err := mysql.Select(&r,sql,id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}

func (this *CourseUser) Del(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "delete from CourseUser where id = ?"
	_,err := mysql.Exec(sql,id)
	exception.CheckMysqlError(err)
}
