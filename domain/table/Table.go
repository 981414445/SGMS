package table

import "database/sql"

//活动
type Activity struct {
	Id          int
	Name        string
	Logo1       string
	Logo2       string
	Logo3       string
	Description string
	//报名开始日期
	SignupStartTime int
	//报名结束日期
	SignupEndTime int
	//循环比赛开始日期
	StartTime int
	//循环比赛结束日期
	EndTime int
	Ct      int
	//0：未开始,1：已开始，2：已完成
	Status int
	//0：体验课，1：专项训练课，2：年级课，100：系统训练
	Type int
	//循环赛报名用户上限度。0：不限
	UserLimit int
	//单位：分
	Price     int
	LowerDuan int
	UpperDuan int
	//1: "死活",2: "布局",3: "定式",4: "中盘",5: "官子"
	Category int
	//负责活动的主要老师
	TeacherId int
	//是否上架
	Online bool
	//活动页排名权重
	Rank      int
	ViewCount int
	//报名人数
	SignupCount int
	//是否允许报名
	EnableSignup bool
	//剩余人数
	Remainder int
	//微信群二维码图片
	WxGroupQr    sql.NullString
	IntroVideoId sql.NullInt64
	//年级课版本 ，1 ： 年级课1.0模式： 2:年级课2.0模式（双师课堂），
	ActivityVersion int
	//能否使用代金券
	CanUseCoupon bool
}
type ActivityAuditCard struct {
	Id         int
	ActivityId int
	StartTime  int
	EndTime    int
	Price      int
	Ct         int
	Uid        int
}
type ActivityConfig struct {
	Id         int
	ActivityId int
	//校验级别 ， 1:上传证书即可，100:强制全部校验
	LimitLevel int
	MinAge     sql.NullInt64
	MaxAge     sql.NullInt64
	MinDuan    sql.NullInt64
	MaxDuan    sql.NullInt64
	Ct         int
}

//活动限时打折
type ActivityDiscount struct {
	Id         int
	ActivityId int
	//折扣，如80折
	Discount  int
	Ct        int
	StartTime int
	EndTime   int
	//打折信息
	Info sql.NullString
}

//课堂2.0分组
type ActivityGroup struct {
	Id         int
	Name       sql.NullString
	Ct         int
	TeacherId  sql.NullInt64
	ActivityId int
	SeasonId   int
	GroupOrder int
}

//课堂2.0分组学生
type ActivityGroupUser struct {
	Id         int
	ActivityId int
	SeasonId   int
	GroupId    int
	Uid        int
	Ct         int
}
type ActivityItem struct {
	Id         int
	Name       string
	ActivityId int
	SeasonId   int
	//0：课程，1：考试
	Type int
	//训练项开始时间,小于这个时间不能开始
	StartTime int
	//分
	Duration int
	Ct       int
	//0:未开始，1,正在行，2：已结束
	Status int
	//0:直播，1：录播
	VideoType int
	//课堂Id
	CourseId sql.NullInt64
	//预习试卷id
	PreparePaperId sql.NullInt64
	//课后作业试卷
	PractisePaperId sql.NullInt64
	//考试试卷
	ExamPaperId   sql.NullInt64
	RecordVideoId sql.NullInt64
	//大纲word,文件路径
	Doc sql.NullString
	//大纲文件名
	DocFileName sql.NullString
	//大纲
	Outline sql.NullString
}
type ActivityItemUser struct {
	Id             int
	ActivityId     int
	ActivityItemId int
	Uid            int
	Ct             int
	//0：未开始，1：正在进行，2：已通过
	Status int
}
type ActivityRecordVideoBaidPan struct {
	Id        int
	Url       string
	Password  string
	StartTime int
	Ct        int
}
type ActivityScore struct {
	Id         int
	Uid        int
	ActivityId int
	//胜一局得2分，负得0分
	Score int
	Ct    int
}
type ActivitySeason struct {
	Id         int
	ActivityId int
	StartTime  int
	EndTime    int
	Ct         int
	//赛季名字
	Name string
	//单位分
	Price int
	//是否允许报名，1：允许
	EnableSignup bool
	Status       int
}
type ActivitySeasonBaiduPan struct {
	Id         int
	ActivityId int
	SeasonId   int
	Url        string
	Ct         int
}
type ActivitySeasonUser struct {
	Id         int
	ActivityId int
	SeasonId   int
	Uid        int
	Ct         int
	//用户赛季状态，0：未开始，1：正在进行， 2：已完成
	Status int
	//初始序号
	No int
	//操作人
	Creator sql.NullInt64
}
type ActivityUserExchange struct {
	Id            int
	Uid           int
	SrcActivityId int
	DstActivityId int
	Ct            int
	OrderId       sql.NullInt64
	//操作人
	Creator int
}
type AdminGroupMenu struct {
	Id     int
	Group  int
	MenuId int
	Ct     int
	//get.add.update.del , 位操作
	Mode int
}
type AdminMenu struct {
	Id    int
	Name  string
	Url   string
	Ct    int
	SetId int
	No    int
	//是否显示在左侧
	Show    bool
	UIClass sql.NullString
}
type AdminMenuSet struct {
	Id   int
	Name string
	Ct   int
	//排序
	No      int
	UIClass sql.NullString
}
type AgainstPlan struct {
	Id       int
	BlackUid int
	WhiteUid sql.NullInt64
	GameId   int
	Ct       int
	//对局状态，0：未开始，1：已完成
	Status int
	//0:普通对局，1：指导棋
	Type int
	//创建对局的来源：0：系统训练，
	Src int
}

