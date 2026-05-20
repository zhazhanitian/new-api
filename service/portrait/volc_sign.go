package portrait

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	volcHost    = "open.volcengineapi.com"
	volcService = "ark"
	volcVersion = "2024-01-01"
)

// signedHeaders 参与签名的请求头（小写，已排序）
var signedHeaders = []string{"content-type", "host", "x-date"}

func hmacSHA256(key []byte, data string) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(data))
	return mac.Sum(nil)
}

func sha256Hex(data string) string {
	h := sha256.Sum256([]byte(data))
	return hex.EncodeToString(h[:])
}

// BuildRequest 构造带签名的 HTTP 请求参数。
// action: API 动作名，如 "CreateAssetGroup"。
// bodyBytes: JSON 请求体。
// 返回 (headers map, query string, error)。
func BuildRequest(accessKeyId, secretKey, region, action string, bodyBytes []byte) (headers map[string]string, queryStr string, err error) {
	if accessKeyId == "" || secretKey == "" {
		return nil, "", fmt.Errorf("火山引擎凭据未配置，请在系统设置中配置 VolcPortraitAccessKeyId / VolcPortraitSecretAccessKey")
	}

	now := time.Now().UTC()
	xDate := now.Format("20060102T150405Z")
	date := now.Format("20060102")

	// 构造 Query String（Action 和 Version 必须）
	queryParams := url.Values{}
	queryParams.Set("Action", action)
	queryParams.Set("Version", volcVersion)
	// 按 key 排序（火山要求）
	keys := make([]string, 0, len(queryParams))
	for k := range queryParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, url.QueryEscape(k)+"="+url.QueryEscape(queryParams.Get(k)))
	}
	canonicalQueryStr := strings.Join(parts, "&")

	// 构造规范化请求头
	canonicalHeadersMap := map[string]string{
		"content-type": "application/json",
		"host":         volcHost,
		"x-date":       xDate,
	}
	canonicalHeaderLines := make([]string, 0, len(signedHeaders))
	for _, h := range signedHeaders {
		canonicalHeaderLines = append(canonicalHeaderLines, h+":"+canonicalHeadersMap[h])
	}
	canonicalHeaders := strings.Join(canonicalHeaderLines, "\n") + "\n"
	signedHeadersStr := strings.Join(signedHeaders, ";")

	// 规范化请求体哈希
	bodyHash := sha256Hex(string(bodyBytes))

	// 规范请求字符串
	canonicalRequest := strings.Join([]string{
		"POST",
		"/",
		canonicalQueryStr,
		canonicalHeaders,
		signedHeadersStr,
		bodyHash,
	}, "\n")

	// 凭证范围
	credentialScope := strings.Join([]string{date, region, volcService, "request"}, "/")

	// 待签字符串
	stringToSign := strings.Join([]string{
		"HMAC-SHA256",
		xDate,
		credentialScope,
		sha256Hex(canonicalRequest),
	}, "\n")

	// 推导签名密钥
	signingKey := hmacSHA256(hmacSHA256(hmacSHA256(hmacSHA256(
		[]byte(secretKey), date), region), volcService), "request")

	// 计算签名
	signature := hex.EncodeToString(hmacSHA256(signingKey, stringToSign))

	// Authorization header
	authorization := fmt.Sprintf(
		"HMAC-SHA256 Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		accessKeyId, credentialScope, signedHeadersStr, signature,
	)

	headers = map[string]string{
		"Content-Type":  "application/json",
		"Host":          volcHost,
		"X-Date":        xDate,
		"Authorization": authorization,
	}
	return headers, canonicalQueryStr, nil
}
