package table

import (
	"database/sql"
)

type User struct {
	Id           int
	Name         string
	Phone        sql.NullString
	Ct           int
	Group        int
	Birthday     sql.NullString
	Password     string
	ProfessionId sql.NullInt64
	Sex          int
	ProfessionNo sql.NullInt64
	No           sql.NullInt64
}

type Profession struct {
	Id        int
	Name      string
	TeacherId int
	Ct        int
	No        int
}

type Course struct {
	Id        int
	Name      string
	TeacherId int
	Ct        int
	Status    int
	StartTime int
	EndTime   int
	Limit     int
	Signup    int
	Address   sql.NullString
}

type CourseUser struct {
	Id       int
	Uid      int
	CourseId int
	Ct       int
	Grade    sql.NullInt64
}

type ProfessionCourse struct {
	Id           int
	CourseId     int
	ProfessionId int
	Ct           int
}