//大区域，粒度最大的地域
type AreaBig struct {
	Id      int
	Name    string
	Ct      int
	Creator sql.NullInt64
}
type AutoLoginUser struct {
	Id             int
	Uid            sql.NullInt64
	AutoLoginToken sql.NullString
	Ip             sql.NullString
	//到期时间
	Expire sql.NullInt64
	Ct     sql.NullInt64
}
type CapacityItemQuizUser struct {
	Id         int
	ItemId     int
	Uid        int
	QuizId     int
	Answer     sql.NullString
	AnswerSgf  sql.NullString
	Ct         int
	Status     int
	IsRight    bool
	ResultType sql.NullInt64
}
type CapacityItemUser struct {
	Id  int
	Uid int
	//1:计算，2：布局，3：定式，4：中盘，5:官子
	Category int
	Ct       int
	Duan     int
	//0:初始状态，1：正在测评，2：已生效测评，3，测评已过期，4，测评已失败（中断）
	Status    int
	StartTime int
	EndTime   int
	//秒
	Duration int
}
type CapacityUser struct {
	Id      int
	Uid     sql.NullInt64
	Calc    sql.NullInt64
	Layout  sql.NullInt64
	Pattern sql.NullInt64
	Mid     sql.NullInt64
	Yose    sql.NullInt64
	Ct      int
	Duan    sql.NullFloat64
}
type Config struct {
	Id      int
	Key     string
	Value   string
	Comment sql.NullString
}
type Coupon struct {
	Id   int
	Name string
	//1:老带新活动 ，2:码券，3:专属代金券 ，
	Type      int
	Intro     sql.NullString
	StartTime sql.NullInt64
	EndTime   sql.NullInt64
	Ct        int
	Creator   int
	//代金券发放上限,0:不限
	Limit int
	//单位：分
	Price int
}
type CouponActivityIntro struct {
	Id int
	//介绍人
	IntroUid   int
	ActivityId int
	//新用户，被介绍的人
	Uid int
	Ct  int
	//分
	Price int
	//关联的订单
	OrderId int
}
type CouponActivityLimit struct {
	Id         int
	CouponId   int
	ActivityId int
	Ct         int
}
type CouponCode struct {
	Id int
	Ct int
	//使用开始时间
	StartTime int
	//使用结束时间，过期时间
	EndTime int
	//分
	Price int
	//0:未使用，1:已使用
	Status int
	//代金券码
	Code string
}
type CouponUser struct {
	Id  int
	Uid int
	Ct  int
	//0:未使用，1：已使用，2，已过期
	Status int
	Price  int
	//代金券开始使用日期
	StartTime int
	//代金券过期日期
	EndTime int
	//0:手动添加， 1：老带新， 2:代金券码
	Type int
	//操作人
	Creator sql.NullInt64
	//老带新活动介绍id
	CouponActivityIntroId sql.NullInt64
	//代金券id
	CouponId sql.NullInt64
	//代金券码id
	CouponCodeId sql.NullInt64
	//0:未读，1:已读
	Readed sql.NullBool
}
type Course struct {
	Id        int
	Name      string
	TeacherId int
	Ct        int
	//讲课状态：
	//0：未开始
	//1：正在讲课
	//2：讲课暂停
	//3：讲课结束
	//4：课程已取消
	Status int
	//上课开始时间
	StartTime int
	//上课时长，45分钟
	Duration int
	//0:公开课，
	//1：系统训练小班课，1-1班课
	//2：系统训练大班课，1-n班课
	//3：专项训练课，1-A课
	Type int
	//0:学生老师听课卡，1:所有已注册的人
	ImVisibility int
	//进入课堂权限，0:活动相关的人能看，1：所有注册的人能看
	Visibility int
	//听课人数上限
	UserLimit int
	//所属活动Id
	ActivityId sql.NullInt64
	//所属活动轮次Id
	RoundId       sql.NullInt64
	RecordVideoId sql.NullInt64
	//ActivityItemId
	ActivityItemId sql.NullInt64
	SeasonId       sql.NullInt64
}
type CourseMessage struct {
	Id       int
	CourseId int
	Message  string
	Uid      int
	Ct       int
}
type CourseQuiz struct {
	Id         int
	CourseId   int
	Sgf        sql.NullString
	Ct         int
	RightCount sql.NullInt64
	FalseCount sql.NullInt64
	//放弃答题数
	NullCount sql.NullInt64
	//题目截图
	SgfImg sql.NullString
	//正确答案
	Answer sql.NullString
}
type CourseQuizUser struct {
	Id           int
	CourseQuizId int
	Uid          int
	Ct           int
	Answer       string
	IsRight      sql.NullBool
	//答题用时
	UseTime sql.NullInt64
}
type CourseRp struct {
	Id       int
	CourseId int
	GroupId  int
	Uid      sql.NullInt64
	Ct       int
	//1-99:punish ,100-: prize
	Action int
}
type CourseSgf struct {
	Id       int
	CourseId int
	//棋谱名称
	Name string
	//排序
	No int
	Ct int
	//学生对弈的棋谱，gameId,和sgf必选其一
	GameId sql.NullInt64
	Sgf    sql.NullString
}
type CourseUserStat struct {
	Id       int
	CourseId int
	//排名0,1,2,3,按班级排名
	Rank       int
	Uid        int
	Ct         int
	RightCount int
	FalseCount int
	//放弃答题数
	NullCount int
	//是否是mvp
	IsMvp bool
	//班级Id
	GroupId sql.NullInt64
	//答题累计用时
	UseTimeTotal int
}
type CourseUserTime struct {
	Id        int
	CourseId  int
	Uid       int
	StartTime int
	EndTime   int
	Ct        int
}
type Evaluation struct {
	Id        int
	Uid       int
	StartTime sql.NullInt64
	EndTime   sql.NullInt64
	Ct        int
	//测评后的段位
	Duan sql.NullInt64
	//0 : 未开始，1:正在进行，2:已正常完成测评
	Status int
	//用户预设的段
	PresetDuan int
	//大区域id
	AreaBigId sql.NullInt64
}
type EvaluationQuiz struct {
	Id           int
	EvaluationId int
	Uid          int
	QuizId       int
	Ct           int
	IsRight      bool
	//题号，从0开始
	No        int
	AnswerSgf sql.NullString
	Answer    sql.NullString
	//题目段位
	Duan int
}
type EvaluationRemark struct {
	Id int
	//段位
	Duan int
	//测试段=预设段
	RemarkEq sql.NullString
	//测试段>预设段
	RemarkGt sql.NullString
	//测试段<预设段
	RemarkLt            sql.NullString
	Ct                  int
	RecommendActivityId sql.NullInt64
}
type Feedback struct {
	Id      int
	Uid     int
	Content string
	Ct      int
}
type Game struct {
	Id       int
	BlackUid int
	WhiteUid int
	//0：没有结果，1:黑胜 2:白胜,3:平
	Result    sql.NullInt64
	GnuResult sql.NullString
	Ct        int
	//预计开始时间
	PlanStartTime int
	StartTime     sql.NullInt64
	EndTime       sql.NullInt64
	//0:未开始，1：正在进行，2：已结束
	Status int
	//黑等级分
	BlackLevelScore sql.NullInt64
	//白等级分
	WhiteLevelScore sql.NullInt64
	//棋谱截图
	Capture1 sql.NullString
	Capture2 sql.NullString
	//黑剩余大时间 秒
	BlackPlentyTime sql.NullInt64
	//黑剩余读秒次数
	BlackPeriodCount sql.NullInt64
	//白剩余大时间，秒
	WhitePlentyTime sql.NullInt64
	//白剩余读秒次数
	WhitePeriodCount sql.NullInt64
	GameRuleId       int
	IsAi             bool
}
type GameAiSetting struct {
	Id       int
	GameId   int
	Uid      int
	Ct       int
	Komi     float32
	Handicap int
	//0:chinese 1:japanese
	Rule int
	//电脑的段位
	Duan int
	//1:黑，2:白 ， 用户颜色选择
	UserColor    int
	PeriodLength int
	//黑定时器，大时间， 秒
	BlackPlentyTime int
	//黑定时器读秒次数
	BlackPeriodCount int
	WhitePlentyTime  int
	WhitePeriodCount int
	//用户上传的残局
	UploadSgfId sql.NullInt64
	//谁先，1:黑，2:白
	WhoFirst int
}
type GameAiUploadSgf struct {
	Id  int
	Sgf string
	Ct  int
}
type GameMessage struct {
	Id      int
	Uid     int
	GameId  int
	Message string
	Ct      int
}
type GameRule struct {
	Id int
	//规则：0:chinese 1:japanese
	Rule      int
	Komi      float32
	BoardSize int
	//0：正常对局，>0 :让子，白先
	Handicap int
	//是否是指导棋，0：否，1：是
	GuideGame    bool
	PlentyTime   int
	PeriodLength int
	PeriodCount  int
	Ct           int
	Name         string
	//谁先，1:黑，2:白
	WhoFirst int
	//离线时间
	OfflineTime int
}
type GameSgf struct {
	Id     int
	Sgf    string
	Ct     int
	GameId int
}
type HistoryAdminActivitySeasonUser struct {
	Id      int
	AdminId int
	//1:添加，2:删除
	Action     int
	ActivityId int
	SeasonId   int
	Uid        int
	Ct         int
}
type HistoryAdminOperation struct {
	Id  int
	Uid int
	//用户的动作
	Action int
	Ct     int
	Param  sql.NullString
}
type HistoryAdminUserValidation struct {
	Id      int
	AdminId int
	//认证状态：0 ：未认证， 1:认证中， 2:已认证，3:认证失败
	Action int
	Uid    int
	Ct     int
}
type HistoryUserSignin struct {
	Id   int
	Uid  int
	Ip   string
	Ct   int
	Date int
}
type IndexActivity struct {
	Id         int
	ActivityId int
	Online     int
	Rank       int
	Ct         int
}
type IndexAd struct {
	Id     int
	Src    string
	Img    string
	Ct     int
	Online int
	Rank   int
	//0:专项广告，1：系统训练广告
	Position int
}
type IndexNews struct {
	Id     int
	Ct     int
	Online int
	Rank   int
	Src    string
	Title  string
}
type IndexSlider struct {
	Id    int
	Title string
	Img   string
	Src   string
	//0:下架不显示，1：上线
	Online int
	//排序,从0开始
	Rank int
	Ct   int
}
type IndexVideo struct {
	Id            int
	PublicVideoId int
	Online        int
	Rank          int
	Ct            int
	//视频类型 0:直播，1:视频，2:视频集
	Type int
	//是否为大图，0:小图，1:大图
	Isbig int
	Title string
	Img   string
	//直播的课程的id
	CourseId sql.NullInt64
	//将文章切换到直播课堂的时间
	ShowTime sql.NullInt64
	//没到切换时间时的展示文章的id
	Article sql.NullInt64
}
type Message struct {
	Id       int
	Sender   int
	Receiver int
	Message  string
	Ct       int
	//1:纯文本消息，2：图片，3：图文混排
	Type   int
	Readed int
}
type MessageSession struct {
	Id int
	//最后一条消息id
	LastMessageId int
	Sender        int
	Receiver      int
	Ct            int
	Ut            int
}
type Order struct {
	Id  int
	Uid int
	//0:待支付，1：已支付
	Status int
	//分
	TotalFee int
	//订单过期时间
	Expire int
	Ct     int
	//订单UUID，outTradeNo
	Uuid    string
	PayTime sql.NullInt64
	//交易ID
	TransactionId sql.NullString
	PayType       int
}
type OrderCommodity struct {
	Id      int
	OrderId int
	//商品Id
	CommodityId int
	//商品类型
	CommodityType int
	Uid           int
	Ct            int
}
type OrderCoupon struct {
	Id           int
	OrderId      int
	CouponUserId int
	Ct           int
	//0:待支付，1:已支付，2:已过期
	Status int
}
type Paper struct {
	Id        int
	Name      string
	Creator   int
	Ct        int
	StartTime int
	//分钟
	Duration int
	//及格分
	PassScore int
	//总分
	TotalScore int
	//0:系统训练，1：专项训练预习，2：专项作业，3，专项考试
	Src        int
	ActivityId sql.NullInt64
	//roundId or itemId
	RoundId sql.NullInt64
	//试卷是否以完成，老师是否创建好了试卷，填充题目
	Completed bool
	//activityItemId
	ActivityItemId sql.NullInt64
	SeasonId       sql.NullInt64
}

