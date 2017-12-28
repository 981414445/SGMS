package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func Sha1(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}

func Random6() string {
	return strconv.Itoa(rand.Intn(99999) + 100000)
}
func Md5File(file string) (string, error) {
	f, err := os.Open(file)
	if nil != err {
		return "", err
	}
	defer f.Close()
	return Md5Reader(f), nil
}

func Md5Reader(reader io.Reader) string {
	m := md5.New()
	io.Copy(m, reader)
	return fmt.Sprintf("%x", m.Sum([]byte("")))
}
func RandomStr(n int, ignoreCase bool) string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if true == ignoreCase {
		letters = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
	}
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomStr32() string {
	return RandomStr(32, true)
}
func AESEncryptBase64(str32 string, data string) (string, error) {
	bs, err := AESEncrypt([]byte(str32[0:16]), []byte(str32[16:32]), []byte(data))
	return base64.StdEncoding.EncodeToString(bs), err
}
func AESEncrypt(key []byte, vector []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ecb := cipher.NewCBCEncrypter(block, vector)
	content := PKCS5Padding(data, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)
	return crypted, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func AESDecryptBase64(ciphertext, key32 string) (string, error) {
	c, err := base64.StdEncoding.DecodeString(ciphertext)
	if nil != err {
		return "", err
	}
	bs, err := AESDecrypt(c, []byte(key32))
	return string(bs), err
}

func AESDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key[0:16]) //选择加密算法
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, key[16:32])
	plantText := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plantText, ciphertext)
	plantText = PKCS5Unpadding(plantText)
	return plantText, nil
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
