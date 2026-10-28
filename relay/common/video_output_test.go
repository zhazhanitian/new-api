package common

import "testing"

func TestNormalizeVideoResolutionRatioKeepsOuterFields(t *testing.T) {
	req := &TaskSubmitReq{
		Resolution: "2K",
		Ratio:      "9:16",
		Size:       "1920x1080",
		Metadata:   map[string]interface{}{"resolution": "768P", "ratio": "1:1"},
	}
	NormalizeVideoResolutionRatio(req)
	if req.Resolution != "2K" || req.Ratio != "9:16" {
		t.Fatalf("outer fields changed: %+v", req)
	}
}

func TestNormalizeVideoResolutionRatioFromSizeAndMetadata(t *testing.T) {
	req := &TaskSubmitReq{
		Size:     "16:9",
		Metadata: map[string]interface{}{"resolution": "1080P"},
	}
	NormalizeVideoResolutionRatio(req)
	if req.Resolution != "1080P" || req.Ratio != "16:9" {
		t.Fatalf("legacy size ratio not folded: %+v", req)
	}
}

func TestNormalizeVideoResolutionRatioFromPixelSize(t *testing.T) {
	req := &TaskSubmitReq{Size: "1920x1080"}
	NormalizeVideoResolutionRatio(req)
	if req.Resolution != "1080p" || req.Ratio != "16:9" {
		t.Fatalf("pixel size not folded: %+v", req)
	}
}

func TestNormalizeVideoResolutionRatioFromMetadata(t *testing.T) {
	req := &TaskSubmitReq{Metadata: map[string]interface{}{"resolution": "768P", "ratio": "4:3"}}
	NormalizeVideoResolutionRatio(req)
	if req.Resolution != "768P" || req.Ratio != "4:3" {
		t.Fatalf("metadata not folded: %+v", req)
	}
}
