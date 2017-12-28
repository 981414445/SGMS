package user

import (
	"SGMS/domain/db"
	"SGMS/domain/exception"
	"SGMS/domain/util"
	"time"

	"SGMS/domain/table"

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

func (this *UserToken) Fetch(mysql *gorp.DbMap, uid int) *table.UserToken {
	var r []table.UserToken
	_, err := mysql.Select(&r, "select * from UserToken where uid=?", uid)
	exception.CheckMysqlError(err)
	if len(r) > 0 {
		return &(r[0])
	}
	return nil
}
func (this *UserToken) RefreshWsToken(uid int) string {
	mysql := db.InitMysql()
	defer mysql.Db.Close()
	return this.refreshWsToken(mysql, uid)
}
func (this *UserToken) refreshWsToken(mysql *gorp.DbMap, uid int) string {
	userToken := this.Fetch(mysql, uid)
	now := int(time.Now().Unix())
	if nil == userToken {
		expire := now + 24*3600
		token := util.RandomStr32()
		_, err := mysql.Exec("insert UserToken (uid,wsToken,wsTokenExpire,ct) values(?,?,?,?)", uid, token, expire, now)
		exception.CheckMysqlError(err)
		return token
	}
	if userToken.WsTokenExpire.Valid && userToken.WsTokenExpire.Int64 > int64(now) {
		return userToken.WsToken.String
	}
	expire := now + 24*3600
	token := util.RandomStr32()
	_, err := mysql.Exec("update UserToken set wsToken=?,wsTokenExpire=? where uid=?", token, expire, uid)
	exception.CheckMysqlError(err)
	return token
}

func (this *UserToken) hasUser(mysql *gorp.DbMap, uid int) bool {
	count, err := mysql.SelectInt("select count(*) from UserToken where uid=?", uid)
	exception.CheckMysqlError(err)
	return count > 0
}

func (this *UserToken) saveToken(mysql *gorp.DbMap, uid int) {

}
