package manager

import (
	"SGMS/domain/util"
	"gopkg.in/gorp.v1"
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
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

func (this *Course) Add(param face.CourseInsertParam) {
	r := table.Course{}
	r.Ct = util.Now()
	r.EndTime = param.EndTime
	r.Limit = param.Limit
	r.Name = param.Name
	r.StartTime = param.StartTime
	r.Status = param.Status
	r.TeacherId = param.TeacherId
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	mysql.AddTable(r).SetKeys(true,"Id")
	_,err := mysql.Insert(&r)
	exception.CheckMysqlError(err)
}

func (this *Course) Update(param face.CourseUpdateParam){

}

func (this *Course) fetch(mysql *gorp.DbMap,id int) *table.Course{
	r := []table.Course{}
	return &r[0]
}

func (this *Course) Del(id int) {

}