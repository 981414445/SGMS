package face

import (
	"github.com/guregu/null"
)

const (
	USER_GROUP_STUDENT = 0
	USER_GROUP_TEACHER = 1
	USER_GROUP_ADMIN   = 99
)

type UserSigninParam struct {
	Key, Password string
}

type User struct {
	Id, Group int
}

type UserQueryParam struct {
	PageParam
	Id, ProfessionId int
	Name             string
}

type UserBasic struct {
	Id, Ct, Group, Sex int
	Name, Password     string
	Phone, Birthday    null.String
	ProfessionId       null.Int
	ProfessionNo       null.Int
}

type UserUpdateParam struct {
	Id, Group, ProfessionId, Sex int
	Name, Phone, Password        string
	Birthday                     null.Int
}

type UserAddParam struct {
	Name, Phone              string
	Group, ProfessionId, Sex int
	Birthday                 null.Int
}
