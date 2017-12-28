package db

import (
    "fmt"
	"testing"
)

func TestInit(t *testing.T) {
    type User struct {
        Id int
    }
    u:=new(User)
    mysql:=InitMysql()
    defer mysql.Db.Close()
    mysql.SelectOne(u , "select id from User where id=1")
    fmt.Println(u)
}