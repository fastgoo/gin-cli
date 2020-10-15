package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

//md5加密
func MD5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	return hex.EncodeToString(ctx.Sum(nil))
}

//获取随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//获取随机数字
func GetRandom(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(end-start+1) + start
	return num
}
