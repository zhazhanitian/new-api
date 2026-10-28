package helper

import (
	"strings"

	"github.com/QuantumNous/new-api/common"
	taskcommon "github.com/QuantumNous/new-api/relay/channel/task/taskcommon"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/tidwall/sjson"
)

// PatchVideoBillingBody adjusts a copy of the video request before tiered expressions run.
// Stored expressions and the parsed task request are left unchanged.
func PatchVideoBillingBody(body []byte) []byte {
	if len(body) == 0 {
		return body
	}
	var req relaycommon.TaskSubmitReq
	if err := common.Unmarshal(body, &req); err != nil {
		return body
	}
	items := billingContentRoles(taskcommon.NormalizeVideoContent(req))
	relaycommon.NormalizeVideoResolutionRatio(&req)
	out := append([]byte(nil), body...)
	if req.Resolution != "" {
		patched, err := sjson.SetBytes(out, "resolution", req.Resolution)
		if err != nil {
			return body
		}
		out = patched
	}
	if req.Ratio != "" {
		patched, err := sjson.SetBytes(out, "ratio", req.Ratio)
		if err != nil {
			return body
		}
		out = patched
	}

	if len(items) > 0 {
		payload := make([]billingContentItem, 0, len(items))
		for _, item := range items {
			payload = append(payload, billingContentItem{
				Type: item.Type,
				Role: item.Role,
				URL:  item.URL,
			})
		}
		patched, err := sjson.SetBytes(out, "metadata.content", payload)
		if err != nil {
			return body
		}
		out = patched
	}

	if taskcommon.IsAspectRatioSize(req.Size) {
		resolution := ""
		if req.Metadata != nil {
			resolution, _ = req.Metadata["resolution"].(string)
		}
		resolution = strings.TrimSpace(resolution)
		var err error
		if resolution == "" {
			out, err = sjson.DeleteBytes(out, "size")
		} else {
			out, err = sjson.SetBytes(out, "size", resolution)
		}
		if err != nil {
			return body
		}
	}

	hasFirstFrame := false
	refImages := make([]string, 0)
	for _, item := range items {
		if item.Role == taskcommon.VideoRoleFirstFrame {
			hasFirstFrame = true
		}
		if item.Type == taskcommon.VideoContentImage && item.Role == taskcommon.VideoRoleReferenceImage && item.URL != "" {
			refImages = append(refImages, item.URL)
		}
	}
	if len(refImages) > 0 && !hasFirstFrame {
		patched, err := sjson.SetBytes(out, "metadata.input_usage", "Reference")
		if err != nil {
			return body
		}
		out = patched
	}
	if len(req.Images) == 0 && len(refImages) > 0 {
		patched, err := sjson.SetBytes(out, "images", refImages)
		if err != nil {
			return body
		}
		out = patched
	}
	return out
}

// billingContentRoles 只改计价副本里的角色。图片、视频不写 role 时按参考素材计数。
func billingContentRoles(items []taskcommon.VideoContentItem) []taskcommon.VideoContentItem {
	if len(items) == 0 {
		return items
	}
	out := make([]taskcommon.VideoContentItem, len(items))
	copy(out, items)
	for i := range out {
		if strings.TrimSpace(out[i].Role) != "" {
			continue
		}
		switch out[i].Type {
		case taskcommon.VideoContentImage:
			out[i].Role = taskcommon.VideoRoleReferenceImage
		case taskcommon.VideoContentVideo:
			out[i].Role = taskcommon.VideoRoleReferenceVideo
		}
	}
	return out
}

type billingContentItem struct {
	Type string `json:"type,omitempty"`
	Role string `json:"role,omitempty"`
	URL  string `json:"url,omitempty"`
}
