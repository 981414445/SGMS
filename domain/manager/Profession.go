package manager

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
	"strconv"

	"github.com/guregu/null"

	"gopkg.in/gorp.v1"
)

type Profession struct {
}

func (this *Profession) Query(param face.ProfessionQueryParam) ([]table.Profession, int64) {
	sql := "select * from Profession "
	csql := "select count(*) from Profession "
	wsql := " where 1=1 "
	if param.Name != "" {
		param.Name = "'%" + param.Name + "%'"
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
	usql := "select id as uid,ct,name,phone,birthday,sex,professionNo as No from User where professionId = ?"
	csql := "select c.* from Course c left join ProfessionCourse pc on c.id = pc.courseid where pc.professionid = ?"
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := []face.ProfessionDetail{}
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		c := []table.Course{}
		_, err := mysql.Select(&c, csql, id)
		exception.CheckMysqlError(err)
		r[0].Courses = c
		u := []face.ProfessionUsers{}
		_, err = mysql.Select(&u, usql, id)
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
	r := this.fetch(mysql, param.Id)
	r.Name = param.Name
	r.No = param.No
	r.TeacherId = param.TeacherId
	mysql.AddTable(&r).SetKeys(true, "Id")
	_, err := mysql.Update(&r)
	exception.CheckMysqlError(err)
}

func (this *Profession) fetch(mysql *gorp.DbMap, id int) *table.Profession {
	r := []table.Profession{}
	sql := "select * from Profession where id = ?"
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}

func (this *Profession) Del(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "delete from Profession where id = ?"
	_, err := mysql.Exec(sql, id)
	exception.CheckMysqlError(err)
}

type TeacherProfession struct {
	ProfessionId, Ct, No, TeacherId int
	Name                            string
}

// 老师专业列表
// 专业详情(学生信息)
func (this *Profession) GetTeacherProfession(teacherId int) ([]TeacherProfession, int64) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "select id as ProfessionId,name,ct,no,teacherId from Profession where teacherId = ?"
	csql := "select count(*) from Profession where teacherId = ?"
	r := []TeacherProfession{}
	_, err := mysql.Select(&r, sql, teacherId)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, teacherId)
	exception.CheckMysqlError(err)
	return r, total
}

type ProfessionDetail struct {
	ProfessionId   int
	ProfessionName string
	ProfessionNo   int64
}

type ProfessionUser struct {
	Info  ProfessionDetail
	Users []Users
}

type Users struct {
	Name, Phone, Sex, No string
	Birthday             null.Int
}

func (this *Profession) QueryProfessionUser(ProfessionId, no int, name string) (ProfessionUser, int64) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "select id as ProfessionId,name as ProfessionName,no as ProfessionNo from Profession where id = ?"
	usql := "select name,phone,birthday,sex,professionno as no from User where professionId = ? and `group` = 0 "
	csql := "select count(*) from User where professionId = ? and `group` = 0 "
	if name != "" {
		name = "'%" + name + "%'"
		usql += " and name like " + name
		csql += " and name like " + name
	}
	if no > 0 {
		usql += " and professionno = " + strconv.Itoa(no)
		csql += " and professionno = " + strconv.Itoa(no)
	}
	r := ProfessionUser{}
	err := mysql.SelectOne(&r.Info, sql, ProfessionId)
	exception.CheckMysqlError(err)
	_, err = mysql.Select(&r.Users, usql, ProfessionId)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, ProfessionId)
	exception.CheckMysqlError(err)
	return r, total
}
