package face

import (
	"SGMS/domain/table"

	"github.com/guregu/null"
)

type CourseQueryParam struct {
	PageParam
	Name                                  string
	TeacherId, Status, StartTime, EndTime int
}

type CourseDetail struct {
	Course table.Course
	Users  []CourseUserDetail
}

type CourseUserDetail struct {
	PageParam
	Uid, Sex     int
	Name         string
	Phone        null.String
	ProfessionId null.Int
}

type CourseInsertParam struct {
	Name, Address                                string
	TeacherId, Status, StartTime, EndTime, Limit int
}

type CourseUpdateParam struct {
	Name, Address                            string
	Id, TeacherId, StartTime, EndTime, Limit int
}
