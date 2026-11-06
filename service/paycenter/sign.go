package paycenter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

const NotifyTimeWindowSeconds int64 = 300

// VerifyNotifySignature 用原始 body 做 HMAC-SHA256，与通知头里的小写或大写 hex 比较。
// 签名长度不对或不是合法 hex 时返回 false，不 panic。
func VerifyNotifySignature(rawBody []byte, notifyKey, signature string) bool {
	mac := hmac.New(sha256.New, []byte(notifyKey))
	mac.Write(rawBody)
	expected := mac.Sum(nil)

	sig, err := hex.DecodeString(strings.TrimSpace(signature))
	if err != nil {
		return false
	}
	return hmac.Equal(expected, sig)
}

// TimestampWithinWindow 要求请求头和 body 的时间戳一致，且与当前时间相差不超过窗口。
func TimestampWithinWindow(headerTS, bodyTS, now, windowSeconds int64) bool {
	if headerTS == 0 || headerTS != bodyTS {
		return false
	}
	if windowSeconds < 0 {
		windowSeconds = 0
	}
	delta := now - bodyTS
	if delta < 0 {
		delta = -delta
	}
	return delta <= windowSeconds
}
