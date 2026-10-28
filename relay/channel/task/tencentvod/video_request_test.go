package tencentvod

import (
	"io"
	"net/http/httptest"
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestBuildVideoRequestBodyReferenceImagesAndSeed(t *testing.T) {
	seed := 7
	body, info := buildVideoBody(t, "kling-video-2.5", relaycommon.TaskSubmitReq{
		Prompt:   "a cat",
		Duration: 5,
		Size:     "9:16",
		Seed:     &seed,
		Images:   []string{"https://img/a.png", "https://img/b.png"},
		Metadata: map[string]interface{}{
			"resolution":       "1080P",
			"seed":             3,
			"audio_generation": "Enabled",
			"input_region":     "Mainland",
			"ext_info":         `{"k":"v"}`,
		},
	})

	require.Equal(t, float64(5), info.BilledDurationSec)
	require.Equal(t, int64(7), gjson.GetBytes(body, "Seed").Int())
	require.Equal(t, "1080P", gjson.GetBytes(body, "OutputConfig.Resolution").String())
	require.False(t, gjson.GetBytes(body, "OutputConfig.AspectRatio").Exists())
	require.Equal(t, "Enabled", gjson.GetBytes(body, "OutputConfig.AudioGeneration").String())
	require.False(t, gjson.GetBytes(body, "InputRegion").Exists())
	require.False(t, gjson.GetBytes(body, "ExtInfo").Exists())
	require.False(t, gjson.GetBytes(body, "LastFrameUrl").Exists())
	require.Equal(t, "Reference", gjson.GetBytes(body, "FileInfos.0.Usage").String())
	require.Equal(t, "Reference", gjson.GetBytes(body, "FileInfos.1.Usage").String())
	require.Equal(t, "Image", gjson.GetBytes(body, "FileInfos.0.Category").String())
}

func TestBuildVideoRequestBodyKlingWithoutFilesKeepsAspectRatio(t *testing.T) {
	body, info := buildVideoBody(t, "kling-video-2.5", relaycommon.TaskSubmitReq{
		Prompt: "a cat",
		Size:   "9:16",
	})
	require.Equal(t, float64(5), info.BilledDurationSec)
	require.Equal(t, "9:16", gjson.GetBytes(body, "OutputConfig.AspectRatio").String())
	require.False(t, gjson.GetBytes(body, "FileInfos").Exists())
}

func TestBuildVideoRequestBodyLastFrameAndReferenceVideo(t *testing.T) {
	body, _ := buildVideoBody(t, "vidu-video-q2-pro", relaycommon.TaskSubmitReq{
		Prompt: "a cat",
		Images: []string{"https://img/a.png"},
		Metadata: map[string]interface{}{
			"last_frame": "https://img/last.png",
			"video_url":  "https://cdn/ref.mp4",
		},
	})
	require.Equal(t, "FirstFrame", gjson.GetBytes(body, "FileInfos.0.Usage").String())
	require.Equal(t, "https://img/a.png", gjson.GetBytes(body, "FileInfos.0.Url").String())
	require.Equal(t, "Video", gjson.GetBytes(body, "FileInfos.1.Category").String())
	require.Equal(t, "Reference", gjson.GetBytes(body, "FileInfos.1.Usage").String())
	require.Equal(t, "LastFrame", gjson.GetBytes(body, "FileInfos.2.Usage").String())
	require.False(t, gjson.GetBytes(body, "LastFrameUrl").Exists())
}

func TestBuildVideoRequestBodyContentReferenceTypeAndAudio(t *testing.T) {
	body, _ := buildVideoBody(t, "pixverse-video-v6", relaycommon.TaskSubmitReq{
		Prompt: "a cat",
		Metadata: map[string]interface{}{
			"audio_generation":  "Enabled",
			"reference_type":    "subject",
			"pixverse_subjects": `[{"url":"https://img/ignored.png","type":"subject","name":"旧"}]`,
			"content": []interface{}{
				map[string]interface{}{
					"type":           "image_url",
					"url":            "https://img/cat.png",
					"role":           "reference_image",
					"reference_type": "background",
					"name":           "小猫",
				},
				map[string]interface{}{
					"type": "audio_url",
					"url":  "https://cdn/a.mp3",
				},
			},
		},
	})
	require.Equal(t, "Enabled", gjson.GetBytes(body, "OutputConfig.AudioGeneration").String())
	require.Equal(t, "background", gjson.GetBytes(body, "FileInfos.0.ReferenceType").String())
	require.Equal(t, "小猫", gjson.GetBytes(body, "FileInfos.0.Text").String())
	require.Equal(t, "Audio", gjson.GetBytes(body, "FileInfos.1.Category").String())
	require.Equal(t, "Reference", gjson.GetBytes(body, "FileInfos.1.Usage").String())
	require.Equal(t, int64(2), gjson.GetBytes(body, "FileInfos.#").Int())
}

func TestBuildVideoRequestBodyAppendsDuplicateContent(t *testing.T) {
	body, _ := buildVideoBody(t, "kling-video-2.5", relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png"},
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type": "image_url",
					"url":  "https://img/a.png",
					"role": "first_frame",
				},
			},
		},
	})
	require.Equal(t, int64(2), gjson.GetBytes(body, "FileInfos.#").Int())
	require.Equal(t, "FirstFrame", gjson.GetBytes(body, "FileInfos.0.Usage").String())
	require.Equal(t, "Reference", gjson.GetBytes(body, "FileInfos.1.Usage").String())
	require.False(t, gjson.GetBytes(body, "OutputConfig.AspectRatio").Exists())
}

func TestBuildVideoRequestBodyHailuoForwardsDurationResolutionAndRatio(t *testing.T) {
	body, info := buildVideoBody(t, "hailuo-video-h3", relaycommon.TaskSubmitReq{
		Prompt:   "a cat",
		Duration: 12,
		Size:     "2K",
		Ratio:    "16:9",
	})
	require.Equal(t, float64(12), info.BilledDurationSec)
	require.Equal(t, float64(12), gjson.GetBytes(body, "OutputConfig.Duration").Float())
	require.Equal(t, "2K", gjson.GetBytes(body, "OutputConfig.Resolution").String())
	require.Equal(t, "16:9", gjson.GetBytes(body, "OutputConfig.AspectRatio").String())
}

func TestBuildVideoRequestBodyHailuoOmitsUnsetDurationAndRatio(t *testing.T) {
	body, info := buildVideoBody(t, "hailuo-video-02", relaycommon.TaskSubmitReq{
		Prompt: "a cat",
		Size:   "1080P",
	})
	require.Equal(t, float64(6), info.BilledDurationSec)
	require.False(t, gjson.GetBytes(body, "OutputConfig.Duration").Exists())
	require.Equal(t, "1080P", gjson.GetBytes(body, "OutputConfig.Resolution").String())
	require.False(t, gjson.GetBytes(body, "OutputConfig.AspectRatio").Exists())
}

func buildVideoBody(t *testing.T, model string, req relaycommon.TaskSubmitReq) ([]byte, *relaycommon.RelayInfo) {
	t.Helper()
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("POST", "/v1/video/generations", nil)
	c.Set("task_request", req)
	info := &relaycommon.RelayInfo{
		OriginModelName: model,
		ChannelMeta:     &relaycommon.ChannelMeta{},
	}
	reader, err := (&TaskAdaptor{subAppId: 1}).buildVideoRequestBody(c, info)
	require.NoError(t, err)
	body, err := io.ReadAll(reader)
	require.NoError(t, err)
	return body, info
}
