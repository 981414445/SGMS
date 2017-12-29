package face

type CourseQueryParam struct {
	PageParam
	Name                                  string
	TeacherId, Status, StartTime, EndTime int
}

type CourseInsertParam struct {
	Name, Address                                string
	TeacherId, Status, StartTime, EndTime, Limit int
}

type CourseUpdateParam struct {
	Name, Address                                    string
	Id, TeacherId, Status, StartTime, EndTime, Limit int
}
