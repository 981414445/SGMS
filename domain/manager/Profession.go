package manager

import (
	"gopkg.in/gorp.v1"
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
)

type Profession struct {
}

func (this *Profession) Query(param face.ProfessionQueryParam) ([]table.Profession, int64) {
	sql := "select * from Profession "
	csql := "select count(*) from Profession "
	wsql := " where 1=1 "
	if param.Name != "" {
		param.Name = "%" + param.Name + "%"
		wsql += " and name like " + param.Name
	}
	if param.No > 0 {
		wsql += " and no = :No "
	}
	if param.TeacherId > 0 {
		wsql += " and teacherId = :TeacherId "
	}
	sql += wsql + " order by ct desc " + param.Limit()
	csql += wsql
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := []table.Profession{}
	_, err := mysql.Select(&r, sql, param)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, param)
	exception.CheckMysqlError(err)
	return r, total
}

func (this *Profession) Get(id int) face.ProfessionDetail {
	sql := "select p.id,p.name as ProfessionName,p.teacherId,p.no,u.name as TeacherName,u.phone as TeacherPhone from Profession p left join User u on p.teacherId = u.id where p.id = ?"
	usql := "select id as uid,ct,name,phone,birthday from User where professionId = ?"
	mysql:= db.InitMysql()
	defer mysql.Db.Close()
	r := []face.ProfessionDetail{}
	_, err:=mysql.Select(&r,sql,id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		u := []face.ProfessionUsers{}
		_,err := mysql.Select(&u,usql,id)
		exception.CheckMysqlError(err)
		r[0].Users = u
		return r[0]
	}
	return face.ProfessionDetail{}
}

func (this *Profession) Add(param face.ProfessionInsertParam) {
	r := table.Profession{}
	r.Name = param.Name
	r.No = param.No
	r.TeacherId = param.TeacherId
	r.Ct = util.Now()
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	mysql.AddTable(r).SetKeys(true, "Id")
	err := mysql.Insert(&r)
	exception.CheckMysqlError(err)
}

func (this *Profession) Update(param face.ProfessionUpdateParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := this.fetch(mysql,param.Id)
	r.Name = param.Name
	r.No = param.No
	r.TeacherId = param.TeacherId
	mysql.AddTable(&r).SetKeys(true,"Id")
	_,err := mysql.Update(&r)
	exception.CheckMysqlError(err)
}

func (this *Profession) fetch(mysql *gorp.DbMap,id int) *table.Profession {
	r := []table.Profession{}
	sql := "select * from Profession where id = ?"
	_,err := mysql.Select(&r,sql,id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}

func (this *Profession) Del(id int) {
	mysql:= db.InitMysql()
	defer mysql.Db.Close()
	sql := "delete from Profession where id = ?"
	_,err :=mysql.Exec(sql,id)
	exception.CheckMysqlError(err)
}
