package apimart_image

import (
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/model"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestBuildUpstreamBodyMapsFourFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:      "gpt-image-2.6-flare",
		Prompt:     "a cat",
		Ratio:      "16:9",
		Resolution: "2K",
		N:          2,
		Images:     []string{"https://example.com/a.png"},
		Metadata: map[string]interface{}{
			"complexity": "high",
		},
	})
	info := &relaycommon.RelayInfo{
		ChannelMeta:   &relaycommon.ChannelMeta{UpstreamModelName: "gpt-image-2.6-flare"},
		TaskRelayInfo: &relaycommon.TaskRelayInfo{},
	}
	reader, err := (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw := readAll(t, reader)
	require.Equal(t, "16:9", gjson.GetBytes(raw, "size").String())
	require.Equal(t, "2k", gjson.GetBytes(raw, "resolution").String())
	require.Equal(t, "low", gjson.GetBytes(raw, "quality").String())
	require.Equal(t, "high", gjson.GetBytes(raw, "complexity").String())
	require.Equal(t, "https://example.com/a.png", gjson.GetBytes(raw, "image_urls.0").String())

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Prompt: "pixels",
		Size:   "1024x1024",
	})
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, "1024x1024", gjson.GetBytes(raw, "size").String())
	require.False(t, gjson.GetBytes(raw, "resolution").Exists())
}

func TestParseTaskResultImages(t *testing.T) {
	body := []byte(`{
		"code": 200,
		"data": {
			"id": "task_1",
			"status": "completed",
			"result": {"images": [{"url": ["https://cdn/a.png", "https://cdn/b.png"]}]},
			"usage": {"total_tokens": 12}
		}
	}`)
	info, err := parseTaskResult(body)
	require.NoError(t, err)
	require.Equal(t, "SUCCESS", info.Status)
	require.Equal(t, []string{"https://cdn/a.png", "https://cdn/b.png"}, info.Urls)
	require.Equal(t, 2, info.ActualImages)
	require.Equal(t, 12, info.TotalTokens)
}

func TestQuotaMatchesAPIMartSample(t *testing.T) {
	body := []byte(`{
		"code": 200,
		"data": {
			"id": "task_01M6TJZF7YAXF11RCVN0HC0KEE",
			"status": "completed",
			"result": {"images": [{"url": ["https://cdn/a.png"]}]},
			"usage": {
				"input_tokens": 1782,
				"input_tokens_details": {"cached_tokens": 0, "image_tokens": 1508, "text_tokens": 274},
				"output_tokens": 1449,
				"output_tokens_details": {"image_tokens": 1449, "text_tokens": 0},
				"total_tokens": 3231
			}
		}
	}`)
	info, err := parseTaskResult(body)
	require.NoError(t, err)
	require.Equal(t, 274, info.TokenUsage.TextInputTokens)
	require.Equal(t, 1508, info.TokenUsage.ImageInputTokens)
	require.Equal(t, 1449, info.TokenUsage.ImageOutputTokens)

	// 系数是平台 $/1M tokens。40 Credits÷10=$4，64÷10=$6.4，240÷10=$24。
	expr := `tier("actual", p * 4 + cr * 1 + img * 6.4 + cc * 1.6 + img_o * 24)`
	tokenTask := &model.Task{}
	tokenTask.PrivateData.BillingContext = &model.TaskBillingContext{
		OriginModelName: "gpt-image-2.6-flare",
		GroupRatio:      1,
		QuotaPerUnit:    500000,
		BillingExpr:     expr,
	}
	require.Equal(t, 22762, (&TaskAdaptor{}).AdjustBillingOnComplete(tokenTask, info))
	tokenTask.PrivateData.BillingContext.GroupRatio = 0.5
	require.Equal(t, 11381, (&TaskAdaptor{}).AdjustBillingOnComplete(tokenTask, info))
	tokenTask.PrivateData.BillingContext.BillingExpr = ""
	require.Equal(t, 0, (&TaskAdaptor{}).AdjustBillingOnComplete(tokenTask, info))

	perImage := &model.Task{}
	perImage.Quota = 1000
	perImage.PrivateData.BillingContext = &model.TaskBillingContext{
		OriginModelName: "apimart-per-image",
		RequestedN:      1,
		GroupRatio:      1,
		QuotaPerUnit:    500000,
	}
	require.Equal(t, 0, (&TaskAdaptor{}).AdjustBillingOnComplete(perImage, info))
}

