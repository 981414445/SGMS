package face

import (
	"SGMS/domain/table"

	"github.com/guregu/null"
)

type ProfessionQueryParam struct {
	PageParam
	Name          string
	No, TeacherId int
}

type ProfessionDetail struct {
	Id, TeacherId, Ct, No                     int
	ProfessionName, TeacherName, TeacherPhone string
	Users                                     []ProfessionUsers
	Courses                                   []table.Course
}

type ProfessionUsers struct {
	Uid, Ct, No, Sex int
	Name             string
	Phone, Birthday  null.String
}

type ProfessionInsertParam struct {
	Name          string
	No, TeacherId int
}

type ProfessionUpdateParam struct {
	Name              string
	Id, No, TeacherId int
}
