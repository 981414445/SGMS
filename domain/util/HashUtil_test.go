package util

import (
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	if "c4ca4238a0b923820dcc509a6f75849b" != Md5("1") {
		t.Error("Md5(1) not correct")
	}
}

func TestMd5File(t *testing.T) {
	fmt.Println(Md5File("HashUtil.go"))
	fmt.Println(RandomStr32())
	fmt.Println(RandomStr32())
	fmt.Println(RandomStr32())
}

func TestAESEncrypt(t *testing.T) {
	k := Md5("23")
	ct, _ := AESEncryptBase64(k, "(;CA[utf-8]AW[oc][pb][pd]AP[MultiGo:4.4.4]SZ[19]AB[pc]MULTIGOGM[1];W[qc])")
	fmt.Println(ct)
	fmt.Println(AESDecryptBase64(ct, k))
}

func eq(q1, q2 interface{}) bool {
	return q1 == q2
}
func TestEqual(t *testing.T) {
	fmt.Println("eq", eq(1, 1))
}

func TestParseDate(t *testing.T) {
	d, _ := ParseDate("2017-01-01")
	fmt.Println(d.Unix())
}
