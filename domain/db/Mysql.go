package db

import (
	"database/sql"
	"encoding/json"
	"SGMS/domain/exception"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

type MysqlTs struct {
	gorp.Transaction
	Con gorp.DbMap
}
type mysqlConfig struct {
	MysqlSGMS string
	Trace        bool
}

var conf = mysqlConfig{}

func init() {
	file, err := os.Open("mysql.json")
	if nil != err {
		file, err = os.Open("../mysql.json")
		if nil != err {
			file, err = os.Open("../../mysql.json")
			if nil != err {
				file, err = os.Open("../../../mysql.json")
			}
		}
	}
	if nil != err {
		panic(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&conf)
}

func InitMysql() *gorp.DbMap {
	//username:password@protocol(address)/dbname?param=value
	db, err := sql.Open("mysql", conf.MysqlSGMS)
	if nil != err {
		panic(err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	if conf.Trace {
		dbmap.TraceOn("[SGMS]", log.New(os.Stdout, "mysql:", log.Lmicroseconds))
	} else {
		dbmap.TraceOff()
	}
	return dbmap
}

func Del(table string, id int) {
	mysql := InitMysql()
	defer mysql.Db.Close()
	_, err := mysql.Exec("delete from `"+table+"` where id=?", id)
	exception.CheckMysqlError(err)
}
