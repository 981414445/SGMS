package userf

import (
	"SGMS/domain/face"
	"SGMS/domain/user"
)

func NewUser() face.IUser {
	return new(user.User)
}
