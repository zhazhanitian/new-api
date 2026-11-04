package doubao

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/relay/channel/volcengine"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
)

func TestTaskReqToImageReqKeepsWatermarkFalse(t *testing.T) {
	watermark := false
	req := &relaycommon.TaskSubmitReq{
		Model:        "doubao-seedream-5-0-260128",
		Prompt:       "仙侠大战",
		Size:         "2K",
		OutputFormat: "png",
		N:            1,
		Watermark:    &watermark,
	}
	info := &relaycommon.RelayInfo{
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "doubao-seedream-5-0-260128",
		},
	}

	upstream := volcengine.ConvertToImageGenRequest(taskReqToImageReq(req, info))
	body, err := json.Marshal(upstream)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), `"watermark":false`) {
		t.Fatalf("upstream body dropped watermark: %s", body)
	}
}

func TestTaskReqToImageReqOmitsUnsetWatermark(t *testing.T) {
	req := &relaycommon.TaskSubmitReq{
		Model:        "doubao-seedream-5-0-260128",
		Prompt:       "仙侠大战",
		Size:         "2K",
		OutputFormat: "png",
	}
	info := &relaycommon.RelayInfo{
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "doubao-seedream-5-0-260128",
		},
	}

	upstream := volcengine.ConvertToImageGenRequest(taskReqToImageReq(req, info))
	body, err := json.Marshal(upstream)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(body), "watermark") {
		t.Fatalf("unset watermark should stay omitted: %s", body)
	}
}
