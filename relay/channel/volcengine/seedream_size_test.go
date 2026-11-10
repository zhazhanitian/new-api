package volcengine

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/dto"
)

func TestSeedreamPixelSize(t *testing.T) {
	cases := []struct {
		model      string
		resolution string
		ratio      string
		want       string
	}{
		{"doubao-seedream-5-0-pro-260628", "1K", "16:9", "1424x800"},
		{"doubao-seedream-5.0-pro", "1.5K", "16:9", "2048x1152"},
		{"doubao-seedream-5.0-flash", "2K", "4:3", "2368x1776"},
		{"doubao-seedream-5-0-flash-260915", "2K", "21:9", "3136x1344"},
		{"doubao-seedream-5.0-lite", "2K", "16:9", "2848x1600"},
		{"doubao-seedream-5-0-260128", "3K", "9:16", "2304x4096"},
		{"doubao-seedream-4-5-251128", "4K", "1:1", "4096x4096"},
		{"doubao-seedream-4-0-250828", "1K", "16:9", "1280x720"},
	}
	for _, tt := range cases {
		got, ok := seedreamPixelSize(tt.model, tt.resolution, tt.ratio)
		if !ok || got != tt.want {
			t.Fatalf("%s %s %s = %q ok=%v, want %s", tt.model, tt.resolution, tt.ratio, got, ok, tt.want)
		}
	}
}

func TestConvertToImageGenRequestMapsResolutionAndRatio(t *testing.T) {
	req := dto.ImageRequest{
		Model:  "doubao-seedream-5.0-pro",
		Prompt: "一只猫",
		Extra:  map[string]json.RawMessage{},
	}
	req.Extra["resolution"] = json.RawMessage(`"1K"`)
	req.Extra["ratio"] = json.RawMessage(`"16:9"`)
	got := ConvertToImageGenRequest(req)
	if got.Size != "1424x800" {
		t.Fatalf("size = %q", got.Size)
	}
	if strings.Contains(got.Prompt, "宽高比") {
		t.Fatalf("prompt = %q", got.Prompt)
	}
}

func TestConvertToImageGenRequestKeepsExplicitSize(t *testing.T) {
	req := dto.ImageRequest{Model: "doubao-seedream-5.0-lite", Prompt: "一只猫", Size: "2048x2048"}
	got := ConvertToImageGenRequest(req)
	if got.Size != "2048x2048" || got.Prompt != "一只猫" {
		t.Fatalf("size=%q prompt=%q", got.Size, got.Prompt)
	}
}

func TestConvertToImageGenRequestDefaultsRatioToSquare(t *testing.T) {
	req := dto.ImageRequest{
		Model:  "doubao-seedream-5.0-pro",
		Prompt: "一只猫",
		Extra:  map[string]json.RawMessage{"resolution": json.RawMessage(`"2K"`)},
	}
	got := ConvertToImageGenRequest(req)
	if got.Size != "2048x2048" || strings.Contains(got.Prompt, "宽高比") {
		t.Fatalf("size=%q prompt=%q", got.Size, got.Prompt)
	}
}

func TestConvertToImageGenRequestDefaultsResolutionAndRatio(t *testing.T) {
	req := dto.ImageRequest{Model: "doubao-seedream-5.0-lite", Prompt: "一只猫"}
	got := ConvertToImageGenRequest(req)
	if got.Size != "2048x2048" || got.Prompt != "一只猫" {
		t.Fatalf("size=%q prompt=%q", got.Size, got.Prompt)
	}
}

func TestConvertToImageGenRequestDefaultsResolutionWhenOnlyRatio(t *testing.T) {
	req := dto.ImageRequest{
		Model:  "doubao-seedream-4-5-251128",
		Prompt: "一只猫",
		Extra:  map[string]json.RawMessage{"ratio": json.RawMessage(`"16:9"`)},
	}
	got := ConvertToImageGenRequest(req)
	if got.Size != "2848x1600" || strings.Contains(got.Prompt, "宽高比") {
		t.Fatalf("size=%q prompt=%q", got.Size, got.Prompt)
	}
}
