package joyagent

import (
	"testing"

	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/require"
)

func TestHappyHorseMediaImagesAreReference(t *testing.T) {
	media := happyHorseMedia(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/a.png", "https://img/b.png"},
	})
	require.Equal(t, []mediaItem{
		{Type: "reference_image", URL: "https://img/a.png"},
		{Type: "reference_image", URL: "https://img/b.png"},
	}, media)
}

func TestHappyHorseMediaFirstFrameRole(t *testing.T) {
	media := happyHorseMedia(relaycommon.TaskSubmitReq{
		Metadata: map[string]interface{}{
			"content": []interface{}{
				map[string]interface{}{
					"type": "image_url",
					"url":  "https://img/first.png",
					"role": "first_frame",
				},
			},
		},
	})
	require.Equal(t, []mediaItem{{Type: "first_frame", URL: "https://img/first.png"}}, media)
}

func TestHappyHorseMediaVideoURL(t *testing.T) {
	media := happyHorseMedia(relaycommon.TaskSubmitReq{
		Images: []string{"https://img/style.png"},
		Metadata: map[string]interface{}{
			"video_url": "https://cdn/source.mp4",
		},
	})
	require.Equal(t, []mediaItem{
		{Type: "reference_image", URL: "https://img/style.png"},
		{Type: "video", URL: "https://cdn/source.mp4"},
	}, media)
}

func TestHappyHorseMediaInputReferenceIsReferenceImage(t *testing.T) {
	media := happyHorseMedia(relaycommon.TaskSubmitReq{
		InputReference: "https://img/legacy.png",
	})
	require.Equal(t, []mediaItem{{Type: "reference_image", URL: "https://img/legacy.png"}}, media)
}
