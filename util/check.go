package util

import (
	"regexp"

	jsoniter "github.com/json-iterator/go"
)

var r = regexp.MustCompile(`^1[3456789]\d{9}$`)

func IsValidPhone(phoneNumber string) bool {
	// 使用正则表达式匹配手机号码
	return r.MatchString(phoneNumber)
}

func IsJSON(s string) bool {
	var js map[string]interface{}
	return jsoniter.Unmarshal([]byte(s), &js) == nil
}
