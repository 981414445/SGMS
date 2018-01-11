package user

import (
	"SGMS/domain/exception"
	"time"

	"strconv"

	gorp "gopkg.in/gorp.v1"
)

//管理用户的各个token
type UserToken struct {
}

type UidTokenUidQueryParam struct {
	WsToken       string
	WebLoginToken string
}

func (this *UserToken) GetUid(mysql *gorp.DbMap, param UidTokenUidQueryParam) int {
	if "" == param.WebLoginToken && "" == param.WsToken {
		return 0
	}
	ps := make(map[string]interface{})
	ssql := "select uid from UserToken where "
	now := time.Now().Unix()
	if param.WsToken != "" {
		ssql += " wsToken=:WsToken and wsTokenExpire>" + strconv.Itoa(int(now))
		ps["WsToken"] = param.WsToken
	}
	if param.WebLoginToken != "" {
		ssql += " webLoginToken=:WebLoginToken and WebLoginTokenExpire>" + strconv.Itoa(int(now))
		ps["WebLoginToken"] = param.WebLoginToken
	}
	ssql += " limit 1"
	var r []struct{ Uid int }
	_, err := mysql.Select(&r, ssql, ps)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return r[0].Uid
	}
	return 0
}

func (this *UserToken) hasUser(mysql *gorp.DbMap, uid int) bool {
	count, err := mysql.SelectInt("select count(*) from UserToken where uid=?", uid)
	exception.CheckMysqlError(err)
	return count > 0
}

func (this *UserToken) saveToken(mysql *gorp.DbMap, uid int) {

}
