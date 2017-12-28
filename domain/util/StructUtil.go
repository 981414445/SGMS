package util

import (
	"fmt"
	"reflect"
	"strconv"
)

var structFields = make(map[string][]string)

func GetFieldNames(obj interface{}) []string {
	typ := reflect.Indirect(reflect.ValueOf(obj))
	structName := typ.Type().PkgPath() + "/" + typ.Type().Name()
	fs, ok := structFields[structName]
	if ok {
		return fs
	}
	fs = make([]string, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		fs[i] = typ.Type().Field(i).Name
	}
	structFields[structName] = fs
	return fs
}

func MysqlKeys(ks []string) string {
	r := ""
	for i, v := range ks {
		r += "`" + v + "`"
		if i < len(ks)-1 {
			r += ","
		}
	}
	return r
}
func MysqlColonKeys(ks []string) string {
	r := ""
	for i, v := range ks {
		r += ":" + v
		if i < len(ks)-1 {
			r += ","
		}
	}
	return r
}
func MysqlKvs(ks []string) string {
	r := ""
	for i, v := range ks {
		r += "`" + v + "`:" + v
		if i < len(ks)-1 {
			r += ","
		}
	}
	return r
}

func GetPointer(p interface{}) int64 {
	a, _ := strconv.ParseInt(fmt.Sprintf("%p", p), 0, 64)
	return a
}
