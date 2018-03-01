package face

import (
	"github.com/guregu/null"
)

type CourseUserQueryParam struct {
	ProfessionId, Uid, Choose int
}

type CourseUserAddParam struct {
	Uid, CourseId int
}

type CourseUserUpdateParam struct {
	Id    int
	Score null.Int
}
