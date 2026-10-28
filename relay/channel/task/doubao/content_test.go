package doubao

import (
	"testing"

	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/require"
)

func TestConvertToRequestPayloadFillsEmptyRole(t *testing.T) {
	payload, err := (&TaskAdaptor{}).convertToRequestPayload(&relaycommon.TaskSubmitReq{
		Model:  "doubao-seedance-2-0-260128",
		Prompt: "a cat",
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type":      "image_url",
					"image_url": map[string]interface{}{"url": "https://img/a.png"},
				},
				map[string]interface{}{
					"type":      "video_url",
					"video_url": map[string]interface{}{"url": "https://cdn/a.mp4"},
				},
			},
			"video_url": "https://cdn/legacy.mp4",
		},
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(payload.Content), 4)
	require.Equal(t, "reference_image", payload.Content[0].Role)
	require.Equal(t, "https://img/a.png", payload.Content[0].ImageURL.URL)
	require.Equal(t, "reference_video", payload.Content[1].Role)
	require.Equal(t, "https://cdn/a.mp4", payload.Content[1].VideoURL.URL)
	require.Equal(t, "reference_video", payload.Content[2].Role)
	require.Equal(t, "https://cdn/legacy.mp4", payload.Content[2].VideoURL.URL)
	require.Equal(t, "text", payload.Content[3].Type)
	require.Equal(t, "a cat", payload.Content[3].Text)
}

func TestConvertToRequestPayloadKeepsCallerRoleAndAppendsImages(t *testing.T) {
	payload, err := (&TaskAdaptor{}).convertToRequestPayload(&relaycommon.TaskSubmitReq{
		Model:  "doubao-seedance-1-0-pro-250528",
		Images: []string{"https://img/extra.png"},
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type":      "image_url",
					"image_url": map[string]interface{}{"url": "https://img/first.png"},
					"role":      "first_frame",
				},
			},
		},
	})
	require.NoError(t, err)
	require.Len(t, payload.Content, 2)
	require.Equal(t, "first_frame", payload.Content[0].Role)
	require.Equal(t, "reference_image", payload.Content[1].Role)
	require.Equal(t, "https://img/extra.png", payload.Content[1].ImageURL.URL)
}

func TestConvertDurationNegativeOneOverridesFrames(t *testing.T) {
	payload, err := (&TaskAdaptor{}).convertToRequestPayload(&relaycommon.TaskSubmitReq{
		Model:    "doubao-seedance-2-5-260628",
		Prompt:   "edit",
		Duration: -1,
		Metadata: map[string]interface{}{"frames": 57},
	})
	require.NoError(t, err)
	require.NotNil(t, payload.Duration)
	require.Equal(t, dto.IntValue(-1), *payload.Duration)
	require.Nil(t, payload.Frames)
}

func TestConvertSecondsNegativeOne(t *testing.T) {
	payload, err := (&TaskAdaptor{}).convertToRequestPayload(&relaycommon.TaskSubmitReq{
		Model:   "doubao-seedance-2-5-260628",
		Prompt:  "edit",
		Seconds: "-1",
	})
	require.NoError(t, err)
	require.NotNil(t, payload.Duration)
	require.Equal(t, dto.IntValue(-1), *payload.Duration)
}

func TestConvertSize4kUsesOfficialCasing(t *testing.T) {
	payload, err := (&TaskAdaptor{}).convertToRequestPayload(&relaycommon.TaskSubmitReq{
		Model:  "doubao-seedance-2-0-260128",
		Prompt: "city",
		Size:   "4K",
	})
	require.NoError(t, err)
	require.Equal(t, "4k", payload.Resolution)
}

func TestHasAudioDefaultsToEnabled(t *testing.T) {
	require.True(t, hasAudioInMetadata(nil))
	require.True(t, hasAudioInMetadata(map[string]interface{}{}))
	require.True(t, hasAudioInMetadata(map[string]interface{}{"generate_audio": true}))
	require.False(t, hasAudioInMetadata(map[string]interface{}{"generate_audio": false}))
}

func TestResolutionRatioAccepts4kCasing(t *testing.T) {
	lower, ok := GetResolutionRatio("doubao-seedance-2-0-260128", "4k")
	require.True(t, ok)
	upper, ok := GetResolutionRatio("doubao-seedance-2-0-260128", "4K")
	require.True(t, ok)
	require.Equal(t, upper, lower)
}
