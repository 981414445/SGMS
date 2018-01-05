package face

type ProfessionCourseQueryParam struct {
	PageParam
	ProfessionId int
}

type ProfessionCourseBasic struct {
	Id, ProfessionId, CourseId, Ct int
	ProfessionName, CourseName     string
}

type ProfessionCourseInsertParam struct {
	ProfessionId, CourseId int
}
