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
}

type Profession struct {
	Id        int
	Name      int
	TeacherId int
	Ct        int
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
	address   sql.NullString
}

type CourseUser struct {
	Id       int
	Uid      int
	CourseId int
	Ct       int
	Grade    sql.NullInt64
}

type ProfessionUser struct {
	Id           int
	CourseId     int
	ProfessionId int
	Ct           int
}
