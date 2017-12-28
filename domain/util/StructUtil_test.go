package util

import (
	"fmt"
	"testing"
)

type A struct {
	B int
}

func TestSort(t *testing.T) {
	a := []string{"2", "a", "c", "b"}
	SortStrs(a)
	fmt.Println(a)
}

func TestGetFieldNames(t *testing.T) {
	a := &A{1}
	fmt.Println(GetFieldNames(a))
}

func TestMysqlKvs(t *testing.T) {
	fmt.Println(MysqlKvs([]string{"k1", "k2"}))
}

func TestMysqlKeys(t *testing.T) {
	fmt.Println(MysqlKeys([]string{"k1", "k2"}))
}
func TestMysqlColonKeys(t *testing.T) {
	fmt.Println(MysqlColonKeys([]string{"k1", "k2"}))
}

func TestPinyin(t *testing.T) {
	fmt.Println(Pinyin("123嘎嘎-43fdfd"))
}
func TestJoinSqlFields(t *testing.T) {
	fmt.Println(JoinSqlFields("a", "`name` as name1"))
}
func TestDistance(t *testing.T) {
	fmt.Println(SqlNowDistance("ct"))
}
func TestDeHtmlTag(t *testing.T) {
	fmt.Println(DeHtmlTag("he<a>ll</a>o"))
}
