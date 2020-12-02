package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
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

// 加密密码
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

// 验证密码
func ComparePassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}

// 邮箱正则表达式
func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

// 手机号正则表达式
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 密码安全级别正则表达式
func VerifyPasswordSecurity(pwd string) int {
	if len(pwd) < 6 {
		return 0
	}
	securityLevel := 0
	if b, err := regexp.MatchString(`[0-9]{1}`, pwd); b && err == nil {
		securityLevel++
	}
	if b, err := regexp.MatchString(`[a-z]{1}`, pwd); b && err == nil {
		securityLevel++
	}
	if b, err := regexp.MatchString(`[A-Z]{1}`, pwd); b && err == nil {
		securityLevel++
	}
	if b, err := regexp.MatchString(`[!@#~$%^&*()+|_]{1}`, pwd); b && err == nil {
		securityLevel++
	}
	return securityLevel
}

// 获取指定ip信息
func GetIPInfo(ip string) (ret map[string]interface{}, retErr error) {
	defer func() {
		if recover() != nil {
			retErr = errors.New("ip get fail,http error")
		}
	}()
	resp, err := http.Get("http://api.tianapi.com/txapi/ipquery/index?ip=" + ip + "&key=f2d405f48efac244ccab33c1ebaf7911")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}
	if ret["code"].(float64) != 200 {
		return nil, errors.New("ip '" + ip + "' error")
	}
	if ipInfo := ret["newslist"].([]interface{}); len(ipInfo) > 0 {
		ret = ipInfo[0].(map[string]interface{})
		return
	}
	return nil, errors.New("ip get fail")
}

func FileExt(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return ""
}
