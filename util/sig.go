package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func SignGinRequest(req *http.Request, secret string) {

	path := req.URL.Path
	raw := req.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	reqBody, _ := io.ReadAll(req.Body)
	req.Body = io.NopCloser(bytes.NewReader(reqBody))

	req.Header.Add("timestamp", fmt.Sprintf("%v", time.Now().Unix()))
	_, sign := GenerateSig2(path, http.MethodPost, req.Header, reqBody, secret)
	req.Header.Add("signature", sign)
}

func GetRawSigContent(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 将参数和值进行拼接，用"&"连接
	var s []string
	for _, k := range keys {
		v := params[k]
		s = append(s, fmt.Sprintf("%s=%s", k, v))
	}
	signStr := strings.Join(s, "&")

	// URL编码
	signStr = url.QueryEscape(signStr)

	return signStr
}

func GetRawHttpSigContent(path, method string, header http.Header, body []byte) string {
	timestamp := header.Get("timestamp")
	params := map[string]string{
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
		"body":      string(body),
	}
	return GetRawSigContent(params)
}

// GenerateSig 目前没有大body，暂不考虑
func GenerateSig(path, method string, header http.Header, body []byte, key string) string {
	timestamp := header.Get("timestamp")
	params := map[string]string{
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
		"body":      string(body),
	}
	_, sign := Sign(params, key)
	return sign
}

func GenerateSig2(path, method string, header http.Header, body []byte, key string) (string, string) {
	timestamp := header.Get("timestamp")
	params := map[string]string{
		"method":    method,
		"path":      path,
		"timestamp": timestamp,
		"body":      string(body),
	}
	return Sign(params, key)
}

// Sign 签名算法
// payload, path, method, timestamp
func Sign(params map[string]string, secret string) (string, string) {
	str := GetRawSigContent(params)
	return str, SignStr(str, secret)
}

func SignStr(signStr string, secret string) string {
	// 拼接密钥，并对字符串进行哈希
	h := md5.New()
	h.Write([]byte(secret + signStr))
	signBytes := h.Sum(nil)

	// 将哈希值转换为16进制字符串
	sign := hex.EncodeToString(signBytes)
	return sign
}
