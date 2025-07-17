package util

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"
	"unsafe"

	"github.com/google/uuid"
)

const (
	letterBytes   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// GenerateTraceId 生成 trace id
func GenerateTraceId() (s string) {
	var bu bytes.Buffer
	bu.WriteString(strconv.Itoa(int(time.Now().UnixNano() / 1000)))
	bu.WriteByte('_')
	bu.WriteString(RandStringBytesMask(13))
	s = String(bu.Bytes())
	return
}

// RandStringBytesMask 随机生成字符串
func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return String(b)
}

// 字节数组转为string （性能高于string(b)),编译器会完成内联处理，不会发生逃逸行为
func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func GenerateTraceId2() string {
	return uuid.New().String()
}

var numGenerator *rand.Rand

func init() {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	numGenerator = rand.New(source)
}

func GenerateRandNumber(length int) string {
	// if util.IsLocal() {
	// 	return "301740"
	// }
	// 生成随机数字的字符集合
	charset := "0123456789"

	// 生成随机字符串
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := numGenerator.Intn(len(charset))
		result[i] = charset[randomIndex]
	}

	return string(result)
}
