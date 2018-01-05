package manager

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/face"
	"SGMS/domain/table"
	"SGMS/domain/util"
)

type ProfessionCourse struct {
}

func (this *ProfessionCourse) Query(param face.ProfessionCourseQueryParam) ([]face.ProfessionCourseBasic, int64) {
	sql := "select pc.id,pc.ProfessionId,pc.CourseId,p.name as ProfessionName,c.name as CourseName from ProfessionCourse pc left join Profession p on pc.professionId = p.id left join Course c on pc.courseId = c.id"
	csql := "select count(*) from ProfessionCourse pc left join Profession p on pc.professionId = p.id left join Course c on pc.courseId = c.id"
	wsql := " where 1=1 "
	if param.ProfessionId > 0 {
		wsql += " and pc.professionId = :Id "
	}
	sql += wsql + " order by pc.ct desc " + param.Limit()
	csql += wsql
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	r := []face.ProfessionCourseBasic{}
	_, err := mysql.Select(&r, sql, param)
	exception.CheckMysqlError(err)
	total, err := mysql.SelectInt(csql, param)
	exception.CheckMysqlError(err)
	return r, total
}

func (this *ProfessionCourse) Add(param face.ProfessionCourseInsertParam) {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	p := table.ProfessionCourse{}
	p.CourseId = param.CourseId
	p.ProfessionId = param.ProfessionId
	p.Ct = util.Now()
	mysql.AddTable(p).SetKeys(true, "Id")
	err := mysql.Insert(&p)
	exception.CheckMysqlError(err)
}

func (this *ProfessionCourse) Del(id int) {
	sql := "delete from ProfessionCourse where id = ?"
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	_, err := mysql.Exec(sql, id)
	exception.CheckMysqlError(err)
}
