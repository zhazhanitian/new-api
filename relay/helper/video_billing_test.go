package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

func TestPatchVideoBillingBodyRatioUsesResolution(t *testing.T) {
	body := []byte(`{"size":"9:16","metadata":{"resolution":"1080P","audio_generation":"Enabled"}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "1080P", gjson.GetBytes(patched, "size").String())
	require.Equal(t, "Enabled", gjson.GetBytes(patched, "metadata.audio_generation").String())
}

func TestPatchVideoBillingBodyRatioWithoutResolutionDropsSize(t *testing.T) {
	body := []byte(`{"size":"9:16","metadata":{"audio_generation":"Enabled"}}`)
	patched := PatchVideoBillingBody(body)
	require.False(t, gjson.GetBytes(patched, "size").Exists())
	require.Equal(t, "Enabled", gjson.GetBytes(patched, "metadata.audio_generation").String())
}

func TestPatchVideoBillingBodyKeepsResolutionSize(t *testing.T) {
	body := []byte(`{"size":"720P","metadata":{"resolution":"1080P"}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "720P", gjson.GetBytes(patched, "size").String())
}

func TestPatchVideoBillingBodyReferenceImages(t *testing.T) {
	body := []byte(`{"metadata":{"content":[{"type":"image_url","url":"https://img/a.png","role":"reference_image"},{"type":"image_url","url":"https://img/b.png","role":"reference_image"}]}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "Reference", gjson.GetBytes(patched, "metadata.input_usage").String())
	require.Equal(t, []string{"https://img/a.png", "https://img/b.png"}, gjsonStringList(patched, "images"))
}

func TestPatchVideoBillingBodyFirstFrameDoesNotFillImages(t *testing.T) {
	body := []byte(`{"metadata":{"content":[{"type":"image_url","url":"https://img/a.png","role":"first_frame"}]}}`)
	patched := PatchVideoBillingBody(body)
	require.False(t, gjson.GetBytes(patched, "images").Exists())
	require.False(t, gjson.GetBytes(patched, "metadata.input_usage").Exists())
}

func TestPatchVideoBillingBodyEmptyImageRoleBecomesReference(t *testing.T) {
	body := []byte(`{"metadata":{"content":[{"type":"image_url","image_url":{"url":"https://img/a.png"}}]}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "reference_image", gjson.GetBytes(patched, "metadata.content.0.role").String())
	require.Equal(t, int64(1), gjson.GetBytes(patched, "metadata.content.#(role==reference_image)#|#").Int())
	require.Equal(t, "Reference", gjson.GetBytes(patched, "metadata.input_usage").String())
	require.Equal(t, []string{"https://img/a.png"}, gjsonStringList(patched, "images"))
}

func TestPatchVideoBillingBodyReferenceVideo(t *testing.T) {
	body := []byte(`{"metadata":{"content":[{"type":"video_url","video_url":{"url":"https://cdn/a.mp4"}}],"video_url":"https://cdn/legacy.mp4"}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, int64(2), gjson.GetBytes(patched, "metadata.content.#(role==reference_video)#|#").Int())
	require.Equal(t, int64(0), gjson.GetBytes(patched, "metadata.content.#(role==reference_image)#|#").Int())
	require.False(t, gjson.GetBytes(patched, "images").Exists())
	require.False(t, gjson.GetBytes(patched, "metadata.input_usage").Exists())
}

func TestPatchVideoBillingBodyLegacyImagesLandInContent(t *testing.T) {
	body := []byte(`{"images":["https://img/old.png"]}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "reference_image", gjson.GetBytes(patched, "metadata.content.0.role").String())
	require.Equal(t, "https://img/old.png", gjson.GetBytes(patched, "metadata.content.0.url").String())
	require.Equal(t, []string{"https://img/old.png"}, gjsonStringList(patched, "images"))
}

func TestPatchVideoBillingBodyFirstFrameStaysFirstFrame(t *testing.T) {
	body := []byte(`{"metadata":{"content":[{"type":"image_url","url":"https://img/a.png","role":"first_frame"},{"type":"image_url","url":"https://img/b.png"}]}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, "first_frame", gjson.GetBytes(patched, "metadata.content.0.role").String())
	require.Equal(t, "reference_image", gjson.GetBytes(patched, "metadata.content.1.role").String())
	require.Equal(t, int64(1), gjson.GetBytes(patched, "metadata.content.#(role==reference_image)#|#").Int())
	require.False(t, gjson.GetBytes(patched, "metadata.input_usage").Exists())
}

func TestPatchVideoBillingBodyKeepsExistingImages(t *testing.T) {
	body := []byte(`{"images":["https://img/kept.png"],"metadata":{"content":[{"type":"image_url","url":"https://img/other.png","role":"reference_image"}]}}`)
	patched := PatchVideoBillingBody(body)
	require.Equal(t, []string{"https://img/kept.png"}, gjsonStringList(patched, "images"))
	require.Equal(t, "Reference", gjson.GetBytes(patched, "metadata.input_usage").String())
}

func gjsonStringList(body []byte, path string) []string {
	values := gjson.GetBytes(body, path).Array()
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, value.String())
	}
	return out
}
