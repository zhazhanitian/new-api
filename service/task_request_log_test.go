package service

import (
	"strings"
	"testing"
)

func TestSanitizeRequestLogKeepsPromptAndDropsBinary(t *testing.T) {
	raw := []byte(`{
		"model":"doubao-seedance",
		"prompt":"一只猫在跑步",
		"image":"data:image/png;base64,` + strings.Repeat("abcd", 80) + `",
		"images":["https://example.com/a.png","` + strings.Repeat("QUJDRA==", 200) + `"],
		"api_key":"sk-secret",
		"duration":5
	}`)

	got := string(SanitizeRequestLog(raw))
	if !strings.Contains(got, "一只猫在跑步") {
		t.Fatalf("prompt missing: %s", got)
	}
	if !strings.Contains(got, "https://example.com/a.png") {
		t.Fatalf("url missing: %s", got)
	}
	if !strings.Contains(got, `"duration":5`) {
		t.Fatalf("duration missing: %s", got)
	}
	if strings.Contains(got, "sk-secret") || strings.Contains(got, "data:image") {
		t.Fatalf("secret or binary leaked: %s", got)
	}
	if !strings.Contains(got, "omitted binary") || !strings.Contains(got, "[redacted]") {
		t.Fatalf("redaction marker missing: %s", got)
	}
}

func TestSanitizeRequestLogWrapsPlainText(t *testing.T) {
	got := string(SanitizeRequestLog([]byte("not json")))
	if !strings.Contains(got, `"raw":"not json"`) {
		t.Fatalf("got %s", got)
	}
}

func TestSanitizeRequestLogEmpty(t *testing.T) {
	if SanitizeRequestLog(nil) != nil {
		t.Fatal("nil input should stay empty")
	}
	if SanitizeRequestLog([]byte("  ")) != nil {
		t.Fatal("blank input should stay empty")
	}
}
