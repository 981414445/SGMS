package manager

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
	"database/sql"

	"gopkg.in/gorp.v1"
)

type Course struct {
}

func (this *Course) Query(param face.CourseQueryParam) ([]table.Course, int64) {
	sql := "select * from Course "
	csql := "select count(*) from Course"
	wsql := " where 1=1 "
	if param.Name != "" {
		param.Name = "%" + param.Name + "%"
		wsql += " and name like " + param.Name
	}
	if param.TeacherId > 0 {
		wsql += " and TeacherId = :TeacherId "
	}
	if param.Status > -1 {
		wsql += " and Status = :Status "
	}
	if param.StartTime > 0 {
		wsql += " and StartTime > :StartTime "
	}
	if param.EndTime > 0 {
		wsql += " and EndTime < :EndTime "
	}
	sql += wsql + " order by ct desc " + param.Limit()
	csql += wsql
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := []table.Course{}
	_, err := mysql.Select(&r, sql, param)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, param)
	exception.CheckMysqlError(err)
	return r, total
}

func (this *Course) Get(id int) face.CourseDetail {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "select * from Course where id = ?"
	r := []face.CourseDetail{}
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		usql := "select u.id as Uid,u.sex,u.name,u.phone,u.professionId from CourseUser cu left join User u on cu.uid = u.id where cu.courseId = ?"
		u := []face.CourseUserDetail{}
		_, err = mysql.Select(&u, usql, id)
		exception.CheckMysqlError(err)
		r[0].Users = u
		return r[0]
	}
	return face.CourseDetail{}
}

func (this *Course) Add(param face.CourseInsertParam) {
	r := table.Course{}
	r.Ct = util.Now()
	r.EndTime = param.EndTime
	r.Limit = param.Limit
	r.Name = param.Name
	r.StartTime = param.StartTime
	r.Status = param.Status
	r.TeacherId = param.TeacherId
	r.Signup = 0
	r.Address = sql.NullString{param.Address, true}
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	mysql.AddTable(r).SetKeys(true, "Id")
	err := mysql.Insert(&r)
	exception.CheckMysqlError(err)
}

func (this *Course) Update(param face.CourseUpdateParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := this.fetch(mysql, param.Id)
	r.EndTime = param.EndTime
	r.Limit = param.Limit
	r.Name = param.Name
	r.StartTime = param.StartTime
	r.TeacherId = param.TeacherId
	mysql.AddTable(r).SetKeys(true, "Id")
	_, err := mysql.Update(&r)
	exception.CheckMysqlError(err)
}

func (this *Course) fetch(mysql *gorp.DbMap, id int) *table.Course {
	sql := "selectr * from Course where id = ?"
	r := []table.Course{}
	_, err := mysql.Select(&r, sql, id)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &r[0]
	}
	return nil
}

func (this *Course) Del(id int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "delete from Course where id = ?"
	_, err := mysql.Exec(sql, id)
	exception.CheckMysqlError(err)
}

// 修改课程人数限制
func (this *Course) UpdateLimit(courseId, num int) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	sql := "update Course set limit = ? where id = ?"
	_, err := mysql.Exec(sql, courseId, num)
	exception.CheckMysqlError(err)
}
