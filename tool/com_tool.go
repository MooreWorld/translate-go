package tool

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

// CheckParams 参数检查
func CheckParams(str ...string) (err error) {
	for _, item := range str {
		if item == "" {
			return fmt.Errorf("str empty")
		}
	}
	return nil
}

// GetMillTimeStampStr 获取毫秒级时间戳
func GetMillTimeStampStr() string {
	t := time.Now().UnixNano() / 1e6
	return strconv.FormatInt(t, 10)
}

// GetMD5 获取 MD5加密数据
func GetMD5(msg string) string {
	data := []byte(msg)
	sum := md5.Sum(data)
	res := fmt.Sprintf("%x", sum)
	return res
}
