package constant

import (
	"net/http"
	"strings"
)

const (
	RelayModeUnknown = iota
	RelayModeChatCompletions
	RelayModeCompletions
	RelayModeEmbeddings
	RelayModeModerations
	RelayModeImagesGenerations
	RelayModeImagesEdits
	RelayModeEdits

	RelayModeMidjourneyImagine
	RelayModeMidjourneyDescribe
	RelayModeMidjourneyBlend
	RelayModeMidjourneyChange
	RelayModeMidjourneySimpleChange
	RelayModeMidjourneyNotify
	RelayModeMidjourneyTaskFetch
	RelayModeMidjourneyTaskImageSeed
	RelayModeMidjourneyTaskFetchByCondition
	RelayModeMidjourneyAction
	RelayModeMidjourneyModal
	RelayModeMidjourneyShorten
	RelayModeSwapFace
	RelayModeMidjourneyUpload
	RelayModeMidjourneyVideo
	RelayModeMidjourneyEdits

	RelayModeAudioSpeech        // tts
	RelayModeAudioTranscription // whisper
	RelayModeAudioTranslation   // whisper

	RelayModeSunoFetch
	RelayModeSunoFetchByID
	RelayModeSunoSubmit

	RelayModeVideoFetchByID
	RelayModeVideoSubmit

	RelayModeImageTaskFetchByID  // v2 图片任务查询
	RelayModeImageTaskSubmit     // v2 图片任务提交，使用 ImageRequest 格式

	RelayModeRerank

	RelayModeResponses

	RelayModeRealtime

	RelayModeGemini

	RelayModeResponsesCompact

	RelayMode3DTaskSubmit    // POST /v1/3d/generations
	RelayMode3DTaskFetchByID // GET  /v1/3d/generations/:task_id

	RelayModeMusicSubmit    // POST /v1/music/generations
	RelayModeMusicFetchByID // GET  /v1/music/generations/:task_id
)

func Path2RelayMode(path string) int {
	relayMode := RelayModeUnknown
	if strings.HasPrefix(path, "/v1/chat/completions") || strings.HasPrefix(path, "/pg/chat/completions") {
		relayMode = RelayModeChatCompletions
	} else if strings.HasPrefix(path, "/v1/completions") {
		relayMode = RelayModeCompletions
	} else if strings.HasPrefix(path, "/v1/embeddings") {
		relayMode = RelayModeEmbeddings
	} else if strings.HasSuffix(path, "embeddings") {
		relayMode = RelayModeEmbeddings
	} else if strings.HasPrefix(path, "/v1/moderations") {
		relayMode = RelayModeModerations
	} else if strings.HasPrefix(path, "/v1/images/generations") {
		relayMode = RelayModeImagesGenerations
	} else if strings.HasPrefix(path, "/v1/images/edits") {
		relayMode = RelayModeImagesEdits
	} else if strings.HasPrefix(path, "/v1/edits") {
		relayMode = RelayModeEdits
	} else if strings.HasPrefix(path, "/v1/responses/compact") {
		relayMode = RelayModeResponsesCompact
	} else if strings.HasPrefix(path, "/v1/responses") {
		relayMode = RelayModeResponses
	} else if strings.HasPrefix(path, "/v1/audio/speech") {
		relayMode = RelayModeAudioSpeech
	} else if strings.HasPrefix(path, "/v1/audio/transcriptions") {
		relayMode = RelayModeAudioTranscription
	} else if strings.HasPrefix(path, "/v1/audio/translations") {
		relayMode = RelayModeAudioTranslation
	} else if strings.HasPrefix(path, "/v1/rerank") {
		relayMode = RelayModeRerank
	} else if strings.HasPrefix(path, "/v1/realtime") {
		relayMode = RelayModeRealtime
	} else if strings.HasPrefix(path, "/v1beta/models") || strings.HasPrefix(path, "/v1/models") {
		relayMode = RelayModeGemini
	} else if strings.HasPrefix(path, "/mj") {
		relayMode = Path2RelayModeMidjourney(path)
	}
	return relayMode
}

func Path2RelayModeMidjourney(path string) int {
	relayMode := RelayModeUnknown
	if strings.HasSuffix(path, "/mj/submit/action") {
		// midjourney plus
		relayMode = RelayModeMidjourneyAction
	} else if strings.HasSuffix(path, "/mj/submit/modal") {
		// midjourney plus
		relayMode = RelayModeMidjourneyModal
	} else if strings.HasSuffix(path, "/mj/submit/shorten") {
		// midjourney plus
		relayMode = RelayModeMidjourneyShorten
	} else if strings.HasSuffix(path, "/mj/insight-face/swap") {
		// midjourney plus
		relayMode = RelayModeSwapFace
	} else if strings.HasSuffix(path, "/submit/upload-discord-images") {
		// midjourney plus
		relayMode = RelayModeMidjourneyUpload
	} else if strings.HasSuffix(path, "/mj/submit/imagine") {
		relayMode = RelayModeMidjourneyImagine
	} else if strings.HasSuffix(path, "/mj/submit/video") {
		relayMode = RelayModeMidjourneyVideo
	} else if strings.HasSuffix(path, "/mj/submit/edits") {
		relayMode = RelayModeMidjourneyEdits
	} else if strings.HasSuffix(path, "/mj/submit/blend") {
		relayMode = RelayModeMidjourneyBlend
	} else if strings.HasSuffix(path, "/mj/submit/describe") {
		relayMode = RelayModeMidjourneyDescribe
	} else if strings.HasSuffix(path, "/mj/notify") {
		relayMode = RelayModeMidjourneyNotify
	} else if strings.HasSuffix(path, "/mj/submit/change") {
		relayMode = RelayModeMidjourneyChange
	} else if strings.HasSuffix(path, "/mj/submit/simple-change") {
		relayMode = RelayModeMidjourneyChange
	} else if strings.HasSuffix(path, "/fetch") {
		relayMode = RelayModeMidjourneyTaskFetch
	} else if strings.HasSuffix(path, "/image-seed") {
		relayMode = RelayModeMidjourneyTaskImageSeed
	} else if strings.HasSuffix(path, "/list-by-condition") {
		relayMode = RelayModeMidjourneyTaskFetchByCondition
	}
	return relayMode
}

func Path2RelaySuno(method, path string) int {
	relayMode := RelayModeUnknown
	if method == http.MethodPost && strings.HasSuffix(path, "/fetch") {
		relayMode = RelayModeSunoFetch
	} else if method == http.MethodGet && strings.Contains(path, "/fetch/") {
		relayMode = RelayModeSunoFetchByID
	} else if strings.Contains(path, "/submit/") {
		relayMode = RelayModeSunoSubmit
	}
	return relayMode
}
