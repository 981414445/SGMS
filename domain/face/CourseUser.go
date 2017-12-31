package face

import (
	"github.com/guregu/null"
)

type CourseUserAddParam struct {
	Uid, CourseId int
}

type CourseUserUpdateParam struct {
	Id    int
	Grade null.Int
}
