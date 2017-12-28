package face

type User struct {
	Id               int
	Group            int
	Duan             float32
	RealName, Pinyin string
	Phone, Email     *string
}

const (
	USER_GROUP_STUDENT = 0
	USER_GROUP_TEACHER = 1
	USER_GROUP_ADMIN   = 99
)
