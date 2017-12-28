package admin

import (
	"SGMS/domain/face/admin"
	"SGMS/route"

	sessions "github.com/kataras/go-sessions"
)

const SESSION_USER_ADMIN = "admin"

func SessionSetAdmin(session sessions.Session, u *admin.Admin) {
	session.Set(route.SESSION_USER, &u.User)
	session.Set(SESSION_USER_ADMIN, u)
}

func SessionGetAdmin(session sessions.Session) *admin.Admin {
	u := session.Get(SESSION_USER_ADMIN)
	if nil == u {
		return nil
	}
	if r, ok := u.(*admin.Admin); ok {
		return r
	} else {
		return nil
	}
}
