package apimart_image

import (
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