func TestBuildFluxUpstreamBody(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	images := make([]string, 11)
	for i := range images {
		images[i] = fmt.Sprintf("https://example.com/%d.png", i)
	}
	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:      "flux-3-image",
		Prompt:     "a coastal road",
		Ratio:      "21:9",
		Resolution: "2K",
		Images:     images,
		Metadata: map[string]interface{}{
			"safety_tolerance": float64(0),
			"grounding":        false,
			"mask_url":         "https://example.com/mask.png",
		},
	})
	info := &relaycommon.RelayInfo{
		ChannelMeta:   &relaycommon.ChannelMeta{UpstreamModelName: "flux-3-image"},
		TaskRelayInfo: &relaycommon.TaskRelayInfo{},
	}
	reader, err := (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw := readAll(t, reader)
	require.Equal(t, "21:9", gjson.GetBytes(raw, "aspect_ratio").String())
	require.Equal(t, "2k", gjson.GetBytes(raw, "resolution").String())
	require.False(t, gjson.GetBytes(raw, "size").Exists())
	require.False(t, gjson.GetBytes(raw, "quality").Exists())
	require.False(t, gjson.GetBytes(raw, "mask_url").Exists())
	require.Equal(t, int64(0), gjson.GetBytes(raw, "safety_tolerance").Int())
	require.True(t, gjson.GetBytes(raw, "safety_tolerance").Exists())
	require.True(t, gjson.GetBytes(raw, "grounding").Exists())
	require.False(t, gjson.GetBytes(raw, "grounding").Bool())
	require.Equal(t, 10, len(gjson.GetBytes(raw, "image_urls").Array()))

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:        "flux-3-image-klein",
		Prompt:       "a fox",
		OutputFormat: "jpg",
		Metadata: map[string]interface{}{
			"seed":  float64(42),
			"steps": float64(16),
		},
	})
	info.ChannelMeta.UpstreamModelName = "flux-3-image-klein"
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, "1k", gjson.GetBytes(raw, "resolution").String())
	require.Equal(t, "jpeg", gjson.GetBytes(raw, "output_format").String())
	require.Equal(t, int64(42), gjson.GetBytes(raw, "seed").Int())
	require.Equal(t, int64(16), gjson.GetBytes(raw, "steps").Int())
	require.False(t, gjson.GetBytes(raw, "aspect_ratio").Exists())
	require.False(t, gjson.GetBytes(raw, "grounding").Exists())
}

func TestValidateFluxExtras(t *testing.T) {
	err := validateExtras(relaycommon.TaskSubmitReq{
		Model:        "flux-3-image",
		OutputFormat: "png",
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model: "flux-3-image",
		Metadata: map[string]interface{}{
			"seed": float64(1),
		},
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model: "flux-3-image",
		Metadata: map[string]interface{}{
			"grounding": "false",
		},
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model: "flux-3-image",
		Metadata: map[string]interface{}{
			"safety_tolerance": float64(0),
			"grounding":        true,
		},
	})
	require.NoError(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model: "flux-3-image-klein",
		Metadata: map[string]interface{}{
			"grounding": false,
		},
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model:        "flux-3-image-klein",
		OutputFormat: "webp",
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model:        "flux-3-image-klein",
		OutputFormat: "jpg",
		Metadata: map[string]interface{}{
			"seed":  float64(0),
			"steps": float64(1),
		},
	})
	require.NoError(t, err)
}

