package paycenter

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

func signBody(body, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

func TestVerifyNotifySignature(t *testing.T) {
	body := []byte(`{"event":"payment.success","amount":100}`)
	key := "notify-key"
	signature := signBody(string(body), key)

	if !VerifyNotifySignature(body, key, signature) {
		t.Fatal("expected valid signature")
	}
	if !VerifyNotifySignature(body, key, "  "+signature+"  ") {
		t.Fatal("expected trimmed signature to pass")
	}
	if !VerifyNotifySignature(body, key, strings.ToUpper(signature)) {
		t.Fatal("expected uppercase hex signature to pass")
	}
	if VerifyNotifySignature(body, "other-key", signature) {
		t.Fatal("wrong key must fail")
	}
	mutated := append([]byte{}, body...)
	mutated[0] = 'x'
	if VerifyNotifySignature(mutated, key, signature) {
		t.Fatal("mutated body must fail")
	}
	if VerifyNotifySignature(body, key, "abcd") {
		t.Fatal("short signature must fail")
	}
	if VerifyNotifySignature(body, key, "zz") {
		t.Fatal("invalid hex must fail")
	}
	if VerifyNotifySignature(body, key, "") {
		t.Fatal("empty signature must fail")
	}
}

func TestTimestampWithinWindow(t *testing.T) {
	const now int64 = 1_700_000_000
	if !TimestampWithinWindow(now, now, now, NotifyTimeWindowSeconds) {
		t.Fatal("same timestamp must pass")
	}
	if !TimestampWithinWindow(now-300, now-300, now, NotifyTimeWindowSeconds) {
		t.Fatal("edge of five minutes must pass")
	}
	if TimestampWithinWindow(now-301, now-301, now, NotifyTimeWindowSeconds) {
		t.Fatal("older than five minutes must fail")
	}
	if TimestampWithinWindow(now+301, now+301, now, NotifyTimeWindowSeconds) {
		t.Fatal("future beyond five minutes must fail")
	}
	if TimestampWithinWindow(now, now-10, now, NotifyTimeWindowSeconds) {
		t.Fatal("header and body timestamps must match")
	}
	if TimestampWithinWindow(0, 0, now, NotifyTimeWindowSeconds) {
		t.Fatal("zero timestamp must fail")
	}
}
