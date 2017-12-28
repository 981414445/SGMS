package admin

import "database/sql"

//
type RetentionNum struct {
	Uid      int
	RealName string
}

type TeacherRetention struct {
	Signups1, Signups2 []RetentionNum
	Remain, Loss       []RetentionNum
	RemainRate         float32
}

type TeacherRetentionBasic struct {
	TeacherName string
	// 开始时间段
	StartTimeBegin, StartTimeFinish int
	// 结束时间段
	EndTimeBegin, EndTimeFinish int
	// 活动数目
	// 第一期报名学生名单
	Signups1, Signups2 []string
	// 留存报名学生名单
	Remain []string
	// 未报名学生名单
	UnSignup []string
	// 留存率
	RemainRate float32
}

type TeacherRetentionQueryParam struct {
	TeacherId int
	// 是否去重 0:不去重，1:去重
	IsRepeat int
	// 开始时间段
	StartTimeBegin, StartTimeFinish int
	// 结束时间段
	EndTimeBegin, EndTimeFinish int
}

// 管理员操作历史操作类型记录
const (
	// 插入
	HISTORY_ADMIN_OPERATION_INSERT = 0
	// 删除
	HISTORY_ADMIN_OPERATION_DELETE = 1
	// 修改
	HISTORY_ADMIN_OPERATION_UPDATE = 2
)

// 管理员操作的模块
const (
	// 活动模块
	ACTIVITY_MODULE = 101
	// 课程模块
	COURSE_MODULE = 102
	// 试卷模块
	PAPER_MODULE = 103
	// 新闻模块
	PUBLIC_NEWS_MODULE = 104
	// 教师模块
	PUBLIC_TEACHERS_MODULE = 105
	// 用户认证
	USER_VALIDAITON_MODULE = 106
)

type HistoryAdminOperationSaveParam struct {
	Uid, Action, Ct int
	Param           sql.NullString
}