//试卷分类表
type PaperClass struct {
	Id      int
	Name    string
	Ct      int
	Creator int
	Duan    int
	//1:每日一套推送试卷
	Code sql.NullInt64
}
type PaperClassRel struct {
	Id           int
	PaperId      int
	PaperClassId int
	Ct           int
}
type PaperQuiz struct {
	Id      int
	PaperId int
	QuizId  int
	Score   int
	//第几题,从0开始
	No int
	Ct int
	//题干
	Name sql.NullString
}
type PaperUser struct {
	Id        int
	Uid       int
	PaperId   int
	Ct        int
	StartTime int
	//0:没做，1：正在做，2：已提交，3：已批改，4：待重做,5：正在重做，6；已重做提交，7：已批改重做，8：已结束
	Status int
	//不及格次数
	FailCount int
	Score     int
	//老师是否已阅时间，0：未阅，1：已阅
	Reviewed int
	//做试卷用时，分钟
	DoneTime sql.NullInt64
	//阅卷老师
	TeacherId sql.NullInt64
	//第一次考试分数
	FirstScore sql.NullInt64
}
type PaperUserScore struct {
	Id      int
	Uid     int
	PaperId int
	Score   int
	Ct      int
}

//Ployv分类仅设两级，第一级年，如"2017"，第二级：活动名+ +老师名
type PolyvCataActivity struct {
	Id         int
	ActivityId int
	Parentid   int64
	Cataid     int64
	Ct         int
	//第一级目录是年分，eg.2017 ， 第二级目录名是：活动名 老师名
	Name string
}
type PolyvLiveChannel struct {
	Id        int
	ChannelId int64
	//0:未开始，1:直播中，2:转码中，3:转码成功，4：转码失败
	Status   int
	Ct       int
	Password string
	Stream   sql.NullString
	RtmpUrl  sql.NullString
	M3u8Url  sql.NullString
	//开始直播时间
	StartTime sql.NullInt64
	//结束直播时间
	EndTime sql.NullInt64
	//保利威视视频vid ， 录播视频vid
	PolyvVideoId  sql.NullString
	RecordVideoId sql.NullInt64
	RecordFileUrl sql.NullString
}
type PolyvLiveCourse struct {
	Id       int
	CourseId int
	//PloyvLiveChannelId
	LiveId int
	Ct     int
}
type PublicAchievements struct {
	Id       int
	Url      string
	Ct       int
	FileName sql.NullString
	No       sql.NullInt64
}
type PublicNews struct {
	Id      int
	Title   string
	Content string
	//秒
	Ct        int
	ViewCount int
	Online    int
	//rank越大越在前
	Rank int
	//publish time
	Pt     int
	Author sql.NullString
	//原图
	Img sql.NullString
	//60x60
	Img1 sql.NullString
	//165x120
	Img2    sql.NullString
	Summary sql.NullString
	//是否是原创，0：否，1：是
	Original int
	ClassId  sql.NullInt64
}
type PublicNewsClass struct {
	Id    int
	Title string
	Ct    int
	Rank  int
}
type PublicRecommend struct {
	Id    int
	Title string
	RefId int
	//秒
	Ct        int64
	ViewCount int
	Online    int
	//rank越大越在前
	Rank int
	//publish time
	Pt int
	//原图
	Img sql.NullString
	//60x60
	Img1 sql.NullString
	//165x120
	Img2 sql.NullString
	//0:新闻，1：活动
	Type int
}
type PublicTeacherRes struct {
	Id              int
	PublicTeacherId sql.NullInt64
	Title           sql.NullString
	Img1            sql.NullString
	VideoScript     sql.NullString
	Ct              sql.NullInt64
	//1:老师视频，2：学生，3：家长评价
	Type sql.NullInt64
	Rank sql.NullInt64
}
type PublicTeachers struct {
	Id          int
	Name        string
	Tags        string
	Summary     string
	Description sql.NullString
	Ct          int64
	Img         sql.NullString
	Img1        sql.NullString
	Online      int
	Rank        int
	Duan        string
	Intro       sql.NullString
}
type PublicVideoPlaylist struct {
	Id      int
	Title   string
	Ct      int
	Creator int
	//视频集封面
	Img         sql.NullString
	Online      int
	Description sql.NullString
	Rank        int
	VideoCount  int
}
type PublicVideoPlaylistRel struct {
	Id                    int
	PublicVideoId         int
	PublicVideoPlaylistId int
	Rank                  int
	Ct                    int
	Creator               int
}
type PublicVideos struct {
	Id    int
	Title string
	//如果是直播的话填写的是直播的课程的id
	CourseId  sql.NullInt64
	ViewCount int
	//0:下线，1:上线
	Online int
	Rank   int
	Ct     int
	Img    string
	Img1   string
	//发布时间
	Pt int64
	//0:直播，1：视频
	Type      int
	TeacherId sql.NullInt64
	//视频
	RecordVideoId int
	Intro         sql.NullString
	//将文章切换到直播课堂的时间
	ShowTime sql.NullInt64
	//没到切换时间时的展示文章的id
	ArticleId sql.NullInt64
}
type Quiz struct {
	Id      int
	Name    string
	Creator int
	//题目类型，1：摆图题，2：选点题，3：自动应对题
	Type int
	//段位：[-9,9]，-9=10K,0=1K,1=1D
	Duan int
	Ct   int
	//题目sgf
	QuizSgf string
	//实体截图
	QuizPic sql.NullString
	Answer  sql.NullString
	//1:黑先，2：白先
	WhoPlay int
	//1: "死活",2: "布局",3: "定式",4: "中盘",5: "官子"
	Category int
	//测试目标 1: "计算",2: "判断",3: "攻防",4: "定式",5: "棋型",6: "价值",7: "大局"
	Aim int
	//0:不需要选择，1：净杀/净活 2. 打劫
	ResultType int
	//上传sgf文件名
	FileName sql.NullString
	//试题视频讲解
	VideoId sql.NullInt64
}
type QuizClass struct {
	Id   int
	Name string
	//0:测试题集
	Code    sql.NullInt64
	Ct      int
	Creator int
}
type QuizClassRel struct {
	Id      int
	ClassId int
	QuizId  int
	Ct      int
}
type QuizMistake struct {
	Id         int
	Uid        int
	QuizId     int
	Ct         int
	ActivityId sql.NullInt64
	RoundId    sql.NullInt64
	PaperId    sql.NullInt64
	//错误的知识点大域
	Field1 sql.NullString
	Field2 sql.NullString
	Field3 sql.NullString
	//错误点的具体描述
	FieldDes       sql.NullString
	Raw            sql.NullString
	IsRight        bool
	ActivityItemId sql.NullInt64
	//是否被用户移出错题本
	Deleted bool
}
type QuizUser struct {
	Id        int
	Status    int
	QuizId    int
	Uid       int
	Ct        int
	FailCount int
	//最后一次是否做对，0：不对，1：对
	IsRight        sql.NullBool
	Answer         sql.NullString
	AnswerSgf      sql.NullString
	ResultType     sql.NullInt64
	PaperId        sql.NullInt64
	ActivityId     sql.NullInt64
	RoundId        sql.NullInt64
	ActivityItemId sql.NullInt64
}
type RecommendContent struct {
	Id    int
	Rank  int
	Ct    int
	Type  int
	Url   string
	Img   string
	Title string
}

