package face

import (
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
}

type ProfessionUsers struct {
	Uid, Ct         int
	Name            string
	Phone, Birthday null.String
}

type ProfessionInsertParam struct {
	Name          string
	No, TeacherId int
}

type ProfessionUpdateParam struct {
	Name              string
	Id, No, TeacherId int
}
