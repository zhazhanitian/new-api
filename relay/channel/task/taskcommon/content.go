package taskcommon

import (
	"strings"
	"unicode"

	"github.com/QuantumNous/new-api/common"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
)

const (
	VideoContentText  = "text"
	VideoContentImage = "image_url"
	VideoContentVideo = "video_url"
	VideoContentAudio = "audio_url"
	VideoContentDraft = "draft_task"

	VideoRoleFirstFrame     = "first_frame"
	VideoRoleLastFrame      = "last_frame"
	VideoRoleReferenceImage = "reference_image"
	VideoRoleReferenceVideo = "reference_video"
	VideoRoleReferenceAudio = "reference_audio"
)

// VideoContentItem is one entry in the shared video attachment list.
// URL is flattened from image_url / video_url / audio_url objects.
// ReferenceType and Name are kept for Tencent; Doubao ignores them.
type VideoContentItem struct {
	Type          string
	URL           string
	Role          string
	ReferenceType string
	Name          string
	Text          string
	DraftTaskID   string
}

type rawVideoContentItem struct {
	Type          string        `json:"type"`
	URL           string        `json:"url"`
	Role          string        `json:"role"`
	ReferenceType string        `json:"reference_type"`
	Name          string        `json:"name"`
	Text          string        `json:"text"`
	ImageURL      *rawMediaURL  `json:"image_url"`
	VideoURL      *rawMediaURL  `json:"video_url"`
	AudioURL      *rawMediaURL  `json:"audio_url"`
	DraftTask     *rawDraftTask `json:"draft_task"`
}

type rawMediaURL struct {
	URL string `json:"url"`
}

type rawDraftTask struct {
	ID string `json:"id"`
}

// NormalizeVideoContent folds caller content and legacy attachment fields into one list.
// Caller roles are kept. Empty roles are not filled here.
// images, metadata.last_frame and metadata.video_url are appended without dedupe.
func NormalizeVideoContent(req relaycommon.TaskSubmitReq) []VideoContentItem {
	items := parseCallerVideoContent(req.Metadata)

	imageStart := len(items)
	firstFrameUsage := strings.EqualFold(metadataString(req.Metadata, "input_usage"), "FirstFrame")
	for i, img := range req.Images {
		img = strings.TrimSpace(img)
		if img == "" {
			continue
		}
		role := VideoRoleReferenceImage
		if i == 0 && firstFrameUsage {
			role = VideoRoleFirstFrame
		}
		items = append(items, VideoContentItem{
			Type: VideoContentImage,
			URL:  img,
			Role: role,
		})
	}

	if videoURL := metadataString(req.Metadata, "video_url"); videoURL != "" {
		items = append(items, VideoContentItem{
			Type: VideoContentVideo,
			URL:  videoURL,
			Role: VideoRoleReferenceVideo,
		})
	}

	if lastFrame := metadataString(req.Metadata, "last_frame"); lastFrame != "" {
		if !hasVideoRole(items, VideoRoleFirstFrame) {
			for i := imageStart; i < len(items); i++ {
				if items[i].Type == VideoContentImage && items[i].URL != "" {
					items[i].Role = VideoRoleFirstFrame
					break
				}
			}
		}
		items = append(items, VideoContentItem{
			Type: VideoContentImage,
			URL:  lastFrame,
			Role: VideoRoleLastFrame,
		})
	}

	return items
}

// IsAspectRatioSize reports whether size is a ratio token such as 9:16.
// Pixel sizes like 1920x1080 and resolution tokens like 720P are not ratios.
func IsAspectRatioSize(size string) bool {
	size = strings.TrimSpace(size)
	if size == "" || strings.ContainsAny(size, "xX") {
		return false
	}
	left, right, ok := strings.Cut(size, ":")
	if !ok || left == "" || right == "" || strings.Contains(right, ":") {
		return false
	}
	return isASCIIDigits(left) && isASCIIDigits(right)
}

func parseCallerVideoContent(metadata map[string]interface{}) []VideoContentItem {
	if metadata == nil {
		return nil
	}
	raw, ok := metadata["content"]
	if !ok || raw == nil {
		return nil
	}
	encoded, err := common.Marshal(raw)
	if err != nil {
		return nil
	}
	var rows []rawVideoContentItem
	if err := common.Unmarshal(encoded, &rows); err != nil {
		return nil
	}
	items := make([]VideoContentItem, 0, len(rows))
	for _, row := range rows {
		item := videoContentFromRaw(row)
		if item.Type == "" && item.URL == "" && item.Text == "" && item.DraftTaskID == "" {
			continue
		}
		items = append(items, item)
	}
	return items
}

func videoContentFromRaw(row rawVideoContentItem) VideoContentItem {
	itemType := strings.TrimSpace(row.Type)
	url := strings.TrimSpace(row.URL)
	switch itemType {
	case VideoContentImage:
		url = firstNonEmpty(nestedMediaURL(row.ImageURL), url)
	case VideoContentVideo:
		url = firstNonEmpty(nestedMediaURL(row.VideoURL), url)
	case VideoContentAudio:
		url = firstNonEmpty(nestedMediaURL(row.AudioURL), url)
	case "":
		switch {
		case row.ImageURL != nil:
			itemType = VideoContentImage
			url = firstNonEmpty(nestedMediaURL(row.ImageURL), url)
		case row.VideoURL != nil:
			itemType = VideoContentVideo
			url = firstNonEmpty(nestedMediaURL(row.VideoURL), url)
		case row.AudioURL != nil:
			itemType = VideoContentAudio
			url = firstNonEmpty(nestedMediaURL(row.AudioURL), url)
		case strings.TrimSpace(row.Text) != "":
			itemType = VideoContentText
		case row.DraftTask != nil && strings.TrimSpace(row.DraftTask.ID) != "":
			itemType = VideoContentDraft
		}
	}
	draftID := ""
	if row.DraftTask != nil {
		draftID = strings.TrimSpace(row.DraftTask.ID)
	}
	return VideoContentItem{
		Type:          itemType,
		URL:           url,
		Role:          strings.TrimSpace(row.Role),
		ReferenceType: strings.TrimSpace(row.ReferenceType),
		Name:          strings.TrimSpace(row.Name),
		Text:          row.Text,
		DraftTaskID:   draftID,
	}
}

func nestedMediaURL(media *rawMediaURL) string {
	if media == nil {
		return ""
	}
	return strings.TrimSpace(media.URL)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return strings.TrimSpace(value)
		}
	}
	return ""
}

func metadataString(metadata map[string]interface{}, key string) string {
	if metadata == nil {
		return ""
	}
	value, _ := metadata[key].(string)
	return strings.TrimSpace(value)
}

func hasVideoRole(items []VideoContentItem, role string) bool {
	for _, item := range items {
		if item.Role == role {
			return true
		}
	}
	return false
}

func isASCIIDigits(value string) bool {
	if value == "" {
		return false
	}
	for _, r := range value {
		if r > unicode.MaxASCII || !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