//视频录像
type RecordVideo struct {
	Id int
	//乐视云脚本
	VideoScript sql.NullString
	//polyv视频id
	PolyvVideoId   sql.NullString
	PolyvCatalogId sql.NullInt64
	Ct             int
	//0：保利威视，1：乐视
	Type int
}
type RecordVideoClass struct {
	Id             int
	Name           sql.NullString
	PolyvCatalogId sql.NullString
	Ct             sql.NullInt64
}

//系统训练轮次
type Round struct {
	Id         int
	ActivityId int
	SeasonId   int
	//0：未开始，1：正在进行，2：已结束
	Status     int
	Ct         int
	RoundOrder int
	StartTime  int
	EndTime    int
}
type RoundAgainstPlan struct {
	Id         int
	ActivityId int
	RoundId    int
	GameId     int
	Ct         int
	ScheduleId int
}
type RoundConfig struct {
	Id int
	//0:安组自动升降，1:手动升降
	GroupMode int
	Ct        int
	//试卷分组，eg "3,4," : 表示1-3做A卷，4-7组做B卷，剩余的做C卷
	PaperGroup sql.NullString
	//默认创建试卷的老师
	PaperCreator sql.NullInt64
	ActivityId   int
	GameRuleId   int
}
type RoundCourse struct {
	Id         int
	ActivityId int
	RoundId    int
	GroupId    int
	CourseId   int
	ScheduleId int
	Ct         int
}
type RoundGroup struct {
	Id         int
	ActivityId int
	RoundId    int
	//组序号，从0开始
	GroupOrder int
	Ct         int
	//负责此组的老师
	TeacherId sql.NullInt64
	SeasonId  int
	VideoId   sql.NullInt64
}
type RoundGroupConfig struct {
	Id int
	//组默认teacher
	TeacherId  int
	GroupOrder int
	Ct         int
	ActivityId int
}
type RoundGroupUser struct {
	Id         int
	ActivityId int
	RoundId    int
	GroupId    int
	Uid        int
	//小组排名，从0开始
	Rank int
	//小组用户初始序号，从0开始
	No int
	Ct int
	//大分，赢一局得2分，输一局得0分 ， 排名1.比大分，2.比小分，3，直胜，4、随机排名
	Score int
	//所赢对手score的和
	SmallScore int
}
type RoundGroupVideo struct {
	Id      int
	GroupId int
	VideoId int
	Ct      int
}
type RoundPaper struct {
	Id         int
	ActivityId int
	RoundId    int
	PaperId    int
	ScheduleId int
	Ct         int
	//试卷的次数，eg.ABC卷 , 0 是A卷
	Rank int
}
type RoundSchedule struct {
	Id        int
	StartTime int
	EndTime   int
	Type      int
	Ct        int
	Param     sql.NullString
	RoundId   int
}

