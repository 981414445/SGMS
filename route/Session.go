package route

import (
	"SGMS/domain/face"

	sessions "github.com/kataras/go-sessions"
)

const (
	SESSION_USER          = "user"
	SESSION_CAPTCHA       = "captcha"
	SESSION_VERIFYCODE    = "verifycode"
	SESSION_VERIFYCODE_ST = "verifycodesendtime"
	SESSION_AUTO_SIGNIN   = "astoken"
	SESSION_PHONE         = "phone"
	SESSION_EMAIL         = "email"
	SESSION_ERROR         = "error"
)

func SessionGetInt(session sessions.Session, name string) int {
	if i, ok := session.Get(name).(int); ok {
		return i
	}
	return -1
}

func SessionGetInt64(session sessions.Session, name string) int64 {
	if i, ok := session.Get(name).(int64); ok {
		return i
	}
	return -1
}
func SessionSetUser(session sessions.Session, u *face.User) {
	session.Set(SESSION_USER, u)
}

// 从Session获取用户信息
func SessionGetUser(session sessions.Session) *face.User {
	u := session.Get(SESSION_USER)
	if nil == u {
		return nil
	}
	if r, ok := u.(*face.User); ok {
		return r
	} else {
		return nil
	}
}

// 从Session获取用户Id
func SessionGetUserId(session sessions.Session) int {
	user := SessionGetUser(session)
	if nil == user {
		return 0
	}
	return user.Id
}

// 清除用户Session
func SessionClearUser(session sessions.Session) {
	session.Delete(SESSION_USER)
}

// Session设置错误
func SessionSetError(session sessions.Session, errors map[string]string) {
	session.Set(SESSION_ERROR, errors)
}

func SessionGetError(session sessions.Session) map[string]string {
	errs := session.Get(SESSION_ERROR)
	if nil == errs {
		return nil
	}
	emap, ok := errs.(map[string]string)
	if !ok || len(emap) <= 0 {
		return nil
	}
	session.Delete(SESSION_ERROR)
	return emap
}
