package taskcommon

import (
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/require"
)

func TestNormalizeVideoContentImagesDefaultToReference(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png", "https://img/b.png"},
	})
	require.Len(t, items, 2)
	require.Equal(t, VideoRoleReferenceImage, items[0].Role)
	require.Equal(t, VideoRoleReferenceImage, items[1].Role)
	require.Equal(t, "https://img/a.png", items[0].URL)
}

func TestNormalizeVideoContentSingleImageIsReference(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png"},
	})
	require.Len(t, items, 1)
	require.Equal(t, VideoContentImage, items[0].Type)
	require.Equal(t, VideoRoleReferenceImage, items[0].Role)
}

func TestNormalizeVideoContentHistoricalFirstFrame(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png", "https://img/b.png"},
		Metadata: map[string]interface{}{
			"input_usage": "FirstFrame",
		},
	})
	require.Equal(t, VideoRoleFirstFrame, items[0].Role)
	require.Equal(t, VideoRoleReferenceImage, items[1].Role)
}

func TestNormalizeVideoContentKeepsCallerRole(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
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
	require.Len(t, items, 1)
	require.Equal(t, VideoRoleFirstFrame, items[0].Role)
	require.Equal(t, "https://img/first.png", items[0].URL)
}

func TestNormalizeVideoContentLastFramePromotesFirstImage(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png", "https://img/b.png"},
		Metadata: map[string]interface{}{
			"last_frame": "https://img/last.png",
		},
	})
	require.Len(t, items, 3)
	require.Equal(t, VideoRoleFirstFrame, items[0].Role)
	require.Equal(t, "https://img/a.png", items[0].URL)
	require.Equal(t, VideoRoleReferenceImage, items[1].Role)
	require.Equal(t, VideoRoleLastFrame, items[2].Role)
	require.Equal(t, "https://img/last.png", items[2].URL)
}

func TestNormalizeVideoContentLastFrameKeepsExistingFirstFrame(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/extra.png"},
		Metadata: map[string]interface{}{
			"last_frame": "https://img/last.png",
			"content": []interface{}{
				map[string]interface{}{
					"type": "image_url",
					"url":  "https://img/first.png",
					"role": "first_frame",
				},
			},
		},
	})
	require.Equal(t, VideoRoleFirstFrame, items[0].Role)
	require.Equal(t, VideoRoleReferenceImage, items[1].Role)
	require.Equal(t, VideoRoleLastFrame, items[2].Role)
}

func TestNormalizeVideoContentAppendsLegacyFieldsWithoutDedupe(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png"},
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type": "image_url",
					"url":  "https://img/a.png",
					"role": "reference_image",
				},
			},
		},
	})
	require.Len(t, items, 2)
	require.Equal(t, "https://img/a.png", items[0].URL)
	require.Equal(t, "https://img/a.png", items[1].URL)
}

func TestNormalizeVideoContentVideoURLIsReference(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Metadata: map[string]interface{}{
			"video_url": "https://cdn/ref.mp4",
		},
	})
	require.Len(t, items, 1)
	require.Equal(t, VideoContentVideo, items[0].Type)
	require.Equal(t, VideoRoleReferenceVideo, items[0].Role)
}

func TestNormalizeVideoContentKeepsItemReferenceType(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Metadata: map[string]interface{}{
			"reference_type": "subject",
			"content": []interface{}{
				map[string]interface{}{
					"type":           "image_url",
					"url":            "https://img/cat.png",
					"role":           "reference_image",
					"reference_type": "background",
					"name":           "小猫",
				},
			},
		},
	})
	require.Equal(t, "background", items[0].ReferenceType)
	require.Equal(t, "小猫", items[0].Name)
}

func TestNormalizeVideoContentEmptyRoleStaysEmpty(t *testing.T) {
	items := NormalizeVideoContent(relaycommon.TaskSubmitReq{
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type":      "image_url",
					"image_url": map[string]interface{}{"url": "https://img/a.png"},
				},
			},
		},
	})
	require.Empty(t, items[0].Role)
}

func TestIsAspectRatioSize(t *testing.T) {
	require.True(t, IsAspectRatioSize("9:16"))
	require.True(t, IsAspectRatioSize(" 21:9 "))
	require.False(t, IsAspectRatioSize("720P"))
	require.False(t, IsAspectRatioSize("1920x1080"))
	require.False(t, IsAspectRatioSize(""))
}