//系统训练分组规则表，分组已这个表为标准
//每次重新分组将会按照这个表进行分组
//每次升降会体现在这个表中
//
type RoundUser struct {
	Id         int
	ActivityId int
	Uid        int
	//名次，升降会互换名次，每6人为一组
	No int
	Ct int
}
type StatDaySystem struct {
	Id   int
	Date int
	//用户注册数
	UserSignupCount int
	//活跃用户数
	UserActiveCount int
	//活动报名数
	ActivitySignupCount int
	//新增订单数
	OrderCount int
}
type User struct {
	Id       int
	RealName string
	Email    sql.NullString
	Phone    sql.NullString
	Password string
	Ct       int
	Duan     float32
	//1:已禁用
	Disabled      bool
	Group         int
	Icon1         string
	Icon2         string
	Icon3         string
	Weixin        sql.NullString
	Qq            sql.NullString
	Gender        sql.NullInt64
	Score         sql.NullInt64
	Birthday      sql.NullString
	LastSiginTime sql.NullInt64
	Pinyin        string
}
type UserGroup struct {
	Id    int
	Uid   int
	Group int
	Ct    int
}
type UserLoginCookie struct {
	Id     int
	Uid    int
	Token  string
	Expire int
	Ct     int
	Ip     sql.NullString
}
type UserToken struct {
	Id  int
	Uid int
	//web socket user token
	WsToken       sql.NullString
	WsTokenExpire sql.NullInt64
	//网页自动登陆token
	WebLoginToken       sql.NullString
	WebLoginTokenExpire sql.NullInt64
	Ct                  int
}
type UserValidation struct {
	Id  int
	Uid int
	//学生真实姓名
	Name string
	Duan sql.NullInt64
	//认证状态：0 ：未认证， 1:认证中， 2:已认证，3:认证失败
	Status int
	//段位证书图片
	DuanImg sql.NullString
	//身份证
	IdImg    sql.NullString
	Birthday sql.NullInt64
	Ct       int
	//1:男，2:女
	Gender sql.NullInt64
}
type WeixinUser struct {
	Id           int
	Uid          sql.NullInt64
	AccessToken  string
	ExpiresIn    int
	RefreshToken string
	Openid       string
	Scope        string
	Unionid      string
	Ct           int
	//0:服务号，1:微信网页登陆
	Src int
}