func TestBuildMaiQwenMuseHyBodies(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	info := &relaycommon.RelayInfo{
		ChannelMeta:   &relaycommon.ChannelMeta{UpstreamModelName: "mai-image-2.6"},
		TaskRelayInfo: &relaycommon.TaskRelayInfo{},
	}
	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:      "mai-image-2.6",
		Prompt:     "a campus",
		Ratio:      "16:9",
		Resolution: "2K",
		Metadata: map[string]interface{}{
			"web_grounding": false,
		},
	})
	reader, err := (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw := readAll(t, reader)
	require.Equal(t, "16:9", gjson.GetBytes(raw, "size").String())
	require.Equal(t, "2K", gjson.GetBytes(raw, "resolution").String())
	require.True(t, gjson.GetBytes(raw, "web_grounding").Exists())
	require.False(t, gjson.GetBytes(raw, "web_grounding").Bool())
	require.False(t, gjson.GetBytes(raw, "quality").Exists())

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:  "mai-image-2.6-flash",
		Prompt: "pixels",
		Size:   "1536x1024",
	})
	info.ChannelMeta.UpstreamModelName = "mai-image-2.6-flash"
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, "1536x1024", gjson.GetBytes(raw, "size").String())
	require.False(t, gjson.GetBytes(raw, "resolution").Exists())

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model: "mai-image-2.6",
		Metadata: map[string]interface{}{
			"width": float64(1024),
		},
	})
	require.Error(t, err)

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:  "qwen-image-3.0",
		Prompt: "a poster",
		Ratio:  "16:9",
		N:      4,
		Metadata: map[string]interface{}{
			"negative_prompt":    "watermark",
			"prompt_extend":      true,
			"prompt_extend_mode": "direct",
		},
	})
	info.ChannelMeta.UpstreamModelName = "qwen-image-3.0"
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, "16:9", gjson.GetBytes(raw, "size").String())
	require.Equal(t, "1K", gjson.GetBytes(raw, "resolution").String())
	require.Equal(t, int64(4), gjson.GetBytes(raw, "n").Int())
	require.Equal(t, "watermark", gjson.GetBytes(raw, "negative_prompt").String())
	require.True(t, gjson.GetBytes(raw, "prompt_extend").Bool())
	require.Equal(t, "direct", gjson.GetBytes(raw, "prompt_extend_mode").String())

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model:  "qwen-image-3.0-pro",
		Images: []string{"https://example.com/a.png"},
		Metadata: map[string]interface{}{
			"prompt_extend":      true,
			"prompt_extend_mode": "agent",
		},
	})
	require.Error(t, err)

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:  "muse-image-ext",
		Prompt: "a cup",
		Ratio:  "1:1",
	})
	info.ChannelMeta.UpstreamModelName = "muse-image-ext"
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, "1:1", gjson.GetBytes(raw, "aspect_ratio").String())
	require.False(t, gjson.GetBytes(raw, "resolution").Exists())
	require.False(t, gjson.GetBytes(raw, "size").Exists())

	c.Set("task_request", relaycommon.TaskSubmitReq{
		Model:  "hy-image-v3.5-preview",
		Prompt: "a dog",
		Metadata: map[string]interface{}{
			"width":           float64(2000),
			"height":          float64(1104),
			"seed":            float64(0),
			"negative_prompt": "blur",
		},
	})
	info.ChannelMeta.UpstreamModelName = "hy-image-v3.5-preview"
	reader, err = (&TaskAdaptor{}).BuildRequestBody(c, info)
	require.NoError(t, err)
	raw = readAll(t, reader)
	require.Equal(t, int64(2000), gjson.GetBytes(raw, "width").Int())
	require.Equal(t, int64(1104), gjson.GetBytes(raw, "height").Int())
	require.True(t, gjson.GetBytes(raw, "seed").Exists())
	require.Equal(t, int64(0), gjson.GetBytes(raw, "seed").Int())
	require.Equal(t, "blur", gjson.GetBytes(raw, "negative_prompt").String())
	require.False(t, gjson.GetBytes(raw, "resolution").Exists())
	require.False(t, gjson.GetBytes(raw, "size").Exists())
}

func TestValidateExtras(t *testing.T) {
	err := validateExtras(relaycommon.TaskSubmitReq{
		Model: "gpt-image-2.5-flare",
		Metadata: map[string]interface{}{
			"markup_url": "https://example.com/m.png",
		},
	})
	require.Error(t, err)

	err = validateExtras(relaycommon.TaskSubmitReq{
		Model:        "gpt-image-2.6-flare",
		OutputFormat: "png",
		Images:       []string{"https://example.com/a.png"},
		Metadata: map[string]interface{}{
			"background": "transparent",
			"complexity": "low",
		},
	})
	require.NoError(t, err)
}

func readAll(t *testing.T, reader io.Reader) []byte {
	t.Helper()
	raw, err := io.ReadAll(reader)
	require.NoError(t, err)
	return raw
}
