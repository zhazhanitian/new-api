package tencentvod

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestV1AndV2ImageBodiesMatch(t *testing.T) {
	gin.SetMode(gin.TestMode)
	modelName := "gemini-3.1-flash-image"
	info := &relaycommon.RelayInfo{
		OriginModelName: modelName,
		ChannelMeta:     &relaycommon.ChannelMeta{UpstreamModelName: modelName},
		TaskRelayInfo:   &relaycommon.TaskRelayInfo{},
	}
	imageReq := dto.ImageRequest{
		Model:  modelName,
		Prompt: "a cat",
		Extra: map[string]json.RawMessage{
			"resolution": json.RawMessage(`"720P"`),
			"ratio":      json.RawMessage(`"1:8"`),
		},
	}
	syncBody, err := (&Adaptor{subAppId: 9}).ConvertImageRequest(nil, info, imageReq)
	require.NoError(t, err)
	syncJSON, err := common.Marshal(syncBody)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Set("task_request", relaycommon.TaskSubmitReq{
		Prompt:     "a cat",
		Resolution: "720P",
		Ratio:      "1:8",
	})
	reader, err := (&TaskAdaptor{subAppId: 9}).BuildRequestBody(c, info)
	require.NoError(t, err)
	asyncJSON, err := io.ReadAll(reader)
	require.NoError(t, err)
	require.True(t, bytes.Equal(syncJSON, asyncJSON), "v1=%s v2=%s", syncJSON, asyncJSON)
}

func TestBuildImageCreateRequestFieldMapping(t *testing.T) {
	ogSize := buildImageCreateRequest(imageBuildInput{
		SubAppId:    1,
		LookupModel: "gpt-image-2",
		Prompt:      "cat",
		Size:        "1024x1024",
		Quality:     "high",
	})
	require.Equal(t, "image2_high", ogSize.ModelVersion)
	require.Contains(t, ogSize.ExtInfo, "1024x1024")
	require.Nil(t, ogSize.OutputConfig)

	ogTier := buildImageCreateRequest(imageBuildInput{
		SubAppId:    1,
		LookupModel: "gpt-image-2",
		Prompt:      "cat",
		Resolution:  "4K",
		Ratio:       "16:9",
		Quality:     "xhigh",
	})
	require.Equal(t, "image2_low", ogTier.ModelVersion)
	require.Empty(t, ogTier.ExtInfo)
	require.Equal(t, "4K", ogTier.OutputConfig.Resolution)
	require.Equal(t, "16:9", ogTier.OutputConfig.AspectRatio)

	gg := buildImageCreateRequest(imageBuildInput{
		SubAppId:    1,
		LookupModel: "GG-3-pro-image",
		Prompt:      "cat",
		Resolution:  "1K",
		Ratio:       "1:1",
	})
	require.Equal(t, "GG", gg.ModelName)
	require.Equal(t, "3.0", gg.ModelVersion)
	require.Equal(t, "1K", gg.OutputConfig.Resolution)
	require.Equal(t, "1:1", gg.OutputConfig.AspectRatio)

	scene := buildImageCreateRequest(imageBuildInput{
		SubAppId:    1,
		LookupModel: "kling-image-scene",
		Prompt:      "expand",
		Extra: map[string]interface{}{
			"scene_type": "image_expand",
			"expansion":  map[string]interface{}{"up": 0.2},
		},
	})
	require.Equal(t, "scene", scene.ModelVersion)
	require.Nil(t, scene.OutputConfig)
	require.Equal(t, "image_expand", scene.SceneType)
	require.NotEmpty(t, scene.ExtInfo)

	kling := buildImageCreateRequest(imageBuildInput{
		SubAppId:    1,
		LookupModel: "kling-image-2.1",
		Prompt:      "cat",
		Resolution:  "2k",
		Ratio:       "9:16",
	})
	raw, err := common.Marshal(kling)
	require.NoError(t, err)
	require.Equal(t, "2k", gjson.GetBytes(raw, "OutputConfig.Resolution").String())
	require.Equal(t, "9:16", gjson.GetBytes(raw, "OutputConfig.AspectRatio").String())
}
